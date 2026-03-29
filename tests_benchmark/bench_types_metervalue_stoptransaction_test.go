//go:build bench

package benchmark

import (
	"testing"

	mv "github.com/aasanchez/ocpp16messages/metervalues"
	stp "github.com/aasanchez/ocpp16messages/stoptransaction"
	types "github.com/aasanchez/ocpp16types"
)

const (
	sampleTimestamp = "2025-01-02T15:00:00Z"
	sampleValue     = "12345"
	measurandEnergy = "Energy.Active.Import.Register"
	unitWh          = "Wh"
)

func BenchmarkNewDateTime(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := types.NewDateTime(sampleTimestamp); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDateTimeString(b *testing.B) {
	b.ReportAllocs()

	dt, _ := types.NewDateTime(sampleTimestamp)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dt.String()
	}
}

func BenchmarkNewSampledValue(b *testing.B) {
	b.ReportAllocs()

	measurand := measurandEnergy
	unit := unitWh
	ctx := "Sample.Periodic"

	input := types.SampledValueInput{
		Value:     sampleValue,
		Context:   &ctx,
		Measurand: &measurand,
		Unit:      &unit,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewSampledValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewMeterValue_Single(b *testing.B) {
	b.ReportAllocs()

	input := types.MeterValueInput{
		Timestamp: sampleTimestamp,
		SampledValue: []types.SampledValueInput{
			{Value: sampleValue},
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewMeterValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewMeterValue_ManySampled(b *testing.B) {
	b.ReportAllocs()

	var samples []types.SampledValueInput
	for i := 0; i < 10; i++ {
		samples = append(samples, types.SampledValueInput{
			Value: sampleValue,
		})
	}

	input := types.MeterValueInput{
		Timestamp:    sampleTimestamp,
		SampledValue: samples,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewMeterValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMeterValuesReq_SingleValue(b *testing.B) {
	b.ReportAllocs()

	reqInput := mv.ReqInput{
		ConnectorId: 1,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: sampleTimestamp,
				SampledValue: []types.SampledValueInput{
					{Value: sampleValue},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := mv.Req(reqInput); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMeterValuesReq_ManySampledValues(b *testing.B) {
	b.ReportAllocs()

	var samples []types.SampledValueInput
	for i := 0; i < 10; i++ {
		samples = append(samples, types.SampledValueInput{
			Value: sampleValue,
		})
	}

	reqInput := mv.ReqInput{
		ConnectorId: 1,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp:    sampleTimestamp,
				SampledValue: samples,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := mv.Req(reqInput); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString25Type(b *testing.B) {
	const sample = "RFID-ABC1234567890123"

	for i := 0; i < b.N; i++ {
		if _, err := types.NewCiString25Type(sample); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString50Type(b *testing.B) {
	const sample = "RFID-ABC1234567890123456789012345678901234567890"

	for i := 0; i < b.N; i++ {
		if _, err := types.NewCiString50Type(sample); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString255Type(b *testing.B) {
	long := make([]byte, 255)
	for i := range long {
		long[i] = 'A'
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewCiString255Type(string(long)); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewIdToken(b *testing.B) {
	b.ReportAllocs()

	ci, _ := types.NewCiString20Type(ciStringSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = types.NewIdToken(ci)
	}
}

func BenchmarkNewIdTagInfo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := types.NewIdTagInfo(types.AuthorizationStatusAccepted); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStopTransactionReq_Minimal(b *testing.B) {
	b.ReportAllocs()

	reqInput := stp.ReqInput{
		TransactionId: 1,
		MeterStop:     100,
		Timestamp:     sampleTimestamp,
	}

	for i := 0; i < b.N; i++ {
		if _, err := stp.Req(reqInput); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStopTransactionReq_WithTransactionData(b *testing.B) {
	b.ReportAllocs()

	reqInput := stp.ReqInput{
		TransactionId: 1,
		MeterStop:     100,
		Timestamp:     sampleTimestamp,
		TransactionData: []types.MeterValueInput{
			{
				Timestamp: sampleTimestamp,
				SampledValue: []types.SampledValueInput{
					{Value: sampleValue},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := stp.Req(reqInput); err != nil {
			b.Fatal(err)
		}
	}
}
