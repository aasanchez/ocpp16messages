//go:build bench

package benchmark

import (
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func BenchmarkNewChargingSchedule(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	startSchedule := sampleTimestamp
	minChargingRate := 0.0

	periods := []types.ChargingSchedulePeriodInput{
		{StartPeriod: 0, Limit: 16},
		{StartPeriod: 10, Limit: 8},
	}

	input := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &startSchedule,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewChargingSchedule(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkChargingScheduleChargingSchedulePeriodGetter(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	startSchedule := sampleTimestamp
	minChargingRate := 0.0

	periods := []types.ChargingSchedulePeriodInput{
		{StartPeriod: 0, Limit: 16},
		{StartPeriod: 10, Limit: 8},
		{StartPeriod: 20, Limit: 6},
		{StartPeriod: 30, Limit: 4},
	}

	input := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &startSchedule,
	}

	schedule, err := types.NewChargingSchedule(input)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		values := schedule.ChargingSchedulePeriod()
		_ = values[0].Limit()
	}
}
