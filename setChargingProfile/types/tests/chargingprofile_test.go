package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testChargingProfileId    = 1
	testStackLevel           = 0
	testPurposeTxProfile     = "TxProfile"
	testKindAbsolute         = "Absolute"
	testRecurrencyDaily      = "Daily"
	testValidFrom            = "2025-01-01T00:00:00Z"
	testValidTo              = "2025-12-31T23:59:59Z"
	testStartPeriod          = 0
	testLimit                = 32.0
	testChargingRateUnitA    = "A"
	testTransactionId        = 123
	testInvalidTransactionId = -1
	testOverflowStackLevel   = 65536
	errWantError             = "NewChargingProfile() error = nil, want error"
	errFailedCreateProfile   = "Failed to create ChargingProfile: %v"
)

func createValidChargingProfile(t *testing.T) types.ChargingProfile {
	t.Helper()

	transactionId := testTransactionId
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
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err != nil {
		t.Fatalf(errFailedCreateProfile, err)
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
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err != nil {
		t.Fatalf(errFailedCreateProfile, err)
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
			"ChargingProfile.ChargingSchedule().ChargingRateUnit() = %v, "+
				"want %v",
			schedule.ChargingRateUnit(),
			st.ChargingRateUnitAmperes,
		)
	}
}

func TestChargingProfile_RecurrencyKind_WhenNil(t *testing.T) {
	t.Parallel()

	profile, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err != nil {
		t.Fatalf(errFailedCreateProfile, err)
	}

	if profile.RecurrencyKind() != nil {
		t.Error("ChargingProfile.RecurrencyKind() = non-nil, want nil")
	}
}

func TestChargingProfile_ValidFrom_WhenNil(t *testing.T) {
	t.Parallel()

	profile, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err != nil {
		t.Fatalf(errFailedCreateProfile, err)
	}

	if profile.ValidFrom() != nil {
		t.Error("ChargingProfile.ValidFrom() = non-nil, want nil")
	}
}

func TestChargingProfile_ValidTo_WhenNil(t *testing.T) {
	t.Parallel()

	profile, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err != nil {
		t.Fatalf(errFailedCreateProfile, err)
	}

	if profile.ValidTo() != nil {
		t.Error("ChargingProfile.ValidTo() = non-nil, want nil")
	}
}

func TestNewChargingProfile_InvalidChargingProfileId(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      -1,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_InvalidTransactionId(t *testing.T) {
	t.Parallel()

	transactionId := testInvalidTransactionId

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          &transactionId,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_InvalidStackLevel(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             -1,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_StackLevelExceedsUint16(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testOverflowStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error("NewChargingProfile() error = nil, want overflow error")
	}
}

func TestNewChargingProfile_InvalidPurpose(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: "InvalidPurpose",
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_InvalidKind(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    "InvalidKind",
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_InvalidRecurrencyKind(t *testing.T) {
	t.Parallel()

	recurrencyKind := "InvalidRecurrency"

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         &recurrencyKind,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_InvalidValidFrom(t *testing.T) {
	t.Parallel()

	invalidTime := "not-a-valid-time"

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              &invalidTime,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_InvalidValidTo(t *testing.T) {
	t.Parallel()

	invalidTime := "not-a-valid-time"

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                &invalidTime,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testChargingRateUnitA,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod:  testStartPeriod,
					Limit:        testLimit,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}

func TestNewChargingProfile_InvalidChargingSchedule(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      testChargingProfileId,
		TransactionId:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurposeTxProfile,
		ChargingProfileKind:    testKindAbsolute,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: st.ChargingScheduleInput{
			Duration:               nil,
			ChargingRateUnit:       "InvalidUnit",
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{},
			MinChargingRate:        nil,
			StartSchedule:          nil,
		},
	})
	if err == nil {
		t.Error(errWantError)
	}
}
