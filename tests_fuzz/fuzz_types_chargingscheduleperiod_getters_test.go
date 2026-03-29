//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzChargingSchedulePeriodGetterConsistency(
	f *testing.F,
) {
	f.Add(0, 10.0, true, 3)
	f.Add(0, 0.0, false, 0)
	f.Add(100, 32.5, true, 1)

	f.Fuzz(func(
		t *testing.T,
		startPeriod int, limit float64,
		hasNumPhases bool, numPhases int,
	) {
		if math.IsNaN(limit) || math.IsInf(limit, 0) {
			t.Skip()
		}

		var numPhasesPtr *int
		if hasNumPhases {
			numPhasesPtr = &numPhases
		}

		period, err := types.NewChargingSchedulePeriod(
			types.ChargingSchedulePeriodInput{
				StartPeriod:  startPeriod,
				Limit:        limit,
				NumberPhases: numPhasesPtr,
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

		// StartPeriod
		if int(period.StartPeriod().Value()) != startPeriod {
			t.Fatalf(
				"StartPeriod() = %d, want %d",
				period.StartPeriod().Value(), startPeriod,
			)
		}

		// Limit
		if period.Limit() != limit {
			t.Fatalf(
				"Limit() = %f, want %f",
				period.Limit(), limit,
			)
		}

		// NumberPhases
		if hasNumPhases {
			if period.NumberPhases() == nil {
				t.Fatal(
					"NumberPhases() = nil, want non-nil",
				)
			}

			if int(period.NumberPhases().Value()) !=
				numPhases {
				t.Fatalf(
					"NumberPhases() = %d, want %d",
					period.NumberPhases().Value(),
					numPhases,
				)
			}
		} else if period.NumberPhases() != nil {
			t.Fatal("NumberPhases() != nil, want nil")
		}

		// String() determinism on StartPeriod
		s1 := period.StartPeriod().String()
		s2 := period.StartPeriod().String()

		if s1 != s2 {
			t.Fatal(
				"StartPeriod().String() not deterministic",
			)
		}
	})
}

func FuzzChargingSchedulePeriodPointerIsolation(
	f *testing.F,
) {
	f.Add(0, 10.0, 3)

	f.Fuzz(func(
		t *testing.T,
		startPeriod int, limit float64, numPhases int,
	) {
		if math.IsNaN(limit) || math.IsInf(limit, 0) {
			t.Skip()
		}

		numPhasesPtr := &numPhases

		period, err := types.NewChargingSchedulePeriod(
			types.ChargingSchedulePeriodInput{
				StartPeriod:  startPeriod,
				Limit:        limit,
				NumberPhases: numPhasesPtr,
			},
		)
		if err != nil {
			return
		}

		// NumberPhases pointer isolation
		np := period.NumberPhases()
		if np == nil {
			return
		}

		orig := np.Value()
		np2 := period.NumberPhases()

		if np2.Value() != orig {
			t.Fatal("NumberPhases changed between calls")
		}
	})
}
