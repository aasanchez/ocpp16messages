//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	spt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewChargingProfile(f *testing.F) {
	f.Add(
		1,
		0,
		"ChargePointMaxProfile",
		"Absolute",
		false,
		0,
		false,
		"",
		false,
		"",
		false,
		"",
	)
	f.Add(
		-1,
		0,
		"ChargePointMaxProfile",
		"Absolute",
		false,
		0,
		false,
		"",
		false,
		"",
		false,
		"",
	)
	f.Add(
		1,
		-1,
		"ChargePointMaxProfile",
		"Absolute",
		false,
		0,
		false,
		"",
		false,
		"",
		false,
		"",
	)
	f.Add(
		1,
		0,
		"invalid-purpose",
		"Absolute",
		false,
		0,
		false,
		"",
		false,
		"",
		false,
		"",
	)

	f.Fuzz(func(
		t *testing.T,
		chargingProfileId int,
		stackLevel int,
		purpose string,
		kind string,
		hasTransactionId bool,
		transactionId int,
		hasRecurrencyKind bool,
		recurrencyKind string,
		hasValidFrom bool,
		validFrom string,
		hasValidTo bool,
		validTo string,
	) {
		if len(purpose) > maxFuzzStringLen ||
			len(kind) > maxFuzzStringLen ||
			len(recurrencyKind) > maxFuzzStringLen ||
			len(validFrom) > maxFuzzStringLen ||
			len(validTo) > maxFuzzStringLen {
			t.Skip()
		}

		var transactionIdPtr *int
		if hasTransactionId {
			transactionIdPtr = &transactionId
		}

		var recurrencyKindPtr *string
		if hasRecurrencyKind {
			recurrencyKindPtr = &recurrencyKind
		}

		var validFromPtr *string
		if hasValidFrom {
			validFromPtr = &validFrom
		}

		var validToPtr *string
		if hasValidTo {
			validToPtr = &validTo
		}

		profile, err := spt.NewChargingProfile(spt.ChargingProfileInput{
			ChargingProfileId:      chargingProfileId,
			TransactionId:          transactionIdPtr,
			StackLevel:             stackLevel,
			ChargingProfilePurpose: purpose,
			ChargingProfileKind:    kind,
			RecurrencyKind:         recurrencyKindPtr,
			ValidFrom:              validFromPtr,
			ValidTo:                validToPtr,
			ChargingSchedule: st.ChargingScheduleInput{
				ChargingRateUnit: st.ChargingRateUnitWatts.String(),
				ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
					{
						StartPeriod:  0,
						Limit:        0,
						NumberPhases: nil,
					},
				},
			},
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if chargingProfileId < 0 || chargingProfileId > math.MaxUint16 {
			t.Fatalf(
				"NewChargingProfile succeeded with chargingProfileId=%d",
				chargingProfileId,
			)
		}

		if got := profile.ChargingProfileId().Value(); got != uint16(chargingProfileId) {
			t.Fatalf("ChargingProfileId() = %d, want %d", got, chargingProfileId)
		}

		if !profile.ChargingProfilePurpose().IsValid() {
			t.Fatalf(
				"ChargingProfilePurpose() = %q, want valid",
				profile.ChargingProfilePurpose().String(),
			)
		}

		if !profile.ChargingProfileKind().IsValid() {
			t.Fatalf(
				"ChargingProfileKind() = %q, want valid",
				profile.ChargingProfileKind().String(),
			)
		}

		if stackLevel < 0 || stackLevel > math.MaxUint16 {
			t.Fatalf("NewChargingProfile succeeded with stackLevel=%d", stackLevel)
		}

		if got := profile.StackLevel().Value(); got != uint16(stackLevel) {
			t.Fatalf("StackLevel() = %d, want %d", got, stackLevel)
		}

		if hasTransactionId {
			if profile.TransactionId() == nil {
				t.Fatal("TransactionId() = nil, want non-nil")
			}

			if transactionId < 0 || transactionId > math.MaxUint16 {
				t.Fatalf(
					"NewChargingProfile succeeded with transactionId=%d",
					transactionId,
				)
			}

			if got := profile.TransactionId().Value(); got != uint16(transactionId) {
				t.Fatalf("TransactionId() = %d, want %d", got, transactionId)
			}
		} else if profile.TransactionId() != nil {
			t.Fatal("TransactionId() != nil, want nil")
		}

		if hasRecurrencyKind {
			if profile.RecurrencyKind() == nil {
				t.Fatal("RecurrencyKind() = nil, want non-nil")
			}

			if !profile.RecurrencyKind().IsValid() {
				t.Fatalf(
					"RecurrencyKind() = %q, want valid",
					profile.RecurrencyKind().String(),
				)
			}
		} else if profile.RecurrencyKind() != nil {
			t.Fatal("RecurrencyKind() != nil, want nil")
		}

		if hasValidFrom {
			if profile.ValidFrom() == nil {
				t.Fatal("ValidFrom() = nil, want non-nil")
			}
			if profile.ValidFrom().Value().Location() != time.UTC {
				t.Fatalf(
					"ValidFrom location = %v, want UTC",
					profile.ValidFrom().Value().Location(),
				)
			}
		} else if profile.ValidFrom() != nil {
			t.Fatal("ValidFrom() != nil, want nil")
		}

		if hasValidTo {
			if profile.ValidTo() == nil {
				t.Fatal("ValidTo() = nil, want non-nil")
			}
			if profile.ValidTo().Value().Location() != time.UTC {
				t.Fatalf(
					"ValidTo location = %v, want UTC",
					profile.ValidTo().Value().Location(),
				)
			}
		} else if profile.ValidTo() != nil {
			t.Fatal("ValidTo() != nil, want nil")
		}

		if len(profile.ChargingSchedule().ChargingSchedulePeriod()) == 0 {
			t.Fatal("ChargingSchedulePeriod() is empty, want at least one")
		}
	})
}
