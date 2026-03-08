//go:build bench

package benchmark

import (
	"testing"

	sll "github.com/aasanchez/ocpp16messages/sendLocalList"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
)

var (
	sinkSLLCustom    sll.ReqMessage
	sinkSLLPrimitive primitiveSendLocalListReq
)

func BenchmarkSendLocalListReq_Custom_1(b *testing.B) {
	benchmarkSendLocalListCustom(b, scaleOne)
}

func BenchmarkSendLocalListReq_Custom_25(b *testing.B) {
	benchmarkSendLocalListCustom(b, scaleTwentyFive)
}

func BenchmarkSendLocalListReq_Custom_100(b *testing.B) {
	benchmarkSendLocalListCustom(b, scaleOneHundred)
}

func BenchmarkSendLocalListReq_Custom_250(b *testing.B) {
	benchmarkSendLocalListCustom(b, scaleTwoHundred)
}

func BenchmarkSendLocalListReq_Custom_500(b *testing.B) {
	benchmarkSendLocalListCustom(b, scaleFiveHundred)
}

func BenchmarkSendLocalListReq_Custom_1000(b *testing.B) {
	benchmarkSendLocalListCustom(b, scaleOneThousand)
}

func BenchmarkSendLocalListReq_PrimitiveDirect_1(b *testing.B) {
	benchmarkSendLocalListPrimitiveDirect(b, scaleOne)
}

func BenchmarkSendLocalListReq_PrimitiveDirect_25(b *testing.B) {
	benchmarkSendLocalListPrimitiveDirect(b, scaleTwentyFive)
}

func BenchmarkSendLocalListReq_PrimitiveDirect_100(b *testing.B) {
	benchmarkSendLocalListPrimitiveDirect(b, scaleOneHundred)
}

func BenchmarkSendLocalListReq_PrimitiveDirect_250(b *testing.B) {
	benchmarkSendLocalListPrimitiveDirect(b, scaleTwoHundred)
}

func BenchmarkSendLocalListReq_PrimitiveDirect_500(b *testing.B) {
	benchmarkSendLocalListPrimitiveDirect(b, scaleFiveHundred)
}

func BenchmarkSendLocalListReq_PrimitiveDirect_1000(b *testing.B) {
	benchmarkSendLocalListPrimitiveDirect(b, scaleOneThousand)
}

func BenchmarkSendLocalListReq_PrimitiveValidated_1(b *testing.B) {
	benchmarkSendLocalListPrimitiveValidated(b, scaleOne)
}

func BenchmarkSendLocalListReq_PrimitiveValidated_25(b *testing.B) {
	benchmarkSendLocalListPrimitiveValidated(b, scaleTwentyFive)
}

func BenchmarkSendLocalListReq_PrimitiveValidated_100(b *testing.B) {
	benchmarkSendLocalListPrimitiveValidated(b, scaleOneHundred)
}

func BenchmarkSendLocalListReq_PrimitiveValidated_250(b *testing.B) {
	benchmarkSendLocalListPrimitiveValidated(b, scaleTwoHundred)
}

func BenchmarkSendLocalListReq_PrimitiveValidated_500(b *testing.B) {
	benchmarkSendLocalListPrimitiveValidated(b, scaleFiveHundred)
}

func BenchmarkSendLocalListReq_PrimitiveValidated_1000(b *testing.B) {
	benchmarkSendLocalListPrimitiveValidated(b, scaleOneThousand)
}

func benchmarkSendLocalListCustom(b *testing.B, size int) {
	b.Helper()
	b.ReportAllocs()

	input := sll.ReqInput{
		ListVersion:            1,
		LocalAuthorizationList: makeAuthorizationInputs(size),
		UpdateType:             slt.UpdateTypeFull.String(),
	}

	for i := 0; i < b.N; i++ {
		message, err := sll.Req(input)
		if err != nil {
			b.Fatal(err)
		}

		sinkSLLCustom = message
	}
}

func benchmarkSendLocalListPrimitiveDirect(b *testing.B, size int) {
	b.Helper()
	b.ReportAllocs()

	input := primitiveSendLocalListReq{
		ListVersion:            1,
		LocalAuthorizationList: makePrimitiveAuthorizationInputs(size),
		UpdateType:             slt.UpdateTypeFull.String(),
	}

	for i := 0; i < b.N; i++ {
		sinkSLLPrimitive = input
	}
}

func benchmarkSendLocalListPrimitiveValidated(b *testing.B, size int) {
	b.Helper()
	b.ReportAllocs()

	input := primitiveSendLocalListReq{
		ListVersion:            1,
		LocalAuthorizationList: makePrimitiveAuthorizationInputs(size),
		UpdateType:             slt.UpdateTypeFull.String(),
	}

	for i := 0; i < b.N; i++ {
		if err := validatePrimitiveSendLocalListReq(input); err != nil {
			b.Fatal(err)
		}

		sinkSLLPrimitive = input
	}
}
