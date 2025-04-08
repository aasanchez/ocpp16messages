package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

func TestMessageTypeConstants(t *testing.T) {
	if types.CALL != 2 {
		t.Errorf("Expected CALL to be 2, got %d", types.CALL)
	}
	if types.CALLRESULT != 3 {
		t.Errorf("Expected CALLRESULT to be 3, got %d", types.CALLRESULT)
	}
	if types.CALLERROR != 4 {
		t.Errorf("Expected CALLERROR to be 4, got %d", types.CALLERROR)
	}
}

func TestMessageTypeIsValidCALL(t *testing.T) {
	if !types.CALL.IsValid() {
		t.Error("CALL should be valid")
	}
}

func TestMessageTypeIsValidCALLRESULT(t *testing.T) {
	if !types.CALLRESULT.IsValid() {
		t.Error("CALLRESULT should be valid")
	}
}

func TestMessageTypeIsValidCALLERROR(t *testing.T) {
	if !types.CALLERROR.IsValid() {
		t.Error("CALLERROR should be valid")
	}
}

func TestMessageTypeIsValidInvalidLow(t *testing.T) {
	if types.MessageType(0).IsValid() {
		t.Error("0 should be invalid")
	}
}

func TestMessageTypeIsValidInvalidHigh(t *testing.T) {
	if types.MessageType(5).IsValid() {
		t.Error("5 should be invalid")
	}
}
