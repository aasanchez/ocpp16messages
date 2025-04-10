package types

import "testing"

func TestAuthorizationStatusAcceptedIsValid(t *testing.T) {
	t.Parallel()

	status := Accepted
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatusBlockedIsValid(t *testing.T) {
	t.Parallel()

	status := Blocked
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatusExpiredIsValid(t *testing.T) {
	t.Parallel()

	status := Expired
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatusInvalidIsValid(t *testing.T) {
	t.Parallel()

	status := Invalid
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", status)
	}
}

func TestAuthorizationStatusConcurrentTxIsValid(t *testing.T) {
	t.Parallel()

	status := ConcurrentTx
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", status)
	}
}

// Test for invalid AuthorizationStatus values.
func TestAuthorizationStatusInvalidStatus(t *testing.T) {
	t.Parallel()

	invalidStatus := AuthorizationStatus("UnknownStatus")
	if invalidStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", invalidStatus)
	}
}

func TestAuthorizationStatusEmptyStatus(t *testing.T) {
	t.Parallel()

	emptyStatus := AuthorizationStatus("")
	if emptyStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for an empty status, but got true")
	}
}

func TestAuthorizationStatusNullEquivalent(t *testing.T) {
	t.Parallel()

	// Simulating the behavior of null (using an empty string or a special case like "null").
	nullStatus := AuthorizationStatus("null")
	if nullStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v (treated as null), but got true", nullStatus)
	}
}

func TestAuthorizationStatusAcceptedConstant(t *testing.T) {
	t.Parallel()

	if string(Accepted) != "Accepted" {
		t.Errorf("Expected %v, but got %v", "Accepted", Accepted)
	}
}

func TestAuthorizationStatusBlockedConstant(t *testing.T) {
	t.Parallel()

	if string(Blocked) != "Blocked" {
		t.Errorf("Expected %v, but got %v", "Blocked", Blocked)
	}
}

func TestAuthorizationStatusExpiredConstant(t *testing.T) {
	t.Parallel()

	if string(Expired) != "Expired" {
		t.Errorf("Expected %v, but got %v", "Expired", Expired)
	}
}

func TestAuthorizationStatusInvalidConstant(t *testing.T) {
	t.Parallel()

	if string(Invalid) != "Invalid" {
		t.Errorf("Expected %v, but got %v", "Invalid", Invalid)
	}
}

func TestAuthorizationStatusConcurrentTxConstant(t *testing.T) {
	t.Parallel()

	if string(ConcurrentTx) != "ConcurrentTx" {
		t.Errorf("Expected %v, but got %v", "ConcurrentTx", ConcurrentTx)
	}
}
