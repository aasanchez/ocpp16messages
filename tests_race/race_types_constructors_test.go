//go:build race

package race

import (
	"fmt"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

const (
	typeWorkers    = 32
	typeIterations = 300
)

func TestRace_NewDateTime(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		input := fmt.Sprintf(
			"2025-01-%02dT%02d:%02d:%02dZ",
			(iteration%28)+1,
			worker%24,
			iteration%60,
			(worker+iteration)%60,
		)

		_, err := types.NewDateTime(input)
		if err != nil {
			return fmt.Errorf("NewDateTime: %w", err)
		}

		return nil
	})
}

func TestRace_NewCiString25Type(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		value, err := types.NewCiString25Type(
			fmt.Sprintf("SERIAL-%d-%d", worker, iteration),
		)
		if err != nil {
			return fmt.Errorf("NewCiString25Type: %w", err)
		}

		_ = value.String()

		return nil
	})
}

func TestRace_NewCiString50Type(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		value, err := types.NewCiString50Type(
			fmt.Sprintf("KEY-%d-%d", worker, iteration),
		)
		if err != nil {
			return fmt.Errorf("NewCiString50Type: %w", err)
		}

		_ = value.String()

		return nil
	})
}

func TestRace_NewCiString255Type(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		value, err := types.NewCiString255Type(
			fmt.Sprintf("https://example.com/%d/%d", worker, iteration),
		)
		if err != nil {
			return fmt.Errorf("NewCiString255Type: %w", err)
		}

		_ = value.String()

		return nil
	})
}

func TestRace_NewCiString500Type(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		value, err := types.NewCiString500Type(
			fmt.Sprintf("VALUE-%d-%d", worker, iteration),
		)
		if err != nil {
			return fmt.Errorf("NewCiString500Type: %w", err)
		}

		_ = value.String()

		return nil
	})
}

func TestRace_NewChargingSchedulePeriod(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
			StartPeriod:  iteration % 3600,
			Limit:        float64((worker%100)+1) / 10,
			NumberPhases: nil,
		})
		if err != nil {
			return fmt.Errorf("NewChargingSchedulePeriod: %w", err)
		}

		return nil
	})
}

func TestRace_NewChargingSchedule(t *testing.T) {
	t.Parallel()

	duration := 60
	startSchedule := "2025-01-02T15:00:00Z"
	minChargingRate := 0.0

	periods := []types.ChargingSchedulePeriodInput{
		{
			StartPeriod:  0,
			Limit:        16,
			NumberPhases: nil,
		},
	}

	input := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &startSchedule,
	}

	runConcurrent(t, typeWorkers, typeIterations, func(_, _ int) error {
		_, err := types.NewChargingSchedule(input)
		if err != nil {
			return fmt.Errorf("NewChargingSchedule: %w", err)
		}

		return nil
	})
}

func TestRace_NewIdTagInfo(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, _ int) error {
		status := types.AuthorizationStatusAccepted
		if worker%2 == 1 {
			status = types.AuthorizationStatusBlocked
		}

		_, err := types.NewIdTagInfo(status)
		if err != nil {
			return fmt.Errorf("NewIdTagInfo: %w", err)
		}

		return nil
	})
}

func TestRace_NewIdToken(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		tag, err := types.NewCiString20Type(
			fmt.Sprintf("ID-%d-%d", worker, iteration),
		)
		if err != nil {
			return fmt.Errorf("NewCiString20Type: %w", err)
		}

		_ = types.NewIdToken(tag).String()

		return nil
	})
}

func TestRace_NewSharedTypeSampledValue(t *testing.T) {
	t.Parallel()

	runConcurrent(t, typeWorkers, typeIterations, func(worker int, iteration int) error {
		context := types.ReadingContextSamplePeriodic.String()
		format := types.ValueFormatRaw.String()
		measurand := types.MeasurandEnergyActiveImportRegister.String()
		phase := types.PhaseL1.String()
		location := types.LocationOutlet.String()
		unit := types.UnitWh.String()

		sv, err := types.NewSampledValue(types.SampledValueInput{
			Value:     fmt.Sprintf("%d", worker+iteration),
			Context:   &context,
			Format:    &format,
			Measurand: &measurand,
			Phase:     &phase,
			Location:  &location,
			Unit:      &unit,
		})
		if err != nil {
			return fmt.Errorf("NewSampledValue: %w", err)
		}

		_ = sv.Value().String()

		return nil
	})
}

func TestRace_NewSharedTypeMeterValue(t *testing.T) {
	t.Parallel()

	sampled := []types.SampledValueInput{
		{
			Value: "100",
		},
	}

	input := types.MeterValueInput{
		Timestamp:    "2025-01-02T15:00:00Z",
		SampledValue: sampled,
	}

	runConcurrent(t, typeWorkers, typeIterations, func(_, _ int) error {
		mv, err := types.NewMeterValue(input)
		if err != nil {
			return fmt.Errorf("NewMeterValue: %w", err)
		}

		_ = mv.Timestamp().String()
		for _, value := range mv.SampledValue() {
			_ = value.Value().String()
		}

		return nil
	})
}
