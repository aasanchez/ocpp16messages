//go:build bench

package benchmark

import (
	"testing"

	gct "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	scpt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

var (
	sinkInteger      *st.Integer
	sinkDateTime     *st.DateTime
	sinkFloat64      *float64
	sinkCiString500  *st.CiString500Type
	sinkScheduleList []st.ChargingSchedulePeriod
)

func BenchmarkChargingScheduleDurationGetter_Copy(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}
	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          nil,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkInteger = schedule.Duration()
	}
}

func BenchmarkChargingScheduleStartScheduleGetter_Copy(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	startSchedule := sampleTimestamp
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          &startSchedule,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkDateTime = schedule.StartSchedule()
	}
}

func BenchmarkChargingScheduleMinChargingRateGetter_Copy(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	minChargingRate := 1.0
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          nil,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkFloat64 = schedule.MinChargingRate()
	}
}

func BenchmarkChargingScheduleChargingSchedulePeriodGetter_Copy(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: 0, Limit: 16},
		{StartPeriod: 10, Limit: 8},
		{StartPeriod: 20, Limit: 6},
		{StartPeriod: 30, Limit: 4},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          nil,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkScheduleList = schedule.ChargingSchedulePeriod()
	}
}

func BenchmarkGetConfigurationKeyValueValueGetter_Copy(b *testing.B) {
	b.ReportAllocs()

	value := "60"
	keyValue, err := gct.NewKeyValue(gct.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkCiString500 = keyValue.Value()
	}
}

func BenchmarkSetChargingProfileTransactionIdGetter_Copy(b *testing.B) {
	b.ReportAllocs()

	transactionId := 1

	duration := 60
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}
	scheduleInput := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          nil,
	}

	profile, err := scpt.NewChargingProfile(scpt.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          &transactionId,
		StackLevel:             0,
		ChargingProfilePurpose: st.TxProfile.String(),
		ChargingProfileKind:    scpt.ChargingProfileKindAbsolute.String(),
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       scheduleInput,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkInteger = profile.TransactionId()
	}
}
