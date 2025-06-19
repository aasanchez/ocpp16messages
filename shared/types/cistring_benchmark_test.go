package types

import (
	"strings"
	"testing"
)

func BenchmarkCiString20_Create(b *testing.B) {
	input := strings.Repeat("A", 20)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = CiString20(input)
	}
}

func BenchmarkCiString20_Validate(b *testing.B) {
	cs, _ := CiString20(strings.Repeat("A", 20))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Validate()
	}
}

func BenchmarkCiString20_Value(b *testing.B) {
	cs, _ := CiString20(strings.Repeat("A", 20))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Value()
	}
}

func BenchmarkCiString25_Create(b *testing.B) {
	input := strings.Repeat("A", 25)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = CiString25(input)
	}
}

func BenchmarkCiString25_Validate(b *testing.B) {
	input := strings.Repeat("A", 25)
	cs, _ := CiString25(input)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Validate()
	}
}

func BenchmarkCiString25_Value(b *testing.B) {
	input := strings.Repeat("A", 25)
	cs, _ := CiString25(input)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Value()
	}
}

func BenchmarkCiString50_Create(b *testing.B) {
	input := strings.Repeat("A", 50)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = CiString50(input)
	}
}

func BenchmarkCiString50_Validate(b *testing.B) {
	input := strings.Repeat("A", 50)
	cs, _ := CiString50(input)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Validate()
	}
}

func BenchmarkCiString50_Value(b *testing.B) {
	input := strings.Repeat("A", 50)
	cs, _ := CiString50(input)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Value()
	}
}

func BenchmarkCiString255_Create(b *testing.B) {
	input := strings.Repeat("D", 255)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = CiString255(input)
	}
}

func BenchmarkCiString255_Validate(b *testing.B) {
	cs, _ := CiString255(strings.Repeat("D", 255))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Validate()
	}
}

func BenchmarkCiString255_Value(b *testing.B) {
	cs, _ := CiString255(strings.Repeat("D", 255))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Value()
	}
}

func BenchmarkCiString500_Create(b *testing.B) {
	input := strings.Repeat("A", 500)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = CiString500(input)
	}
}

func BenchmarkCiString500_Validate(b *testing.B) {
	input := strings.Repeat("A", 500)
	cs, _ := CiString500(input)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Validate()
	}
}

func BenchmarkCiString500_Value(b *testing.B) {
	input := strings.Repeat("A", 500)
	cs, _ := CiString500(input)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = cs.Value()
	}
}
