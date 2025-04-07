package errorTypes

import "testing"

// Valid tests

func TestErrorCode_IsValid_NotSupported(t *testing.T) {
	if !ErrorNotSupported.IsValid() {
		t.Error("ErrorNotSupported should be valid")
	}
}

func TestErrorCode_IsValid_InternalError(t *testing.T) {
	if !ErrorInternalError.IsValid() {
		t.Error("ErrorInternalError should be valid")
	}
}

func TestErrorCode_IsValid_ProtocolError(t *testing.T) {
	if !ErrorProtocolError.IsValid() {
		t.Error("ErrorProtocolError should be valid")
	}
}

func TestErrorCode_IsValid_SecurityError(t *testing.T) {
	if !ErrorSecurityError.IsValid() {
		t.Error("ErrorSecurityError should be valid")
	}
}

func TestErrorCode_IsValid_FormationViolation(t *testing.T) {
	if !ErrorFormationViolation.IsValid() {
		t.Error("ErrorFormationViolation should be valid")
	}
}

func TestErrorCode_IsValid_PropertyConstraintViolation(t *testing.T) {
	if !ErrorPropertyConstraintViolation.IsValid() {
		t.Error("ErrorPropertyConstraintViolation should be valid")
	}
}

func TestErrorCode_IsValid_OccurrenceConstraintViolation(t *testing.T) {
	if !ErrorOccurrenceConstraintViolation.IsValid() {
		t.Error("ErrorOccurrenceConstraintViolation should be valid")
	}
}

func TestErrorCode_IsValid_TypeConstraintViolation(t *testing.T) {
	if !ErrorTypeConstraintViolation.IsValid() {
		t.Error("ErrorTypeConstraintViolation should be valid")
	}
}

func TestErrorCode_IsValid_GenericError(t *testing.T) {
	if !ErrorGenericError.IsValid() {
		t.Error("ErrorGenericError should be valid")
	}
}

// Invalid tests
func TestErrorCode_IsValid_Empty(t *testing.T) {
	if ErrorCode("").IsValid() {
		t.Error("Empty string should not be valid")
	}
}

func TestErrorCode_IsValid_UnknownValue(t *testing.T) {
	if ErrorCode("SomeUnknownError").IsValid() {
		t.Error("Unknown error code should not be valid")
	}
}

func TestErrorCode_IsValid_Misspelled(t *testing.T) {
	if ErrorCode("InternelError").IsValid() {
		t.Error("Misspelled error code should not be valid")
	}
}
