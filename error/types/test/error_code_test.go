package errorTypes_test

import (
	"testing"

	errorTypes "github.com/aasanchez/ocpp16_messages/error/types"
)

// Test for the String method
func TestErrorCode_String_ErrorNotSupported(t *testing.T) {
	code := errorTypes.ErrorNotSupported
	expected := "NotSupported"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorInternalError(t *testing.T) {
	code := errorTypes.ErrorInternalError
	expected := "InternalError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorProtocolError(t *testing.T) {
	code := errorTypes.ErrorProtocolError
	expected := "ProtocolError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorSecurityError(t *testing.T) {
	code := errorTypes.ErrorSecurityError
	expected := "SecurityError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorFormationViolation(t *testing.T) {
	code := errorTypes.ErrorFormationViolation
	expected := "FormationViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorPropertyConstraintViolation(t *testing.T) {
	code := errorTypes.ErrorPropertyConstraintViolation
	expected := "PropertyConstraintViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorOccurrenceConstraintViolation(t *testing.T) {
	code := errorTypes.ErrorOccurrenceConstraintViolation
	expected := "OccurrenceConstraintViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorTypeConstraintViolation(t *testing.T) {
	code := errorTypes.ErrorTypeConstraintViolation
	expected := "TypeConstraintViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorGenericError(t *testing.T) {
	code := errorTypes.ErrorGenericError
	expected := "GenericError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

// Test for the IsValid method
func TestErrorCode_IsValid_ErrorNotSupported(t *testing.T) {
	code := errorTypes.ErrorNotSupported
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorInternalError(t *testing.T) {
	code := errorTypes.ErrorInternalError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorProtocolError(t *testing.T) {
	code := errorTypes.ErrorProtocolError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorSecurityError(t *testing.T) {
	code := errorTypes.ErrorSecurityError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorFormationViolation(t *testing.T) {
	code := errorTypes.ErrorFormationViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorPropertyConstraintViolation(t *testing.T) {
	code := errorTypes.ErrorPropertyConstraintViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorOccurrenceConstraintViolation(t *testing.T) {
	code := errorTypes.ErrorOccurrenceConstraintViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorTypeConstraintViolation(t *testing.T) {
	code := errorTypes.ErrorTypeConstraintViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorGenericError(t *testing.T) {
	code := errorTypes.ErrorGenericError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_InvalidErrorCode(t *testing.T) {
	code := errorTypes.ErrorCode("InvalidErrorCode")
	expected := false
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}
