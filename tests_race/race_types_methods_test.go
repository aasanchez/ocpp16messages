//go:build race

package race

import (
	"fmt"
	"testing"

	gct "github.com/aasanchez/ocpp16messages/getconfiguration/types"
	gllt "github.com/aasanchez/ocpp16messages/getlocallistversion/types"
	scpt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestRace_IdTagInfoWithAndString(t *testing.T) {
	t.Parallel()

	base, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf("NewIdTagInfo: %v", err)
	}

	expiry, err := st.NewDateTime("2025-01-02T15:00:00Z")
	if err != nil {
		t.Fatalf("NewDateTime: %v", err)
	}

	ciTag, err := st.NewCiString20Type("PARENT-1")
	if err != nil {
		t.Fatalf("NewCiString20Type: %v", err)
	}
	parent := st.NewIdToken(ciTag)

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_ = base.String()

		info := base.WithExpiryDate(expiry).WithParentIdTag(parent)
		_ = info.String()

		return nil
	})
}

func TestRace_ChargingScheduleGetters(t *testing.T) {
	t.Parallel()

	duration := 60
	startSchedule := "2025-01-02T15:00:00Z"
	minChargingRate := 0.0

	periods := []st.ChargingSchedulePeriodInput{
		{
			StartPeriod:  0,
			Limit:        16,
			NumberPhases: nil,
		},
	}

	input := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &startSchedule,
	}

	schedule, err := st.NewChargingSchedule(input)
	if err != nil {
		t.Fatalf("NewChargingSchedule: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		if schedule.Duration() != nil {
			_ = schedule.Duration().String()
		}
		if schedule.StartSchedule() != nil {
			_ = schedule.StartSchedule().String()
		}

		_ = schedule.ChargingRateUnit().String()

		for _, period := range schedule.ChargingSchedulePeriod() {
			_ = period.StartPeriod().String()
			_ = period.Limit()
			if period.NumberPhases() != nil {
				_ = period.NumberPhases().String()
			}
		}

		if schedule.MinChargingRate() != nil {
			_ = *schedule.MinChargingRate()
		}

		return nil
	})
}

func TestRace_SetChargingProfileChargingProfileGetters(t *testing.T) {
	t.Parallel()

	duration := 60
	scheduleStart := "2025-01-02T15:00:00Z"
	minChargingRate := 0.0
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	profile, err := scpt.NewChargingProfile(scpt.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          nil,
		StackLevel:             0,
		ChargingProfilePurpose: st.TxProfile.String(),
		ChargingProfileKind:    scpt.ChargingProfileKindAbsolute.String(),
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       scheduleInput,
	})
	if err != nil {
		t.Fatalf("NewChargingProfile: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_ = profile.ChargingProfileId().String()

		if profile.TransactionId() != nil {
			_ = profile.TransactionId().String()
		}

		_ = profile.StackLevel().String()
		_ = profile.ChargingProfilePurpose().String()
		_ = profile.ChargingProfileKind().String()

		if profile.RecurrencyKind() != nil {
			_ = profile.RecurrencyKind().String()
		}
		if profile.ValidFrom() != nil {
			_ = profile.ValidFrom().String()
		}
		if profile.ValidTo() != nil {
			_ = profile.ValidTo().String()
		}

		schedule := profile.ChargingSchedule()
		_ = schedule.ChargingRateUnit().String()
		for _, period := range schedule.ChargingSchedulePeriod() {
			_ = period.StartPeriod().Value()
			_ = period.Limit()
		}

		return nil
	})
}

func TestRace_GetConfigurationKeyValueGetters(t *testing.T) {
	t.Parallel()

	value := "60"
	keyValue, err := gct.NewKeyValue(gct.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		t.Fatalf("NewKeyValue: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_ = keyValue.Key().String()
		_ = keyValue.Readonly()
		if keyValue.Value() != nil {
			_ = keyValue.Value().String()
		}
		return nil
	})
}

func TestRace_ListVersionNumberMethods(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		value := ((worker + iteration) % 3) - 1

		listVersion, err := gllt.NewListVersionNumber(value)
		if err != nil {
			return fmt.Errorf("NewListVersionNumber: %w", err)
		}

		_ = listVersion.Value()
		_ = listVersion.IsUnsupported()
		_ = listVersion.IsEmpty()
		_ = listVersion.String()

		return nil
	})
}
