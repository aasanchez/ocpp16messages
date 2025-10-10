//go:build benchmark

package sharedtypes_test

import (
	"strconv"
	"sync/atomic"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func BenchmarkSetInteger_ValidInput(b *testing.B) {
	for range b.N {
		_, err := st.SetInteger("73")
		if err != nil {
			b.Errorf("unexpected error: %v", err)
		}
	}
}

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

func BenchmarkSetInteger_VariousInvalid(b *testing.B) {
	cases := []string{
		"",                // empty
		" ",               // space
		"-1",              // negative
		"65536",           // overflow
		"999999999999999", // way overflow
		"1.5",             // float
		"abc",             // non-numeric
		"+12",             // plus sign not allowed by ParseUint base10?
		"0x10",            // hex-like string but base10
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
