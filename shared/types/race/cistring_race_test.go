//go:build race

package sharedtypes_test

import (
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// randomSleepCS increases goroutine interleaving to maximize race detection.
func randomSleepCS() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
}

// makeASCIIString returns a string of length n with printable ASCII chars.
func makeASCIIString(n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat("A", n)
}

// makeStringWithRune returns a string of length n with a special rune in the middle.
func makeStringWithRune(n int, r rune) string {
	if n <= 0 {
		return string(r)
	}
	mid := n / 2
	b := make([]rune, n)
	for i := range b {
		b[i] = 'A'
	}
	b[mid] = r
	return string(b)
}

func TestCiStringRace_ConcurrentSetAndValue(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	types := []struct {
		name string
		max  int
		set  func(string) (any, error)
		val  func(any) string
	}{
		{"CiString20", 20, func(s string) (any, error) { return st.SetCiString20Type(s) }, func(v any) string { return v.(st.CiString20Type).Value() }},
		{"CiString25", 25, func(s string) (any, error) { return st.SetCiString25Type(s) }, func(v any) string { return v.(st.CiString25Type).Value() }},
		{"CiString50", 50, func(s string) (any, error) { return st.SetCiString50Type(s) }, func(v any) string { return v.(st.CiString50Type).Value() }},
		{"CiString255", 255, func(s string) (any, error) { return st.SetCiString255Type(s) }, func(v any) string { return v.(st.CiString255Type).Value() }},
		{"CiString500", 500, func(s string) (any, error) { return st.SetCiString500Type(s) }, func(v any) string { return v.(st.CiString500Type).Value() }},
	}
	for _, typ := range types {
		for i := 0; i < 50; i++ {
			inputs := []string{
				makeASCIIString(typ.max),
				makeASCIIString(typ.max - 1),
				makeASCIIString(1),
				makeStringWithRune(typ.max, rune(0x1F)), // non-printable
				makeStringWithRune(typ.max, rune(0x7F)), // DEL
				makeStringWithRune(typ.max, rune(0xE9)), // non-ASCII
				makeASCIIString(typ.max + 1),            // too long
				"",                                      // empty
			}
			for _, input := range inputs {
				wg.Add(1)
				go func(typ struct {
					name string
					max  int
					set  func(string) (any, error)
					val  func(any) string
				}, input string) {
					defer wg.Done()
					randomSleepCS()
					v, err := typ.set(input)
					if err == nil {
						_ = typ.val(v)
					}
				}(typ, input)
			}
		}
	}
	wg.Wait()
}

func TestCiStringRace_SharedInstanceValue(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	v, err := st.SetCiString255Type(makeASCIIString(255))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randomSleepCS()
			val := v.Value()
			if len(val) != 255 {
				t.Errorf("expected length 255, got %d", len(val))
			}
		}()
	}
	wg.Wait()
}

func TestCiStringRace_InterleavedSetAndValue(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	typ := struct {
		name string
		max  int
		set  func(string) (any, error)
		val  func(any) string
	}{
		"CiString50", 50, func(s string) (any, error) { return st.SetCiString50Type(s) }, func(v any) string { return v.(st.CiString50Type).Value() },
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var input string
			if i%2 == 0 {
				input = makeASCIIString(typ.max)
			} else {
				input = makeStringWithRune(typ.max, rune(0x1F))
			}
			v, err := typ.set(input)
			randomSleepCS()
			if err == nil {
				val := typ.val(v)
				if len(val) != typ.max {
					t.Errorf("expected length %d, got %d", typ.max, len(val))
				}
			}
		}(i)
	}
	wg.Wait()
}

func TestCiStringRace_Stress(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	types := []struct {
		name string
		max  int
		set  func(string) (any, error)
		val  func(any) string
	}{
		{"CiString20", 20, func(s string) (any, error) { return st.SetCiString20Type(s) }, func(v any) string { return v.(st.CiString20Type).Value() }},
		{"CiString255", 255, func(s string) (any, error) { return st.SetCiString255Type(s) }, func(v any) string { return v.(st.CiString255Type).Value() }},
	}
	nGoroutines := 2000
	for i := 0; i < nGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			typ := types[rand.Intn(len(types))]
			var input string
			if rand.Intn(2) == 0 {
				input = makeASCIIString(typ.max)
			} else {
				input = makeStringWithRune(typ.max, rune(0x1F))
			}
			v, err := typ.set(input)
			randomSleepCS()
			if err == nil {
				_ = typ.val(v)
			}
		}(i)
	}
	wg.Wait()
}

// Document: If CiString types ever become mutable, add tests for concurrent mutation and reading.
