//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

// maxFuzzStringLen is kept for backward compatibility;
// new tests should use maxFuzzLen from fuzz_helpers_test.go.
const maxFuzzStringLen = maxFuzzLen

func FuzzNewChargingSchedule(f *testing.F) {
	f.Add(false, 0, "W", true, 0, 0.0, false, 0, false, 0.0, false, "")
	f.Add(true, 60, "A", true, 0, 10.5, true, 3, true, 0.1, true, "2025-01-15T10:30:00Z")
	f.Add(true, -1, "W", true, 0, 0.0, false, 0, false, 0.0, false, "")
	f.Add(false, 0, "invalid-unit", true, 0, 0.0, false, 0, false, 0.0, false, "")
	f.Add(false, 0, "W", false, 0, 0.0, false, 0, false, 0.0, false, "")

	f.Fuzz(func(
		t *testing.T,
		hasDuration bool,
		duration int,
		chargingRateUnit string,
		hasPeriod bool,
		startPeriod int,
		limit float64,
		hasNumberPhases bool,
		numberPhases int,
		hasMinChargingRate bool,
		minChargingRate float64,
		hasStartSchedule bool,
		startSchedule string,
	) {
		if len(chargingRateUnit) > maxFuzzStringLen || len(startSchedule) > maxFuzzStringLen {
			t.Skip()
		}

		var durationPtr *int
		if hasDuration {
			durationPtr = &duration
		}

		var minChargingRatePtr *float64
		if hasMinChargingRate {
			minChargingRatePtr = &minChargingRate
		}

		var startSchedulePtr *string
		if hasStartSchedule {
			startSchedulePtr = &startSchedule
		}

		var numberPhasesPtr *int
		if hasNumberPhases {
			numberPhasesPtr = &numberPhases
		}

		var periods []types.ChargingSchedulePeriodInput
		if hasPeriod {
			if math.IsNaN(limit) || math.IsInf(limit, 0) {
				t.Skip()
			}

			periods = []types.ChargingSchedulePeriodInput{
				{
					StartPeriod:  startPeriod,
					Limit:        limit,
					NumberPhases: numberPhasesPtr,
				},
			}
		}

		schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
			Duration:               durationPtr,
			ChargingRateUnit:       chargingRateUnit,
			ChargingSchedulePeriod: periods,
			MinChargingRate:        minChargingRatePtr,
			StartSchedule:          startSchedulePtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if len(schedule.ChargingSchedulePeriod()) == 0 {
			t.Fatal("ChargingSchedulePeriod() is empty, want at least one")
		}

		if !schedule.ChargingRateUnit().IsValid() {
			t.Fatalf("ChargingRateUnit() = %q, want valid", schedule.ChargingRateUnit())
		}

		if hasDuration {
			if schedule.Duration() == nil {
				t.Fatal("Duration() = nil, want non-nil")
			}
		} else if schedule.Duration() != nil {
			t.Fatal("Duration() != nil, want nil")
		}

		if hasMinChargingRate {
			if schedule.MinChargingRate() == nil {
				t.Fatal("MinChargingRate() = nil, want non-nil")
			}
		} else if schedule.MinChargingRate() != nil {
			t.Fatal("MinChargingRate() != nil, want nil")
		}

		if hasStartSchedule {
			if schedule.StartSchedule() == nil {
				t.Fatal("StartSchedule() = nil, want non-nil")
			}
			if schedule.StartSchedule().Value().Location() != time.UTC {
				t.Fatalf(
					"StartSchedule location = %v, want UTC",
					schedule.StartSchedule().Value().Location(),
				)
			}
		} else if schedule.StartSchedule() != nil {
			t.Fatal("StartSchedule() != nil, want nil")
		}
	})
}
