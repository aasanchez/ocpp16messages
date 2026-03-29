//go:build race

package race

import (
	"fmt"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func TestRace_ChargingScheduleDurationGetterReturnsCopy(t *testing.T) {
	t.Parallel()

	duration := 60
	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          nil,
	})
	if err != nil {
		t.Fatalf("NewChargingSchedule: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		value := schedule.Duration()
		if value == nil {
			return fmt.Errorf("Duration: expected non-nil")
		}

		if worker == 0 {
			mutated, err := types.NewInteger(iteration)
			if err != nil {
				return fmt.Errorf("NewInteger: %w", err)
			}

			*value = mutated
			return nil
		}

		if value.Value() != 60 {
			return fmt.Errorf("Duration.Value() = %d, want %d", value.Value(), 60)
		}

		return nil
	})
}

func TestRace_ChargingScheduleStartScheduleGetterReturnsCopy(t *testing.T) {
	t.Parallel()

	duration := 60
	startSchedule := "2025-01-02T15:00:00Z"
	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          &startSchedule,
	})
	if err != nil {
		t.Fatalf("NewChargingSchedule: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		value := schedule.StartSchedule()
		if value == nil {
			return fmt.Errorf("StartSchedule: expected non-nil")
		}

		if worker == 0 {
			mutated, err := types.NewDateTime(
				fmt.Sprintf("2025-01-%02dT00:00:00Z", (iteration%28)+1),
			)
			if err != nil {
				return fmt.Errorf("NewDateTime: %w", err)
			}

			*value = mutated
			return nil
		}

		if value.String() != startSchedule {
			return fmt.Errorf("StartSchedule.String() = %q, want %q", value.String(), startSchedule)
		}

		return nil
	})
}

func TestRace_ChargingSchedulePeriodNumberPhasesGetterReturnsCopy(t *testing.T) {
	t.Parallel()

	phases := 3
	period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  0,
		Limit:        16,
		NumberPhases: &phases,
	})
	if err != nil {
		t.Fatalf("NewChargingSchedulePeriod: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		value := period.NumberPhases()
		if value == nil {
			return fmt.Errorf("NumberPhases: expected non-nil")
		}

		if worker == 0 {
			mutated, err := types.NewInteger((iteration % 3) + 1)
			if err != nil {
				return fmt.Errorf("NewInteger: %w", err)
			}

			*value = mutated
			return nil
		}

		if value.Value() != 3 {
			return fmt.Errorf("NumberPhases.Value() = %d, want %d", value.Value(), 3)
		}

		return nil
	})
}

func TestRace_GetConfigurationKeyValueValueGetterReturnsCopy(t *testing.T) {
	t.Parallel()

	value := "60"
	keyValue, err := types.NewKeyValue(types.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		t.Fatalf("NewKeyValue: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		current := keyValue.Value()
		if current == nil {
			return fmt.Errorf("KeyValue.Value: expected non-nil")
		}

		if worker == 0 {
			mutated, err := types.NewCiString500Type(fmt.Sprintf("mutated-%d", iteration))
			if err != nil {
				return fmt.Errorf("NewCiString500Type: %w", err)
			}

			*current = mutated
			return nil
		}

		if current.String() != "60" {
			return fmt.Errorf("KeyValue.Value().String() = %q, want %q", current.String(), "60")
		}

		return nil
	})
}

func TestRace_SetChargingProfileTransactionIdGetterReturnsCopy(t *testing.T) {
	t.Parallel()

	transactionId := 1
	duration := 60
	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          nil,
	}

	profile, err := types.NewChargingProfile(types.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          &transactionId,
		StackLevel:             0,
		ChargingProfilePurpose: types.TxProfile.String(),
		ChargingProfileKind:    types.ChargingProfileKindAbsolute.String(),
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       scheduleInput,
	})
	if err != nil {
		t.Fatalf("NewChargingProfile: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		current := profile.TransactionId()
		if current == nil {
			return fmt.Errorf("TransactionId: expected non-nil")
		}

		if worker == 0 {
			mutated, err := types.NewInteger(iteration)
			if err != nil {
				return fmt.Errorf("NewInteger: %w", err)
			}

			*current = mutated
			return nil
		}

		if current.Value() != 1 {
			return fmt.Errorf("TransactionId.Value() = %d, want %d", current.Value(), 1)
		}

		return nil
	})
}
