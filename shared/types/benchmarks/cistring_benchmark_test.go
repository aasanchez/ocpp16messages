package types_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/shared/types"
)

func BenchmarkCiString20_Create(b *testing.B) {
	input := strings.Repeat("A", 20)

	b.ReportAllocs()

	for range b.N {
		_, _ = types.SetCiString20(input)
	}
}

func BenchmarkCiString20_Validate(b *testing.B) {
	cs, _ := types.SetCiString20(strings.Repeat("A", 20))

	b.ReportAllocs()

	for range b.N {
		_ = cs.Validate()
	}
}

func BenchmarkCiString20_Value(b *testing.B) {
	cs, _ := types.SetCiString20(strings.Repeat("A", 20))

	b.ReportAllocs()

	for range b.N {
		_ = cs.Value()
	}
}

func BenchmarkCiString25_Create(b *testing.B) {
	input := strings.Repeat("A", 25)

	b.ReportAllocs()

	for range b.N {
		_, _ = types.SetCiString25(input)
	}
}

func BenchmarkCiString25_Validate(b *testing.B) {
	input := strings.Repeat("A", 25)
	cs, _ := types.SetCiString25(input)

	b.ReportAllocs()

	for range b.N {
		_ = cs.Validate()
	}
}

func BenchmarkCiString25_Value(b *testing.B) {
	input := strings.Repeat("A", 25)
	cs, _ := types.SetCiString25(input)

	b.ReportAllocs()

	for range b.N {
		_ = cs.Value()
	}
}

func BenchmarkCiString50_Create(b *testing.B) {
	input := strings.Repeat("A", 50)

	b.ReportAllocs()

	for range b.N {
		_, _ = types.SetCiString50(input)
	}
}

func BenchmarkCiString50_Validate(b *testing.B) {
	input := strings.Repeat("A", 50)
	cs, _ := types.SetCiString50(input)

	b.ReportAllocs()

	for range b.N {
		_ = cs.Validate()
	}
}

func BenchmarkCiString50_Value(b *testing.B) {
	input := strings.Repeat("A", 50)
	cs, _ := types.SetCiString50(input)

	b.ReportAllocs()

	for range b.N {
		_ = cs.Value()
	}
}

func BenchmarkCiString255_Create(b *testing.B) {
	input := strings.Repeat("D", 255)

	b.ReportAllocs()

	for range b.N {
		_, _ = types.SetCiString255(input)
	}
}

func BenchmarkCiString255_Validate(b *testing.B) {
	cs, _ := types.SetCiString255(strings.Repeat("D", 255))

	b.ReportAllocs()

	for range b.N {
		_ = cs.Validate()
	}
}

func BenchmarkCiString255_Value(b *testing.B) {
	cs, _ := types.SetCiString255(strings.Repeat("D", 255))

	b.ReportAllocs()

	for range b.N {
		_ = cs.Value()
	}
}

func BenchmarkCiString500_Create(b *testing.B) {
	input := strings.Repeat("A", 500)

	b.ReportAllocs()

	for range b.N {
		_, _ = types.SetCiString500(input)
	}
}

func BenchmarkCiString500_Validate(b *testing.B) {
	input := strings.Repeat("A", 500)
	cs, _ := types.SetCiString500(input)

	b.ReportAllocs()

	for range b.N {
		_ = cs.Validate()
	}
}

func BenchmarkCiString500_Value(b *testing.B) {
	input := strings.Repeat("A", 500)
	cs, _ := types.SetCiString500(input)

	b.ReportAllocs()

	for range b.N {
		_ = cs.Value()
	}
}
