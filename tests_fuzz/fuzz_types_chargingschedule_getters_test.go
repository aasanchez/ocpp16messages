//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzChargingScheduleGetterConsistency(f *testing.F) {
	f.Add(
		true, 60, "W", 0, 10.5,
		true, 3, true, 0.1, true, "2025-01-15T10:30:00Z",
	)
	f.Add(
		false, 0, "A", 0, 0.0,
		false, 0, false, 0.0, false, "",
	)

	f.Fuzz(func(
		t *testing.T,
		hasDuration bool, duration int,
		rateUnit string,
		startPeriod int, limit float64,
		hasNumPhases bool, numPhases int,
		hasMinRate bool, minRate float64,
		hasStartSched bool, startSched string,
	) {
		if len(rateUnit) > maxFuzzLen ||
			len(startSched) > maxFuzzLen {
			t.Skip()
		}

		if math.IsNaN(limit) || math.IsInf(limit, 0) ||
			math.IsNaN(minRate) || math.IsInf(minRate, 0) {
			t.Skip()
		}

		var durationPtr *int
		if hasDuration {
			durationPtr = &duration
		}

		var numPhasesPtr *int
		if hasNumPhases {
			numPhasesPtr = &numPhases
		}

		var minRatePtr *float64
		if hasMinRate {
			minRatePtr = &minRate
		}

		var startSchedPtr *string
		if hasStartSched {
			startSchedPtr = &startSched
		}

		sched, err := types.NewChargingSchedule(
			types.ChargingScheduleInput{
				Duration:         durationPtr,
				ChargingRateUnit: rateUnit,
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod:  startPeriod,
						Limit:        limit,
						NumberPhases: numPhasesPtr,
					},
				},
				MinChargingRate: minRatePtr,
				StartSchedule:   startSchedPtr,
			},
		)
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) &&
				!errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want sentinel",
					err,
				)
			}

			return
		}

		// ChargingRateUnit
		if !sched.ChargingRateUnit().IsValid() {
			t.Fatal("ChargingRateUnit() not valid")
		}

		if sched.ChargingRateUnit().String() != rateUnit {
			t.Fatalf(
				"ChargingRateUnit() = %q, want %q",
				sched.ChargingRateUnit().String(), rateUnit,
			)
		}

		// String() determinism
		u1 := sched.ChargingRateUnit().String()
		u2 := sched.ChargingRateUnit().String()

		if u1 != u2 {
			t.Fatal(
				"ChargingRateUnit String() not deterministic",
			)
		}

		// Periods
		periods := sched.ChargingSchedulePeriod()
		if len(periods) == 0 {
			t.Fatal("ChargingSchedulePeriod() empty")
		}

		// Duration
		if hasDuration {
			if sched.Duration() == nil {
				t.Fatal("Duration() = nil, want non-nil")
			}

			if int(sched.Duration().Value()) != duration {
				t.Fatalf(
					"Duration() = %d, want %d",
					sched.Duration().Value(), duration,
				)
			}
		} else if sched.Duration() != nil {
			t.Fatal("Duration() != nil, want nil")
		}

		// MinChargingRate
		if hasMinRate {
			if sched.MinChargingRate() == nil {
				t.Fatal(
					"MinChargingRate() = nil, want non-nil",
				)
			}

			if *sched.MinChargingRate() != minRate {
				t.Fatalf(
					"MinChargingRate() = %f, want %f",
					*sched.MinChargingRate(), minRate,
				)
			}
		} else if sched.MinChargingRate() != nil {
			t.Fatal("MinChargingRate() != nil, want nil")
		}

		// StartSchedule
		if hasStartSched {
			if sched.StartSchedule() == nil {
				t.Fatal(
					"StartSchedule() = nil, want non-nil",
				)
			}

			if sched.StartSchedule().Value().Location() !=
				time.UTC {
				t.Fatal("StartSchedule() not in UTC")
			}
		} else if sched.StartSchedule() != nil {
			t.Fatal("StartSchedule() != nil, want nil")
		}
	})
}

func FuzzChargingSchedulePointerIsolation(f *testing.F) {
	f.Add(60, "W", 0, 10.0, 3, 0.1, "2025-01-15T10:30:00Z")

	f.Fuzz(func(
		t *testing.T,
		duration int, rateUnit string,
		startPeriod int, limit float64,
		numPhases int, minRate float64,
		startSched string,
	) {
		if len(rateUnit) > maxFuzzLen ||
			len(startSched) > maxFuzzLen {
			t.Skip()
		}

		if math.IsNaN(limit) || math.IsInf(limit, 0) ||
			math.IsNaN(minRate) || math.IsInf(minRate, 0) {
			t.Skip()
		}

		durationPtr := &duration
		numPhasesPtr := &numPhases
		minRatePtr := &minRate
		startSchedPtr := &startSched

		sched, err := types.NewChargingSchedule(
			types.ChargingScheduleInput{
				Duration:         durationPtr,
				ChargingRateUnit: rateUnit,
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod:  startPeriod,
						Limit:        limit,
						NumberPhases: numPhasesPtr,
					},
				},
				MinChargingRate: minRatePtr,
				StartSchedule:   startSchedPtr,
			},
		)
		if err != nil {
			return
		}

		// Duration pointer isolation
		if d := sched.Duration(); d != nil {
			orig := d.Value()
			d2 := sched.Duration()

			if d2.Value() != orig {
				t.Fatal("Duration changed between calls")
			}
		}

		// StartSchedule pointer isolation
		if ss := sched.StartSchedule(); ss != nil {
			orig := ss.String()
			ss2 := sched.StartSchedule()

			if ss2.String() != orig {
				t.Fatal(
					"StartSchedule changed between calls",
				)
			}
		}

		// MinChargingRate pointer isolation
		if mr := sched.MinChargingRate(); mr != nil {
			orig := *mr
			*mr = orig + 999.0
			mr2 := sched.MinChargingRate()

			if *mr2 != orig {
				t.Fatal(
					"MinChargingRate mutation leaked",
				)
			}
		}

		// Slice copy isolation
		periods := sched.ChargingSchedulePeriod()
		origLen := len(periods)
		periods = append(
			periods,
			types.ChargingSchedulePeriod{},
		)

		if len(sched.ChargingSchedulePeriod()) != origLen {
			t.Fatalf(
				"slice append leaked: got %d, want %d",
				len(sched.ChargingSchedulePeriod()),
				origLen,
			)
		}
	})
}
