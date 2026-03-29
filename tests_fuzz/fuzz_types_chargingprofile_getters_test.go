//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzChargingProfileGetterConsistency(f *testing.F) {
	f.Add(
		1, 0, "ChargePointMaxProfile", "Absolute",
		true, 100, false, "", false, "",
		true, 60, "W", 0, 1.0, false, 0,
	)
	f.Add(
		0, 1, "TxDefaultProfile", "Recurring",
		false, 0, true, "2025-01-15T10:30:00Z",
		true, "2025-12-31T23:59:59Z",
		true, 0, "A", 0, 0.0, true, 3,
	)

	f.Fuzz(func(
		t *testing.T,
		profileId int,
		stackLevel int,
		purpose string,
		kind string,
		hasTxId bool, txId int,
		hasValidFrom bool, validFrom string,
		hasValidTo bool, validTo string,
		hasDuration bool, duration int,
		rateUnit string,
		startPeriod int, limit float64,
		hasNumPhases bool, numPhases int,
	) {
		if len(purpose) > maxFuzzLen ||
			len(kind) > maxFuzzLen ||
			len(validFrom) > maxFuzzLen ||
			len(validTo) > maxFuzzLen ||
			len(rateUnit) > maxFuzzLen {
			t.Skip()
		}

		if math.IsNaN(limit) || math.IsInf(limit, 0) {
			t.Skip()
		}

		var txIdPtr *int
		if hasTxId {
			txIdPtr = &txId
		}

		var validFromPtr, validToPtr *string
		if hasValidFrom {
			validFromPtr = &validFrom
		}

		if hasValidTo {
			validToPtr = &validTo
		}

		var durationPtr *int
		if hasDuration {
			durationPtr = &duration
		}

		var numPhasesPtr *int
		if hasNumPhases {
			numPhasesPtr = &numPhases
		}

		profile, err := types.NewChargingProfile(
			types.ChargingProfileInput{
				ChargingProfileId:      profileId,
				TransactionId:          txIdPtr,
				StackLevel:             stackLevel,
				ChargingProfilePurpose: purpose,
				ChargingProfileKind:    kind,
				ValidFrom:              validFromPtr,
				ValidTo:                validToPtr,
				ChargingSchedule: types.ChargingScheduleInput{
					Duration:         durationPtr,
					ChargingRateUnit: rateUnit,
					ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
						{
							StartPeriod:  startPeriod,
							Limit:        limit,
							NumberPhases: numPhasesPtr,
						},
					},
				},
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

		// Getter consistency: every getter must reflect input
		if int(profile.ChargingProfileId().Value()) !=
			profileId {
			t.Fatalf(
				"ChargingProfileId() = %d, want %d",
				profile.ChargingProfileId().Value(),
				profileId,
			)
		}

		if int(profile.StackLevel().Value()) != stackLevel {
			t.Fatalf(
				"StackLevel() = %d, want %d",
				profile.StackLevel().Value(), stackLevel,
			)
		}

		if profile.ChargingProfilePurpose().String() !=
			purpose {
			t.Fatalf(
				"Purpose() = %q, want %q",
				profile.ChargingProfilePurpose().String(),
				purpose,
			)
		}

		if !profile.ChargingProfilePurpose().IsValid() {
			t.Fatal("ChargingProfilePurpose() not valid")
		}

		if profile.ChargingProfileKind().String() != kind {
			t.Fatalf(
				"Kind() = %q, want %q",
				profile.ChargingProfileKind().String(),
				kind,
			)
		}

		if !profile.ChargingProfileKind().IsValid() {
			t.Fatal("ChargingProfileKind() not valid")
		}

		// TransactionId
		if hasTxId {
			if profile.TransactionId() == nil {
				t.Fatal("TransactionId() = nil, want non-nil")
			}

			if int(profile.TransactionId().Value()) != txId {
				t.Fatalf(
					"TransactionId() = %d, want %d",
					profile.TransactionId().Value(), txId,
				)
			}
		} else if profile.TransactionId() != nil {
			t.Fatal("TransactionId() != nil, want nil")
		}

		// ValidFrom
		if hasValidFrom {
			if profile.ValidFrom() == nil {
				t.Fatal("ValidFrom() = nil, want non-nil")
			}

			if profile.ValidFrom().Value().Location() !=
				time.UTC {
				t.Fatal("ValidFrom() not in UTC")
			}
		} else if profile.ValidFrom() != nil {
			t.Fatal("ValidFrom() != nil, want nil")
		}

		// ValidTo
		if hasValidTo {
			if profile.ValidTo() == nil {
				t.Fatal("ValidTo() = nil, want non-nil")
			}

			if profile.ValidTo().Value().Location() !=
				time.UTC {
				t.Fatal("ValidTo() not in UTC")
			}
		} else if profile.ValidTo() != nil {
			t.Fatal("ValidTo() != nil, want nil")
		}

		// ChargingSchedule is always present
		sched := profile.ChargingSchedule()
		if !sched.ChargingRateUnit().IsValid() {
			t.Fatal("ChargingRateUnit() not valid")
		}

		if len(sched.ChargingSchedulePeriod()) == 0 {
			t.Fatal(
				"ChargingSchedulePeriod() empty after success",
			)
		}
	})
}

func FuzzChargingProfilePointerIsolation(f *testing.F) {
	f.Add(
		1, 0, "ChargePointMaxProfile", "Absolute",
		100, "2025-01-15T10:30:00Z", "2025-12-31T23:59:59Z",
	)

	f.Fuzz(func(
		t *testing.T,
		profileId, stackLevel int,
		purpose, kind string,
		txId int,
		validFrom, validTo string,
	) {
		if len(purpose) > maxFuzzLen ||
			len(kind) > maxFuzzLen ||
			len(validFrom) > maxFuzzLen ||
			len(validTo) > maxFuzzLen {
			t.Skip()
		}

		txIdPtr := &txId
		validFromPtr := &validFrom
		validToPtr := &validTo

		profile, err := types.NewChargingProfile(
			types.ChargingProfileInput{
				ChargingProfileId:      profileId,
				TransactionId:          txIdPtr,
				StackLevel:             stackLevel,
				ChargingProfilePurpose: purpose,
				ChargingProfileKind:    kind,
				ValidFrom:              validFromPtr,
				ValidTo:                validToPtr,
				ChargingSchedule: types.ChargingScheduleInput{
					ChargingRateUnit: "W",
					ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
						{StartPeriod: 0, Limit: 10.0},
					},
				},
			},
		)
		if err != nil {
			return
		}

		// TransactionId pointer isolation
		if ptr := profile.TransactionId(); ptr != nil {
			original := ptr.Value()
			ptr2 := profile.TransactionId()

			if ptr2.Value() != original {
				t.Fatal(
					"TransactionId changed between calls",
				)
			}
		}

		// ValidFrom pointer isolation
		if ptr := profile.ValidFrom(); ptr != nil {
			original := ptr.String()
			ptr2 := profile.ValidFrom()

			if ptr2.String() != original {
				t.Fatal("ValidFrom changed between calls")
			}
		}

		// ValidTo pointer isolation
		if ptr := profile.ValidTo(); ptr != nil {
			original := ptr.String()
			ptr2 := profile.ValidTo()

			if ptr2.String() != original {
				t.Fatal("ValidTo changed between calls")
			}
		}
	})
}
