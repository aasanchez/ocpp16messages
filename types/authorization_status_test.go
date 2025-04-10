package types

import "testing"

func TestAuthorizationStatusAcceptedIsValid(t *testing.T) {
	status := Accepted
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatusBlockedIsValid(t *testing.T) {
	status := Blocked
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatusExpiredIsValid(t *testing.T) {
	status := Expired
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatusInvalidIsValid(t *testing.T) {
	status := Invalid
	if status.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", status)
	}
}

func TestAuthorizationStatusConcurrentTxIsValid(t *testing.T) {
	status := ConcurrentTx
	if status.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", status)
	}
}

// Test for invalid AuthorizationStatus values
func TestAuthorizationStatusInvalidStatus(t *testing.T) {
	invalidStatus := AuthorizationStatus("UnknownStatus")
	if invalidStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", invalidStatus)
	}
}

func TestAuthorizationStatusEmptyStatus(t *testing.T) {
	emptyStatus := AuthorizationStatus("")
	if emptyStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for an empty status, but got true")
	}
}

func TestAuthorizationStatusNullEquivalent(t *testing.T) {
	// Simulating the behavior of null (using an empty string or a special case like "null")
	nullStatus := AuthorizationStatus("null")
	if nullStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v (treated as null), but got true", nullStatus)
	}
}

func TestAuthorizationStatusAcceptedConstant(t *testing.T) {
	if string(Accepted) != "Accepted" {
		t.Errorf("Expected %v, but got %v", "Accepted", Accepted)
	}
}

func TestAuthorizationStatusBlockedConstant(t *testing.T) {
	if string(Blocked) != "Blocked" {
		t.Errorf("Expected %v, but got %v", "Blocked", Blocked)
	}
}

func TestAuthorizationStatusExpiredConstant(t *testing.T) {
	if string(Expired) != "Expired" {
		t.Errorf("Expected %v, but got %v", "Expired", Expired)
	}
}

func TestAuthorizationStatusInvalidConstant(t *testing.T) {
	if string(Invalid) != "Invalid" {
		t.Errorf("Expected %v, but got %v", "Invalid", Invalid)
	}
}

func TestAuthorizationStatusConcurrentTxConstant(t *testing.T) {
	if string(ConcurrentTx) != "ConcurrentTx" {
		t.Errorf("Expected %v, but got %v", "ConcurrentTx", ConcurrentTx)
	}
}
