//go:build race

package race

import (
	"fmt"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

const (
	workers          = 32
	iterations       = 500
	ciStringTemplate = "TAG-%d-%d"
)

func TestRace_NewCiString20Type(t *testing.T) {
	t.Parallel()

	runConcurrent(t, workers, iterations, func(worker int, iteration int) error {
		value, err := types.NewCiString20Type(
			fmt.Sprintf(ciStringTemplate, worker, iteration),
		)
		if err != nil {
			return fmt.Errorf("NewCiString20Type: %w", err)
		}

		_ = value.String()

		return nil
	})
}

func TestRace_NewInteger(t *testing.T) {
	t.Parallel()

	const maxValue = 65535

	runConcurrent(t, workers, iterations, func(worker int, iteration int) error {
		n := (worker + iteration) % maxValue
		value, err := types.NewInteger(n)
		if err != nil {
			return fmt.Errorf("NewInteger: %w", err)
		}

		_ = value.String()

		return nil
	})
}
