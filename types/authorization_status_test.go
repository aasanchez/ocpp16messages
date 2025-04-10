package types

import "testing"

// Error message constants.
const (
	errExpectedOutput       = "expected string output, got %s"
	errExpectedValidate     = "expected Validate to return nil, got %v"
	errExpectedInvalid      = "expected IsValid() to return false for %v, but got true" // For invalid status checks
	errExpectedValid        = "expected IsValid() to return true for %v, but got false" // For valid status checks
	errExpectedInvalidNoArg = "expected IsValid() to return false, but got true"        // For cases with no specific arg
)

// Tests for valid AuthorizationStatus values.
func TestAuthorizationStatusAcceptedIsValid(t *testing.T) {
	t.Parallel()

	status := Accepted
	if !status.IsValid() {
		t.Errorf(errExpectedValid, status)
	}
}

func TestAuthorizationStatusBlockedIsValid(t *testing.T) {
	t.Parallel()

	status := Blocked
	if !status.IsValid() {
		t.Errorf(errExpectedValid, status)
	}
}

func TestAuthorizationStatusExpiredIsValid(t *testing.T) {
	t.Parallel()

	status := Expired
	if !status.IsValid() {
		t.Errorf(errExpectedValid, status)
	}
}

// Tests for invalid AuthorizationStatus values.
func TestAuthorizationStatusInvalidIsValid(t *testing.T) {
	t.Parallel()

	status := Invalid
	if !status.IsValid() {
		t.Errorf(errExpectedInvalid, status)
	}
}

func TestAuthorizationStatusConcurrentTxIsValid(t *testing.T) {
	t.Parallel()

	status := ConcurrentTx
	if !status.IsValid() {
		t.Errorf(errExpectedInvalid, status)
	}
}

// Test for invalid AuthorizationStatus values.
func TestAuthorizationStatusInvalidStatus(t *testing.T) {
	t.Parallel()

	invalidStatus := AuthorizationStatus("UnknownStatus")
	if invalidStatus.IsValid() {
		t.Errorf(errExpectedInvalid, invalidStatus)
	}
}

func TestAuthorizationStatusEmptyStatus(t *testing.T) {
	t.Parallel()

	emptyStatus := AuthorizationStatus("")
	if emptyStatus.IsValid() {
		t.Error(errExpectedInvalidNoArg) // Use t.Error() here instead of t.Errorf
	}
}

func TestAuthorizationStatusNullEquivalent(t *testing.T) {
	t.Parallel()

	// Simulating the behavior of null (using an empty string or a special case like "null").
	nullStatus := AuthorizationStatus("null")
	if nullStatus.IsValid() {
		t.Errorf(errExpectedInvalid, nullStatus)
	}
}

// Constant value checks.
func TestAuthorizationStatusAcceptedConstant(t *testing.T) {
	t.Parallel()

	if string(Accepted) != "Accepted" {
		t.Errorf(errExpectedOutput, Accepted)
	}
}

func TestAuthorizationStatusBlockedConstant(t *testing.T) {
	t.Parallel()

	if string(Blocked) != "Blocked" {
		t.Errorf(errExpectedOutput, Blocked)
	}
}

func TestAuthorizationStatusExpiredConstant(t *testing.T) {
	t.Parallel()

	if string(Expired) != "Expired" {
		t.Errorf(errExpectedOutput, Expired)
	}
}

func TestAuthorizationStatusInvalidConstant(t *testing.T) {
	t.Parallel()

	if string(Invalid) != "Invalid" {
		t.Errorf(errExpectedOutput, Invalid)
	}
}

func TestAuthorizationStatusConcurrentTxConstant(t *testing.T) {
	t.Parallel()

	if string(ConcurrentTx) != "ConcurrentTx" {
		t.Errorf(errExpectedOutput, ConcurrentTx)
	}
}
