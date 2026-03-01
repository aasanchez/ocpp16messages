//go:build bench

package benchmark

import (
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

// Benchmarks are opt-in; run with `make test-bench`.

const (
	maxUint16PlusOne = 65536
	outOfRangeBase   = 70000
	outOfRangeModulo = 10
	integerSample    = 12345
	ciStringSample   = "RFID-ABC123"
)

func BenchmarkNewCiString20Type(b *testing.B) {
	const sample = "RFID-ABC1234567890"

	for i := 0; i < b.N; i++ {
		err := runBenchmarkCiString(sample)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString20Type_InvalidNonASCII(b *testing.B) {
	const sample = "bad\x01value"

	for i := 0; i < b.N; i++ {
		_, _ = st.NewCiString20Type(sample)
	}
}

func BenchmarkNewInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := runBenchmarkInteger(i % maxUint16PlusOne)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewInteger_OutOfRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = st.NewInteger(outOfRangeBase + (i % outOfRangeModulo))
	}
}

func BenchmarkIntegerString(b *testing.B) {
	val, _ := st.NewInteger(integerSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = val.String()
	}
}

func BenchmarkCiStringString(b *testing.B) {
	cs, _ := st.NewCiString20Type(ciStringSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cs.String()
	}
}

// Provide a varied set of numeric strings to parse into integers.
func BenchmarkNewInteger_ParseStrings(b *testing.B) {
	inputs := []string{"0", "123", "65535", "99999", "-1"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := inputs[i%len(inputs)]
		err := runBenchmarkInteger(parseStringToInt(s))
		if err != nil {
			// ignore errors; this benchmark includes invalid inputs
			_ = err
		}
	}
}

func parseStringToInt(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

func runBenchmarkCiString(value string) error {
	_, err := st.NewCiString20Type(value)
	return err
}

func runBenchmarkInteger(value int) error {
	_, err := st.NewInteger(value)
	return err
}
