//go:build bench

package benchmark

import (
	"testing"

	mv "github.com/aasanchez/ocpp16messages/metervalues"
	mt "github.com/aasanchez/ocpp16messages/metervalues/types"
	stp "github.com/aasanchez/ocpp16messages/stoptransaction"
	st "github.com/aasanchez/ocpp16messages/types"
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
		if _, err := st.NewDateTime(sampleTimestamp); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDateTimeString(b *testing.B) {
	b.ReportAllocs()

	dt, _ := st.NewDateTime(sampleTimestamp)

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

	input := st.SampledValueInput{
		Value:     sampleValue,
		Context:   &ctx,
		Measurand: &measurand,
		Unit:      &unit,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewSampledValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewMeterValue_Single(b *testing.B) {
	b.ReportAllocs()

	input := st.MeterValueInput{
		Timestamp: sampleTimestamp,
		SampledValue: []st.SampledValueInput{
			{Value: sampleValue},
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewMeterValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewMeterValue_ManySampled(b *testing.B) {
	b.ReportAllocs()

	var samples []st.SampledValueInput
	for i := 0; i < 10; i++ {
		samples = append(samples, st.SampledValueInput{
			Value: sampleValue,
		})
	}

	input := st.MeterValueInput{
		Timestamp:    sampleTimestamp,
		SampledValue: samples,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewMeterValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMeterValuesReq_SingleValue(b *testing.B) {
	b.ReportAllocs()

	reqInput := mv.ReqInput{
		ConnectorId: 1,
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: sampleTimestamp,
				SampledValue: []st.SampledValueInput{
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

	var samples []st.SampledValueInput
	for i := 0; i < 10; i++ {
		samples = append(samples, st.SampledValueInput{
			Value: sampleValue,
		})
	}

	reqInput := mv.ReqInput{
		ConnectorId: 1,
		MeterValue: []st.MeterValueInput{
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
		if _, err := st.NewCiString25Type(sample); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewCiString50Type(b *testing.B) {
	const sample = "RFID-ABC1234567890123456789012345678901234567890"

	for i := 0; i < b.N; i++ {
		if _, err := st.NewCiString50Type(sample); err != nil {
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
		if _, err := st.NewCiString255Type(string(long)); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewIdToken(b *testing.B) {
	b.ReportAllocs()

	ci, _ := st.NewCiString20Type(ciStringSample)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = st.NewIdToken(ci)
	}
}

func BenchmarkNewIdTagInfo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted); err != nil {
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
		TransactionData: []mt.MeterValueInput{
			{
				Timestamp: sampleTimestamp,
				SampledValue: []mt.SampledValueInput{
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
