//go:build race

package race

import (
	"fmt"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func TestRace_GetConfigurationNewKeyValue(t *testing.T) {
	t.Parallel()

	value := "60"
	input := types.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := types.NewKeyValue(input)
		if err != nil {
			return fmt.Errorf("getconfiguration/types.NewKeyValue: %w", err)
		}
		return nil
	})
}

func TestRace_GetLocalListVersionNewListVersionNumber(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		value := (worker + iteration) % 10
		_, err := types.NewListVersionNumber(value)
		if err != nil {
			return fmt.Errorf("getlocallistversion/types.NewListVersionNumber: %w", err)
		}
		return nil
	})
}

func TestRace_SendLocalListNewAuthorizationData(t *testing.T) {
	t.Parallel()

	input := types.AuthorizationDataInput{IdTag: "TAG-1", IdTagInfo: nil}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := types.NewAuthorizationData(input)
		if err != nil {
			return fmt.Errorf("sendlocallist/types.NewAuthorizationData: %w", err)
		}
		return nil
	})
}

func TestRace_SetChargingProfileNewChargingProfile(t *testing.T) {
	t.Parallel()

	duration := 60
	scheduleStart := "2025-01-02T15:00:00Z"
	minChargingRate := 0.0
	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	input := types.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          nil,
		StackLevel:             0,
		ChargingProfilePurpose: types.TxProfile.String(),
		ChargingProfileKind:    types.ChargingProfileKindAbsolute.String(),
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       scheduleInput,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := types.NewChargingProfile(input)
		if err != nil {
			return fmt.Errorf("setchargingprofile/types.NewChargingProfile: %w", err)
		}
		return nil
	})
}

func TestRace_MeterValuesTypesNewSampledValue(t *testing.T) {
	t.Parallel()

	input := types.SampledValueInput{Value: "100"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := types.NewSampledValue(input)
		if err != nil {
			return fmt.Errorf("metervalues/types.NewSampledValue: %w", err)
		}
		return nil
	})
}

func TestRace_MeterValuesTypesNewMeterValue(t *testing.T) {
	t.Parallel()

	sampledValues := []types.SampledValueInput{{Value: "100"}}
	input := types.MeterValueInput{
		Timestamp:    "2025-01-02T15:00:00Z",
		SampledValue: sampledValues,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := types.NewMeterValue(input)
		if err != nil {
			return fmt.Errorf("metervalues/types.NewMeterValue: %w", err)
		}
		return nil
	})
}
