//go:build bench

package benchmark

import (
	"fmt"
	"strconv"
	"testing"

	gconf "github.com/aasanchez/ocpp16messages/getConfiguration"
	st "github.com/aasanchez/ocpp16messages/types"
)

type primitiveGetConfigurationReq struct {
	Key []string
}

var (
	sinkGetConfigurationCustom    gconf.ReqMessage
	sinkGetConfigurationPrimitive primitiveGetConfigurationReq
)

func BenchmarkGetConfigurationReq_Custom_1(b *testing.B) {
	benchmarkGetConfigurationCustom(b, scaleOne)
}

func BenchmarkGetConfigurationReq_Custom_25(b *testing.B) {
	benchmarkGetConfigurationCustom(b, scaleTwentyFive)
}

func BenchmarkGetConfigurationReq_Custom_100(b *testing.B) {
	benchmarkGetConfigurationCustom(b, scaleOneHundred)
}

func BenchmarkGetConfigurationReq_Custom_250(b *testing.B) {
	benchmarkGetConfigurationCustom(b, scaleTwoHundred)
}

func BenchmarkGetConfigurationReq_Custom_500(b *testing.B) {
	benchmarkGetConfigurationCustom(b, scaleFiveHundred)
}

func BenchmarkGetConfigurationReq_Custom_1000(b *testing.B) {
	benchmarkGetConfigurationCustom(b, scaleOneThousand)
}

func BenchmarkGetConfigurationReq_PrimitiveDirect_1(b *testing.B) {
	benchmarkGetConfigurationPrimitiveDirect(b, scaleOne)
}

func BenchmarkGetConfigurationReq_PrimitiveDirect_25(b *testing.B) {
	benchmarkGetConfigurationPrimitiveDirect(b, scaleTwentyFive)
}

func BenchmarkGetConfigurationReq_PrimitiveDirect_100(b *testing.B) {
	benchmarkGetConfigurationPrimitiveDirect(b, scaleOneHundred)
}

func BenchmarkGetConfigurationReq_PrimitiveDirect_250(b *testing.B) {
	benchmarkGetConfigurationPrimitiveDirect(b, scaleTwoHundred)
}

func BenchmarkGetConfigurationReq_PrimitiveDirect_500(b *testing.B) {
	benchmarkGetConfigurationPrimitiveDirect(b, scaleFiveHundred)
}

func BenchmarkGetConfigurationReq_PrimitiveDirect_1000(b *testing.B) {
	benchmarkGetConfigurationPrimitiveDirect(b, scaleOneThousand)
}

func BenchmarkGetConfigurationReq_PrimitiveValidated_1(b *testing.B) {
	benchmarkGetConfigurationPrimitiveValidated(b, scaleOne)
}

func BenchmarkGetConfigurationReq_PrimitiveValidated_25(b *testing.B) {
	benchmarkGetConfigurationPrimitiveValidated(b, scaleTwentyFive)
}

func BenchmarkGetConfigurationReq_PrimitiveValidated_100(b *testing.B) {
	benchmarkGetConfigurationPrimitiveValidated(b, scaleOneHundred)
}

func BenchmarkGetConfigurationReq_PrimitiveValidated_250(b *testing.B) {
	benchmarkGetConfigurationPrimitiveValidated(b, scaleTwoHundred)
}

func BenchmarkGetConfigurationReq_PrimitiveValidated_500(b *testing.B) {
	benchmarkGetConfigurationPrimitiveValidated(b, scaleFiveHundred)
}

func BenchmarkGetConfigurationReq_PrimitiveValidated_1000(b *testing.B) {
	benchmarkGetConfigurationPrimitiveValidated(b, scaleOneThousand)
}

func benchmarkGetConfigurationCustom(b *testing.B, size int) {
	b.Helper()
	b.ReportAllocs()

	input := gconf.ReqInput{Key: makeConfigurationKeys(size)}

	for i := 0; i < b.N; i++ {
		message, err := gconf.Req(input)
		if err != nil {
			b.Fatal(err)
		}

		sinkGetConfigurationCustom = message
	}
}

func benchmarkGetConfigurationPrimitiveDirect(b *testing.B, size int) {
	b.Helper()
	b.ReportAllocs()

	input := primitiveGetConfigurationReq{Key: makeConfigurationKeys(size)}

	for i := 0; i < b.N; i++ {
		sinkGetConfigurationPrimitive = input
	}
}

func benchmarkGetConfigurationPrimitiveValidated(b *testing.B, size int) {
	b.Helper()
	b.ReportAllocs()

	input := primitiveGetConfigurationReq{Key: makeConfigurationKeys(size)}

	for i := 0; i < b.N; i++ {
		if err := validatePrimitiveGetConfigurationReq(input); err != nil {
			b.Fatal(err)
		}

		sinkGetConfigurationPrimitive = input
	}
}

func makeConfigurationKeys(size int) []string {
	keys := make([]string, 0, size)

	for index := 0; index < size; index++ {
		keys = append(keys, "HeartbeatInterval"+strconv.Itoa(index%10))
	}

	return keys
}

func validatePrimitiveGetConfigurationReq(input primitiveGetConfigurationReq) error {
	if len(input.Key) == 0 {
		return nil
	}

	for index, key := range input.Key {
		if err := validatePrimitiveCiString50(key); err != nil {
			return fmt.Errorf("key[%d]: %w", index, err)
		}
	}

	return nil
}

func validatePrimitiveCiString50(value string) error {
	if value == "" {
		return errPrimitiveEmpty
	}

	if len(value) > st.CiString50Max {
		return errPrimitiveTooLong
	}

	for index := 0; index < len(value); index++ {
		char := value[index]
		if char < 32 || char > 126 {
			return errPrimitiveBadASCII
		}
	}

	return nil
}
