package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testChargingProfileId = 1
	testStackLevel        = 0
	testPurposeTxProfile  = "TxProfile"
	testKindAbsolute      = "Absolute"
	testRecurrencyDaily   = "Daily"
	testValidFrom         = "2025-01-01T00:00:00Z"
	testValidTo           = "2025-12-31T23:59:59Z"
	testStartPeriod       = 0
	testLimit             = 32.0
	testChargingRateUnitA = "A"
)

func createValidChargingProfile(t *testing.T) types.ChargingProfile {
	t.Helper()

	transactionId := 123
	recurrencyKind := testRecurrencyDaily
	validFrom := testValidFrom
	validTo := testValidTo

	profile, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          &transactionId,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         &recurrencyKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule: st.ChargingScheduleInput{
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{StartPeriod: testStartPeriod, Limit: testLimit},
			},
		},
	})
	if err != nil {
		t.Fatalf("Failed to create valid ChargingProfile: %v", err)
	}

	return profile
}

func TestChargingProfile_TransactionId(t *testing.T) {
	t.Parallel()

	profile := createValidChargingProfile(t)
	transactionId := profile.TransactionId()

	if transactionId == nil {
		t.Error("ChargingProfile.TransactionId() = nil, want non-nil")
	}
}

func TestChargingProfile_TransactionId_WhenNil(t *testing.T) {
	t.Parallel()

	profile, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		ChargingSchedule: st.ChargingScheduleInput{
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{StartPeriod: testStartPeriod, Limit: testLimit},
			},
		},
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingProfile: %v", err)
	}

	if profile.TransactionId() != nil {
		t.Error("ChargingProfile.TransactionId() = non-nil, want nil")
	}
}

func TestChargingProfile_RecurrencyKind(t *testing.T) {
	t.Parallel()

	profile := createValidChargingProfile(t)
	recurrencyKind := profile.RecurrencyKind()

	if recurrencyKind == nil {
		t.Error("ChargingProfile.RecurrencyKind() = nil, want non-nil")
	}
}

func TestChargingProfile_ValidFrom(t *testing.T) {
	t.Parallel()

	profile := createValidChargingProfile(t)
	validFrom := profile.ValidFrom()

	if validFrom == nil {
		t.Error("ChargingProfile.ValidFrom() = nil, want non-nil")
	}
}

func TestChargingProfile_ValidTo(t *testing.T) {
	t.Parallel()

	profile := createValidChargingProfile(t)
	validTo := profile.ValidTo()

	if validTo == nil {
		t.Error("ChargingProfile.ValidTo() = nil, want non-nil")
	}
}

func TestChargingProfile_ChargingSchedule(t *testing.T) {
	t.Parallel()

	profile := createValidChargingProfile(t)
	schedule := profile.ChargingSchedule()

	if schedule.ChargingRateUnit() != st.ChargingRateUnitAmperes {
		t.Errorf(
			"ChargingProfile.ChargingSchedule().ChargingRateUnit() = %v, want %v",
			schedule.ChargingRateUnit(),
			st.ChargingRateUnitAmperes,
		)
	}
}
