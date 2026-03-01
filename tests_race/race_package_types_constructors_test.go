//go:build race

package race

import (
	"fmt"
	"testing"

	gct "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	gllt "github.com/aasanchez/ocpp16messages/getLocalListVersion/types"
	mvt "github.com/aasanchez/ocpp16messages/meterValues/types"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	scpt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestRace_GetConfigurationNewKeyValue(t *testing.T) {
	t.Parallel()

	value := "60"
	input := gct.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := gct.NewKeyValue(input)
		if err != nil {
			return fmt.Errorf("getConfiguration/types.NewKeyValue: %w", err)
		}
		return nil
	})
}

func TestRace_GetLocalListVersionNewListVersionNumber(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		value := (worker + iteration) % 10
		_, err := gllt.NewListVersionNumber(value)
		if err != nil {
			return fmt.Errorf("getLocalListVersion/types.NewListVersionNumber: %w", err)
		}
		return nil
	})
}

func TestRace_SendLocalListNewAuthorizationData(t *testing.T) {
	t.Parallel()

	input := slt.AuthorizationDataInput{IdTag: "TAG-1", IdTagInfo: nil}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := slt.NewAuthorizationData(input)
		if err != nil {
			return fmt.Errorf("sendLocalList/types.NewAuthorizationData: %w", err)
		}
		return nil
	})
}

func TestRace_SetChargingProfileNewChargingProfile(t *testing.T) {
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

	input := scpt.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          nil,
		StackLevel:             0,
		ChargingProfilePurpose: st.TxProfile.String(),
		ChargingProfileKind:    scpt.ChargingProfileKindAbsolute.String(),
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       scheduleInput,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := scpt.NewChargingProfile(input)
		if err != nil {
			return fmt.Errorf("setChargingProfile/types.NewChargingProfile: %w", err)
		}
		return nil
	})
}

func TestRace_MeterValuesTypesNewSampledValue(t *testing.T) {
	t.Parallel()

	input := mvt.SampledValueInput{Value: "100"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := mvt.NewSampledValue(input)
		if err != nil {
			return fmt.Errorf("meterValues/types.NewSampledValue: %w", err)
		}
		return nil
	})
}

func TestRace_MeterValuesTypesNewMeterValue(t *testing.T) {
	t.Parallel()

	sampledValues := []mvt.SampledValueInput{{Value: "100"}}
	input := mvt.MeterValueInput{
		Timestamp:    "2025-01-02T15:00:00Z",
		SampledValue: sampledValues,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := mvt.NewMeterValue(input)
		if err != nil {
			return fmt.Errorf("meterValues/types.NewMeterValue: %w", err)
		}
		return nil
	})
}
