package types

import "errors"

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// joinErrors joins multiple errors into a single error.
func joinErrors(errs []error) error {
	return errors.Join(errs...)
}
