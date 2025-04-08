package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/error/types"
)

// Test for the String method
func TestErrorCode_String_ErrorNotSupported(t *testing.T) {
	code := types.ErrorNotSupported
	expected := "NotSupported"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorInternalError(t *testing.T) {
	code := types.ErrorInternalError
	expected := "InternalError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorProtocolError(t *testing.T) {
	code := types.ErrorProtocolError
	expected := "ProtocolError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorSecurityError(t *testing.T) {
	code := types.ErrorSecurityError
	expected := "SecurityError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorFormationViolation(t *testing.T) {
	code := types.ErrorFormationViolation
	expected := "FormationViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorPropertyConstraintViolation(t *testing.T) {
	code := types.ErrorPropertyConstraintViolation
	expected := "PropertyConstraintViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorOccurrenceConstraintViolation(t *testing.T) {
	code := types.ErrorOccurrenceConstraintViolation
	expected := "OccurrenceConstraintViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorTypeConstraintViolation(t *testing.T) {
	code := types.ErrorTypeConstraintViolation
	expected := "TypeConstraintViolation"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_String_ErrorGenericError(t *testing.T) {
	code := types.ErrorGenericError
	expected := "GenericError"
	result := code.String()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorNotSupported(t *testing.T) {
	code := types.ErrorNotSupported
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorInternalError(t *testing.T) {
	code := types.ErrorInternalError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorProtocolError(t *testing.T) {
	code := types.ErrorProtocolError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorSecurityError(t *testing.T) {
	code := types.ErrorSecurityError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorFormationViolation(t *testing.T) {
	code := types.ErrorFormationViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorPropertyConstraintViolation(t *testing.T) {
	code := types.ErrorPropertyConstraintViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorOccurrenceConstraintViolation(t *testing.T) {
	code := types.ErrorOccurrenceConstraintViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorTypeConstraintViolation(t *testing.T) {
	code := types.ErrorTypeConstraintViolation
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_ErrorGenericError(t *testing.T) {
	code := types.ErrorGenericError
	expected := true
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}

func TestErrorCode_IsValid_InvalidErrorCode(t *testing.T) {
	code := types.ErrorCode("InvalidErrorCode")
	expected := false
	result := code.IsValid()
	if result != expected {
		t.Errorf("expected validity %v, got %v", expected, result)
	}
}
