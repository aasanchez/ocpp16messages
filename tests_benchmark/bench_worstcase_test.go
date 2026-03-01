//go:build bench

package benchmark

import (
	"testing"

	mv "github.com/aasanchez/ocpp16messages/meterValues/types"
	scpt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	stp "github.com/aasanchez/ocpp16messages/stopTransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

func BenchmarkNewSampledValue_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	context := st.ReadingContextSamplePeriodic.String()
	format := st.ValueFormatRaw.String()
	measurand := st.MeasurandEnergyActiveImportRegister.String()
	phase := st.PhaseL1.String()
	location := st.LocationOutlet.String()
	unit := st.UnitWh.String()

	input := st.SampledValueInput{
		Value:     sampleValue,
		Context:   &context,
		Format:    &format,
		Measurand: &measurand,
		Phase:     &phase,
		Location:  &location,
		Unit:      &unit,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewSampledValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSetChargingProfileNewChargingProfile_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	transactionId := 1

	recurrencyKind := scpt.RecurrencyKindDaily.String()
	validFrom := sampleTimestamp
	validTo := "2025-01-02T16:00:00Z"

	duration := 60
	scheduleStart := sampleTimestamp
	minChargingRate := 0.0

	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	input := scpt.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          &transactionId,
		StackLevel:             0,
		ChargingProfilePurpose: st.TxProfile.String(),
		ChargingProfileKind:    scpt.ChargingProfileKindRecurring.String(),
		RecurrencyKind:         &recurrencyKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       scheduleInput,
	}

	for i := 0; i < b.N; i++ {
		if _, err := scpt.NewChargingProfile(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStopTransactionReq_LargeTransactionData(b *testing.B) {
	b.ReportAllocs()

	const meterValuesCount = 10
	const sampledValuesCount = 10

	var sampledValues []mv.SampledValueInput
	for i := 0; i < sampledValuesCount; i++ {
		sampledValues = append(sampledValues, mv.SampledValueInput{Value: sampleValue})
	}

	var transactionData []mv.MeterValueInput
	for i := 0; i < meterValuesCount; i++ {
		transactionData = append(transactionData, mv.MeterValueInput{
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
