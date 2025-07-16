package sharedtypes_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/shared/types"
)

func BenchmarkSetCiString20Type(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString20Type("test")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString20Type_Validate(b *testing.B) {
	ciString, err := sharedtypes.SetCiString20Type("test")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString25Type(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString25Type("test")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString25Type_Validate(b *testing.B) {
	ciString, err := sharedtypes.SetCiString25Type("test")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString50Type(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString50Type("test")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString50Type_Validate(b *testing.B) {
	ciString, err := sharedtypes.SetCiString50Type("test")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString255Type(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString255Type("test")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString255Type_Validate(b *testing.B) {
	ciString, err := sharedtypes.SetCiString255Type("test")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString500Type(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString500Type("test")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString500Type_Validate(b *testing.B) {
	ciString, err := sharedtypes.SetCiString500Type("test")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// Benchmarks with max length strings

func BenchmarkSetCiString20Type_Max(b *testing.B) {
	s := strings.Repeat("a", 20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString20Type(s)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString20Type_Validate_Max(b *testing.B) {
	s := strings.Repeat("a", 20)
	ciString, err := sharedtypes.SetCiString20Type(s)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString25Type_Max(b *testing.B) {
	s := strings.Repeat("a", 25)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString25Type(s)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString25Type_Validate_Max(b *testing.B) {
	s := strings.Repeat("a", 25)
	ciString, err := sharedtypes.SetCiString25Type(s)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString50Type_Max(b *testing.B) {
	s := strings.Repeat("a", 50)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString50Type(s)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString50Type_Validate_Max(b *testing.B) {
	s := strings.Repeat("a", 50)
	ciString, err := sharedtypes.SetCiString50Type(s)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString255Type_Max(b *testing.B) {
	s := strings.Repeat("a", 255)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString255Type(s)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString255Type_Validate_Max(b *testing.B) {
	s := strings.Repeat("a", 255)
	ciString, err := sharedtypes.SetCiString255Type(s)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkSetCiString500Type_Max(b *testing.B) {
	s := strings.Repeat("a", 500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sharedtypes.SetCiString500Type(s)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkCiString500Type_Validate_Max(b *testing.B) {
	s := strings.Repeat("a", 500)
	ciString, err := sharedtypes.SetCiString500Type(s)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ciString.Validate(); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}