//go:build bench

package benchmark

import (
	"testing"

	stt "github.com/aasanchez/ocpp16messages/startTransaction"
)

var (
	sinkStartTxCustom    stt.ReqMessage
	sinkStartTxPrimitive primitiveStartTransactionReq
)

func BenchmarkStartTransactionReq_Custom(b *testing.B) {
	b.ReportAllocs()

	reservationId := 42
	input := stt.ReqInput{
		ConnectorId:   1,
		IdTag:         "TAG-123",
		MeterStart:    100,
		Timestamp:     benchmarkTimestamp,
		ReservationId: &reservationId,
	}

	for i := 0; i < b.N; i++ {
		message, err := stt.Req(input)
		if err != nil {
			b.Fatal(err)
		}

		sinkStartTxCustom = message
	}
}

func BenchmarkStartTransactionReq_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	reservationId := 42
	input := primitiveStartTransactionReq{
		ConnectorId:   1,
		IdTag:         "TAG-123",
		MeterStart:    100,
		Timestamp:     benchmarkTimestamp,
		ReservationId: &reservationId,
	}

	for i := 0; i < b.N; i++ {
		sinkStartTxPrimitive = input
	}
}

func BenchmarkStartTransactionReq_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	reservationId := 42
	input := primitiveStartTransactionReq{
		ConnectorId:   1,
		IdTag:         "TAG-123",
		MeterStart:    100,
		Timestamp:     benchmarkTimestamp,
		ReservationId: &reservationId,
	}

	for i := 0; i < b.N; i++ {
		if err := validatePrimitiveStartTransactionReq(input); err != nil {
			b.Fatal(err)
		}

		sinkStartTxPrimitive = input
	}
}
