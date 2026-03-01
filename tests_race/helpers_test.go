//go:build race

package race

import (
	"sync"
	"sync/atomic"
	"testing"
)

const (
	raceWorkers    = 16
	raceIterations = 200
)

func runConcurrent(
	t *testing.T,
	workers int,
	iterations int,
	fn func(worker int, iteration int) error,
) {
	t.Helper()

	var wg sync.WaitGroup
	wg.Add(workers)

	errCh := make(chan error, 1)

	var shouldStop atomic.Bool

	for i := 0; i < workers; i++ {
		go func(worker int) {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				if shouldStop.Load() {
					return
				}

				if err := fn(worker, j); err != nil {
					shouldStop.Store(true)
					select {
					case errCh <- err:
					default:
					}

					return
				}
			}
		}(i)
	}

	wg.Wait()

	select {
	case err := <-errCh:
		t.Fatal(err)
	default:
	}
}
