//go:build bench

package benchmark

import (
	"testing"

	stp "github.com/aasanchez/ocpp16messages/stoptransaction"
	types "github.com/aasanchez/ocpp16types"
)

func BenchmarkNewSampledValue_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	context := types.ReadingContextSamplePeriodic.String()
	format := types.ValueFormatRaw.String()
	measurand := types.MeasurandEnergyActiveImportRegister.String()
	phase := types.PhaseL1.String()
	location := types.LocationOutlet.String()
	unit := types.UnitWh.String()

	input := types.SampledValueInput{
		Value:     sampleValue,
		Context:   &context,
		Format:    &format,
		Measurand: &measurand,
		Phase:     &phase,
		Location:  &location,
		Unit:      &unit,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewSampledValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSetChargingProfileNewChargingProfile_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	transactionId := 1

	recurrencyKind := types.RecurrencyKindDaily.String()
	validFrom := sampleTimestamp
	validTo := "2025-01-02T16:00:00Z"

	duration := 60
	scheduleStart := sampleTimestamp
	minChargingRate := 0.0

	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	input := types.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          &transactionId,
		StackLevel:             0,
		ChargingProfilePurpose: types.TxProfile.String(),
		ChargingProfileKind:    types.ChargingProfileKindRecurring.String(),
		RecurrencyKind:         &recurrencyKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       scheduleInput,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewChargingProfile(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStopTransactionReq_LargeTransactionData(b *testing.B) {
	b.ReportAllocs()

	const metervaluesCount = 10
	const sampledValuesCount = 10

	var sampledValues []types.SampledValueInput
	for i := 0; i < sampledValuesCount; i++ {
		sampledValues = append(sampledValues, types.SampledValueInput{Value: sampleValue})
	}

	var transactionData []types.MeterValueInput
	for i := 0; i < metervaluesCount; i++ {
		transactionData = append(transactionData, types.MeterValueInput{
			Timestamp:    sampleTimestamp,
			SampledValue: sampledValues,
		})
	}

	input := stp.ReqInput{
		TransactionId:   1,
		IdTag:           nil,
		MeterStop:       100,
		Timestamp:       sampleTimestamp,
		Reason:          nil,
		TransactionData: transactionData,
	}

	for i := 0; i < b.N; i++ {
		if _, err := stp.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}
