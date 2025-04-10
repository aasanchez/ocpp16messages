package types

import "testing"

func TestAuthorizationStatus_AcceptedIsValid(t *testing.T) {
	status := Accepted
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatus_BlockedIsValid(t *testing.T) {
	status := Blocked
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatus_ExpiredIsValid(t *testing.T) {
	status := Expired
	if !status.IsValid() {
		t.Errorf("Expected IsValid() to return true for %v, but got false", status)
	}
}

func TestAuthorizationStatus_InvalidIsValid(t *testing.T) {
	status := Invalid
	if status.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", status)
	}
}

func TestAuthorizationStatus_ConcurrentTxIsValid(t *testing.T) {
	status := ConcurrentTx
	if status.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", status)
	}
}

// Test for invalid AuthorizationStatus values
func TestAuthorizationStatus_InvalidStatus(t *testing.T) {
	invalidStatus := AuthorizationStatus("UnknownStatus")
	if invalidStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v, but got true", invalidStatus)
	}
}

func TestAuthorizationStatus_EmptyStatus(t *testing.T) {
	emptyStatus := AuthorizationStatus("")
	if emptyStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for an empty status, but got true")
	}
}

func TestAuthorizationStatus_NullEquivalent(t *testing.T) {
	// Simulating the behavior of null (using an empty string or a special case like "null")
	nullStatus := AuthorizationStatus("null")
	if nullStatus.IsValid() {
		t.Errorf("Expected IsValid() to return false for %v (treated as null), but got true", nullStatus)
	}
}

func TestAuthorizationStatus_AcceptedConstant(t *testing.T) {
	if string(Accepted) != "Accepted" {
		t.Errorf("Expected %v, but got %v", "Accepted", Accepted)
	}
}

func TestAuthorizationStatus_BlockedConstant(t *testing.T) {
	if string(Blocked) != "Blocked" {
		t.Errorf("Expected %v, but got %v", "Blocked", Blocked)
	}
}

func TestAuthorizationStatus_ExpiredConstant(t *testing.T) {
	if string(Expired) != "Expired" {
		t.Errorf("Expected %v, but got %v", "Expired", Expired)
	}
}

func TestAuthorizationStatus_InvalidConstant(t *testing.T) {
	if string(Invalid) != "Invalid" {
		t.Errorf("Expected %v, but got %v", "Invalid", Invalid)
	}
}

func TestAuthorizationStatus_ConcurrentTxConstant(t *testing.T) {
	if string(ConcurrentTx) != "ConcurrentTx" {
		t.Errorf("Expected %v, but got %v", "ConcurrentTx", ConcurrentTx)
	}
}
