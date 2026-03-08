//go:build bench

package benchmark

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	sll "github.com/aasanchez/ocpp16messages/sendLocalList"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	stt "github.com/aasanchez/ocpp16messages/startTransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	benchmarkTimestamp = "2025-01-02T15:00:00Z"

	scaleOneThousand = 1000
	scaleFiveHundred = 500
	scaleTwoHundred  = 250
	scaleOneHundred  = 100
	scaleTwentyFive  = 25
	scaleOne         = 1

	maxUint16 = 65535
)

var (
	errPrimitiveInvalidRange = errors.New("value out of uint16 range")

	sinkDateTimeCustom    st.DateTime
	sinkDateTimePrimitive string

	sinkStartTxCustom    stt.ReqMessage
	sinkStartTxPrimitive primitiveStartTransactionReq

	sinkSLLCustom    sll.ReqMessage
	sinkSLLPrimitive primitiveSendLocalListReq
)

type primitiveStartTransactionReq struct {
	ConnectorId   int
	IdTag         string
	MeterStart    int
	Timestamp     string
	ReservationId *int
}

type primitiveAuthorizationDataInput struct {
	IdTag string
}

type primitiveSendLocalListReq struct {
	ListVersion            int
	LocalAuthorizationList []primitiveAuthorizationDataInput
	UpdateType             string
}

func BenchmarkDateTime_CustomType(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		dt, err := st.NewDateTime(benchmarkTimestamp)
		if err != nil {
			b.Fatal(err)
		}

		sinkDateTimeCustom = dt
	}
}

func BenchmarkDateTime_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sinkDateTimePrimitive = benchmarkTimestamp
	}
}

func BenchmarkDateTime_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := validatePrimitiveTimestampUTC(benchmarkTimestamp); err != nil {
			b.Fatal(err)
		}

		sinkDateTimePrimitive = benchmarkTimestamp
	}
}

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

func makeAuthorizationInputs(size int) []slt.AuthorizationDataInput {
	entries := make([]slt.AuthorizationDataInput, 0, size)

	for index := 0; index < size; index++ {
		entries = append(entries, slt.AuthorizationDataInput{
			IdTag: "TAG-" + strconv.Itoa(index),
		})
	}

	return entries
}

func makePrimitiveAuthorizationInputs(size int) []primitiveAuthorizationDataInput {
	entries := make([]primitiveAuthorizationDataInput, 0, size)

	for index := 0; index < size; index++ {
		entries = append(entries, primitiveAuthorizationDataInput{
			IdTag: "TAG-" + strconv.Itoa(index),
		})
	}

	return entries
}

func validatePrimitiveStartTransactionReq(input primitiveStartTransactionReq) error {
	if err := validatePrimitiveIntegerRange(input.ConnectorId); err != nil {
		return fmt.Errorf("connectorId: %w", err)
	}

	if err := validatePrimitiveCiString20(input.IdTag); err != nil {
		return fmt.Errorf("idTag: %w", err)
	}

	if err := validatePrimitiveIntegerRange(input.MeterStart); err != nil {
		return fmt.Errorf("meterStart: %w", err)
	}

	if err := validatePrimitiveTimestampUTC(input.Timestamp); err != nil {
		return fmt.Errorf("timestamp: %w", err)
	}

	if input.ReservationId != nil {
		if err := validatePrimitiveIntegerRange(*input.ReservationId); err != nil {
			return fmt.Errorf("reservationId: %w", err)
		}
	}

	return nil
}

func validatePrimitiveSendLocalListReq(input primitiveSendLocalListReq) error {
	if err := validatePrimitiveIntegerRange(input.ListVersion); err != nil {
		return fmt.Errorf("listVersion: %w", err)
	}

	if input.UpdateType != slt.UpdateTypeFull.String() &&
		input.UpdateType != slt.UpdateTypeDifferential.String() {
		return fmt.Errorf("updateType: %w", st.ErrInvalidValue)
	}

	if input.LocalAuthorizationList == nil {
		return nil
	}

	for index, entry := range input.LocalAuthorizationList {
		if err := validatePrimitiveCiString20(entry.IdTag); err != nil {
			return fmt.Errorf("localAuthorizationList[%d]: %w", index, err)
		}
	}

	return nil
}

func validatePrimitiveIntegerRange(value int) error {
	if value < 0 || value > maxUint16 {
		return errPrimitiveInvalidRange
	}

	return nil
}

func validatePrimitiveCiString20(value string) error {
	if value == "" {
		return errPrimitiveEmpty
	}

	if len(value) > st.CiString20Max {
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

func validatePrimitiveTimestampUTC(value string) error {
	if value == "" {
		return errPrimitiveEmpty
	}

	timestamp, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return fmt.Errorf("%w: %w", st.ErrInvalidValue, err)
	}

	_, offset := timestamp.Zone()
	if offset != 0 {
		return fmt.Errorf(
			"%w: expected UTC offset, got %s",
			st.ErrInvalidValue,
			timestamp.Format("Z07:00"),
		)
	}

	return nil
}
