//go:build benchmark

package sharedtypes_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type setterFn func(s string) error

func makeASCIIString(n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat("A", n)
}

func makeStringWithRune(n int, r rune) string {
	if n <= 0 {
		return string(r)
	}
	// place the rune roughly in the middle
	mid := n / 2
	b := make([]rune, n)
	for i := range b {
		b[i] = 'A'
	}
	if mid >= n {
		mid = n - 1
	}
	if mid < 0 {
		mid = 0
	}
	b[mid] = r
	return string(b)
}

func benchmarkSetter(b *testing.B, name string, set setterFn, input string, wantErr bool) {
	b.Helper()
	b.ReportAllocs()
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := set(input)
			if wantErr {
				if err == nil {
					b.Fatalf("expected error, got nil")
				}
			} else {
				if err != nil {
					b.Fatalf("unexpected error: %v", err)
				}
			}
		}
	})
}

func benchmarkSetterParallel(b *testing.B, name string, set setterFn, input string, wantErr bool) {
	b.Helper()
	b.ReportAllocs()
	b.Run(name+"/Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				err := set(input)
				if wantErr {
					if err == nil {
						b.Fatalf("expected error, got nil")
					}
				} else {
					if err != nil {
						b.Fatalf("unexpected error: %v", err)
					}
				}
			}
		})
	})
}

func setters() []struct {
	name string
	max  int
	fn   setterFn
} {
	return []struct {
		name string
		max  int
		fn   setterFn
	}{
		{"CiString20", 20, func(s string) error { _, err := st.SetCiString20Type(s); return err }},
		{"CiString25", 25, func(s string) error { _, err := st.SetCiString25Type(s); return err }},
		{"CiString50", 50, func(s string) error { _, err := st.SetCiString50Type(s); return err }},
		{"CiString255", 255, func(s string) error { _, err := st.SetCiString255Type(s); return err }},
		{"CiString500", 500, func(s string) error { _, err := st.SetCiString500Type(s); return err }},
	}
}

func BenchmarkSetCiString_AllScenarios(b *testing.B) {
	for _, s := range setters() {
		set := s
		b.Run(set.name, func(b *testing.B) {
			short := makeASCIIString(min(10, set.max))
			maxMinus1 := makeASCIIString(max(set.max-1, 0))
			exactMax := makeASCIIString(set.max)
			tooLong := makeASCIIString(set.max + 1)
			withNonPrintable := makeStringWithRune(max(set.max-1, 1), rune(0x1F))
			withDEL := makeStringWithRune(max(set.max-1, 1), rune(0x7F))
			withHighASCII := makeStringWithRune(max(set.max-1, 1), rune(0xE9)) // 'é'

			benchmarkSetter(b, "Short", set.fn, short, false)
			benchmarkSetter(b, "MaxMinus1", set.fn, maxMinus1, false)
			benchmarkSetter(b, "ExactMax", set.fn, exactMax, false)
			benchmarkSetter(b, "TooLong", set.fn, tooLong, true)
			benchmarkSetter(b, "NonPrintable_0x1F", set.fn, withNonPrintable, true)
			benchmarkSetter(b, "NonPrintable_0x7F", set.fn, withDEL, true)
			benchmarkSetter(b, "NonASCII_High", set.fn, withHighASCII, true)

			benchmarkSetterParallel(b, "Short", set.fn, short, false)
			benchmarkSetterParallel(b, "MaxMinus1", set.fn, maxMinus1, false)
			benchmarkSetterParallel(b, "ExactMax", set.fn, exactMax, false)
			benchmarkSetterParallel(b, "TooLong", set.fn, tooLong, true)
			benchmarkSetterParallel(b, "NonPrintable_0x1F", set.fn, withNonPrintable, true)
			benchmarkSetterParallel(b, "NonPrintable_0x7F", set.fn, withDEL, true)
			benchmarkSetterParallel(b, "NonASCII_High", set.fn, withHighASCII, true)
		})
	}
}

func BenchmarkCiString_Value(b *testing.B) {
	cases := []struct {
		name string
		mk   func() (any, func() string)
	}{
		{
			"CiString20",
			func() (any, func() string) { v, _ := st.SetCiString20Type(makeASCIIString(20)); return v, v.Value },
		},
		{
			"CiString25",
			func() (any, func() string) { v, _ := st.SetCiString25Type(makeASCIIString(25)); return v, v.Value },
		},
		{
			"CiString50",
			func() (any, func() string) { v, _ := st.SetCiString50Type(makeASCIIString(50)); return v, v.Value },
		},
		{
			"CiString255",
			func() (any, func() string) { v, _ := st.SetCiString255Type(makeASCIIString(255)); return v, v.Value },
		},
		{
			"CiString500",
			func() (any, func() string) { v, _ := st.SetCiString500Type(makeASCIIString(500)); return v, v.Value },
		},
	}

	for _, c := range cases {
		c := c
		b.Run(c.name, func(b *testing.B) {
			_, val := c.mk()
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = val()
			}
		})
	}
}

func BenchmarkSetCiString_MixedWorkload(b *testing.B) {
	types := setters()
	inputs := [][]struct {
		in      string
		wantErr bool
	}{}
	for _, s := range types {
		set := s
		arr := []struct {
			in      string
			wantErr bool
		}{
			{makeASCIIString(min(8, set.max)), false},
			{makeASCIIString(set.max), false},
			{makeASCIIString(set.max + 1), true},
			{makeStringWithRune(max(set.max-1, 1), rune(0x1F)), true},
			{makeStringWithRune(max(set.max-1, 1), rune(0xE9)), true},
		}
		inputs = append(inputs, arr)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for idx, s := range types {
			set := s
			inp := inputs[idx]
			for _, item := range inp {
				err := set.fn(item.in)
				if item.wantErr {
					if err == nil {
						b.Fatalf("expected error, got nil")
					}
				} else if err != nil {
					b.Fatalf("unexpected error: %v", err)
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
