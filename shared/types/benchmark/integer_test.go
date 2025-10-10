//go:build benchmark

// Package sharedtypes_test provides micro-
// benchmarks for the Integer type used in
// OCPP 1.6 payloads. It measures parse,
// allocation and concurrency cost for
// 0..65535 values that represent connector
// indices, counters and compact ids.
//
// These numbers help keep CP and CSMS
// message handling fast and predictable,
// even under invalid inputs and bursts.
//
// Run with:
//
//	go test -tags benchmark -bench . -benchmem
package sharedtypes_test

import (
	"strconv"
	"sync/atomic"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// BenchmarkSetInteger_ValidInput measures the
// common-case cost to parse a typical valid
// value (e.g., connector index). Helps ensure
// low latency in CP/CSMS message flows.
func BenchmarkSetInteger_ValidInput(b *testing.B) {
	for range b.N {
		_, err := st.SetInteger("73")
		if err != nil {
			b.Errorf("unexpected error: %v", err)
		}
	}
}

// BenchmarkInteger_Value measures accessor
// overhead to read the stored uint16. It
// should be O(1) with zero allocations.
func BenchmarkInteger_Value(b *testing.B) {
	integer, err := st.SetInteger("73")
	if err != nil {
		b.Errorf("unexpected error: %v", err)
	}

	b.ResetTimer()

	for range b.N {
		_ = integer.Value()
	}
}

// BenchmarkSetInteger_VariousValid exercises
// valid edge points from 0 up to 65535.
// It reports allocations to catch regressions.
func BenchmarkSetInteger_VariousValid(b *testing.B) {
	cases := []string{
		"0", // min
		"1",
		"9",
		"10",
		"99",
		"100",
		"1024",
		"65535", // max
	}
	b.ReportAllocs()
	for _, s := range cases {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := st.SetInteger(s)
				if err != nil {
					b.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

// BenchmarkSetInteger_VariousInvalid measures
// rejection paths: empty, negative, overflow,
// floats, and non decimal text. Ensures fast
// and safe failure under bad inputs.
func BenchmarkSetInteger_VariousInvalid(b *testing.B) {
	cases := []string{
		"",                // empty
		" ",               // space
		"-1",              // negative
		"65536",           // overflow
		"999999999999999", // way overflow
		"1.5",             // float
		"abc",             // non-numeric
		"+12",             // explicit plus sign
		"0x10",            // hex-like but base10
	}
	b.ReportAllocs()
	for _, s := range cases {
		name := s
		if len(name) > 12 {
			name = name[:12]
		}
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := st.SetInteger(s)
				if err == nil {
					b.Errorf("expected error for %q, got nil", s)
				}
			}
		})
	}
}

// BenchmarkSetInteger_MixedWorkload simulates
// production streams mixing valid and invalid
// values. Useful to gauge steady-state cost.
func BenchmarkSetInteger_MixedWorkload(b *testing.B) {
	inputs := []struct {
		in  string
		err bool
	}{
		{"0", false},
		{"1", false},
		{"65535", false},
		{"73", false},
		{"65536", true},
		{"-1", true},
		{"abc", true},
		{"", true},
	}
	b.ReportAllocs()
	b.ResetTimer()
	idx := 0
	for i := 0; i < b.N; i++ {
		item := inputs[idx]
		idx++
		if idx == len(inputs) {
			idx = 0
		}
		_, err := st.SetInteger(item.in)
		if item.err {
			if err == nil {
				b.Errorf("expected error for %q, got nil", item.in)
			}
		} else if err != nil {
			b.Errorf("unexpected error for %q: %v", item.in, err)
		}
	}
}

// BenchmarkSetInteger_Parallel exercises parse
// under parallelism to observe contention and
// allocation behavior in concurrent handlers.
func BenchmarkSetInteger_Parallel(b *testing.B) {
	inputs := []string{"0", "1", "10", "100", "1024", "65535"}
	var ctr uint64
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idx := int(atomic.AddUint64(&ctr, 1))
			s := inputs[idx%len(inputs)]
			_, err := st.SetInteger(s)
			if err != nil {
				b.Errorf("unexpected error: %v", err)
			}
		}
	})
}

// BenchmarkSetInteger_ParseUintBaseline provides
// a baseline using strconv.ParseUint. Compare to
// SetInteger to understand wrapper overhead.
func BenchmarkSetInteger_ParseUintBaseline(b *testing.B) {
	// Baseline to understand SetInteger overhead vs direct strconv.ParseUint
	inputs := []string{"73", "1024", "65535"}
	b.ReportAllocs()
	for _, s := range inputs {
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := strconv.ParseUint(s, 10, 16)
				if err != nil {
					b.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
