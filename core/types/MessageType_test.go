package types

import (
	"testing"
)

func TestMessageTypeConstants(t *testing.T) {
	if CALL != 2 {
		t.Errorf("Expected CALL to be 2, got %d", CALL)
	}
	if CALLRESULT != 3 {
		t.Errorf("Expected CALLRESULT to be 3, got %d", CALLRESULT)
	}
	if CALLERROR != 4 {
		t.Errorf("Expected CALLERROR to be 4, got %d", CALLERROR)
	}
}

func TestMessageType_IsValid_CALL(t *testing.T) {
	if !CALL.IsValid() {
		t.Error("CALL should be valid")
	}
}

func TestMessageType_IsValid_CALLRESULT(t *testing.T) {
	if !CALLRESULT.IsValid() {
		t.Error("CALLRESULT should be valid")
	}
}

func TestMessageType_IsValid_CALLERROR(t *testing.T) {
	if !CALLERROR.IsValid() {
		t.Error("CALLERROR should be valid")
	}
}

func TestMessageType_IsValid_InvalidLow(t *testing.T) {
	if MessageType(0).IsValid() {
		t.Error("0 should be invalid")
	}
}

func TestMessageType_IsValid_InvalidHigh(t *testing.T) {
	if MessageType(5).IsValid() {
		t.Error("5 should be invalid")
	}
}
