//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewChargingSchedulePeriod(f *testing.F) {
	f.Add(0, 0.0, false, 0)
	f.Add(1, 10.5, true, 1)
	f.Add(1, 10.5, true, 3)
	f.Add(-1, 0.0, false, 0)
	f.Add(0, -0.1, false, 0)
	f.Add(0, 0.0, true, 0)
	f.Add(0, 0.0, true, 4)

	f.Fuzz(func(
		t *testing.T,
		startPeriod int,
		limit float64,
		hasNumberPhases bool,
		numberPhases int,
	) {
		if math.IsNaN(limit) || math.IsInf(limit, 0) {
			t.Skip()
		}

		var numberPhasesPtr *int
		if hasNumberPhases {
			numberPhasesPtr = &numberPhases
		}

		period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
			StartPeriod:  startPeriod,
			Limit:        limit,
			NumberPhases: numberPhasesPtr,
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

		if startPeriod < 0 || startPeriod > math.MaxUint16 {
			t.Fatalf(
				"NewChargingSchedulePeriod succeeded with startPeriod=%d",
				startPeriod,
			)
		}

		if limit < 0 {
			t.Fatalf("NewChargingSchedulePeriod succeeded with limit=%v", limit)
		}

		if hasNumberPhases && (numberPhases < 1 || numberPhases > 3) {
			t.Fatalf(
				"NewChargingSchedulePeriod succeeded with numberPhases=%d",
				numberPhases,
			)
		}

		if got := period.StartPeriod().Value(); got != uint16(startPeriod) {
			t.Fatalf("StartPeriod() = %d, want %d", got, startPeriod)
		}

		if period.Limit() != limit {
			t.Fatalf("Limit() = %v, want %v", period.Limit(), limit)
		}

		if !hasNumberPhases {
			if period.NumberPhases() != nil {
				t.Fatal("NumberPhases() != nil, want nil")
			}

			return
		}

		if period.NumberPhases() == nil {
			t.Fatal("NumberPhases() = nil, want non-nil")
		}

		if got := period.NumberPhases().Value(); got != uint16(numberPhases) {
			t.Fatalf("NumberPhases() = %d, want %d", got, numberPhases)
		}
	})
}
