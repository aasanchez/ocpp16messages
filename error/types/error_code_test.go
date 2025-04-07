package errorTypes

import "testing"

// Valid tests

func TestErrorCodeIsValidNotSupported(t *testing.T) {
	if !ErrorNotSupported.IsValid() {
		t.Error("ErrorNotSupported should be valid")
	}
}

func TestErrorCodeIsValidInternalError(t *testing.T) {
	if !ErrorInternalError.IsValid() {
		t.Error("ErrorInternalError should be valid")
	}
}

func TestErrorCodeIsValidProtocolError(t *testing.T) {
	if !ErrorProtocolError.IsValid() {
		t.Error("ErrorProtocolError should be valid")
	}
}

func TestErrorCodeIsValidSecurityError(t *testing.T) {
	if !ErrorSecurityError.IsValid() {
		t.Error("ErrorSecurityError should be valid")
	}
}

func TestErrorCodeIsValidFormationViolation(t *testing.T) {
	if !ErrorFormationViolation.IsValid() {
		t.Error("ErrorFormationViolation should be valid")
	}
}

func TestErrorCodeIsValidPropertyConstraintViolation(t *testing.T) {
	if !ErrorPropertyConstraintViolation.IsValid() {
		t.Error("ErrorPropertyConstraintViolation should be valid")
	}
}

func TestErrorCodeIsValidOccurrenceConstraintViolation(t *testing.T) {
	if !ErrorOccurrenceConstraintViolation.IsValid() {
		t.Error("ErrorOccurrenceConstraintViolation should be valid")
	}
}

func TestErrorCodeIsValidTypeConstraintViolation(t *testing.T) {
	if !ErrorTypeConstraintViolation.IsValid() {
		t.Error("ErrorTypeConstraintViolation should be valid")
	}
}

func TestErrorCodeIsValidGenericError(t *testing.T) {
	if !ErrorGenericError.IsValid() {
		t.Error("ErrorGenericError should be valid")
	}
}

// Invalid tests
func TestErrorCodeIsValidEmpty(t *testing.T) {
	if ErrorCode("").IsValid() {
		t.Error("Empty string should not be valid")
	}
}

func TestErrorCodeIsValidUnknownValue(t *testing.T) {
	if ErrorCode("SomeUnknownError").IsValid() {
		t.Error("Unknown error code should not be valid")
	}
}

func TestErrorCodeIsValidMisspelled(t *testing.T) {
	if ErrorCode("InternelError").IsValid() {
		t.Error("Misspelled error code should not be valid")
	}
}
