package core_test

import (
	"encoding/json"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestParseJSONMessage_ValidCALL(t *testing.T) {
	raw := []byte(`[2, "12345", "Authorize", {"idTag": "ABC123"}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALL {
		t.Errorf("expected TypeID %v, got %v", core.CALL, msg.TypeID)
	}
	if msg.UniqueID != "12345" {
		t.Errorf("expected UniqueID '12345', got '%s'", msg.UniqueID)
	}
	if msg.Action != "Authorize" {
		t.Errorf("expected Action 'Authorize', got '%s'", msg.Action)
	}
	if !jsonEqual(msg.Payload, `{"idTag": "ABC123"}`) {
		t.Errorf("unexpected Payload: %s", string(msg.Payload))
	}
}

func TestParseJSONMessage_ValidCALLRESULT(t *testing.T) {
	raw := []byte(`[3, "12345", {"status": "Accepted"}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALLRESULT {
		t.Errorf("expected TypeID %v, got %v", core.CALLRESULT, msg.TypeID)
	}
	if msg.UniqueID != "12345" {
		t.Errorf("expected UniqueID '12345', got '%s'", msg.UniqueID)
	}
	if !jsonEqual(msg.Payload, `{"status": "Accepted"}`) {
		t.Errorf("unexpected Payload: %s", string(msg.Payload))
	}
}

func TestParseJSONMessage_ValidCALLERROR(t *testing.T) {
	raw := []byte(`[4, "12345", "InternalError", "Something went wrong", {}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALLERROR {
		t.Errorf("expected TypeID %v, got %v", core.CALLERROR, msg.TypeID)
	}
	if msg.UniqueID != "12345" {
		t.Errorf("expected UniqueID '12345', got '%s'", msg.UniqueID)
	}
	if msg.ErrorCode != "InternalError" {
		t.Errorf("expected ErrorCode 'InternalError', got '%s'", msg.ErrorCode)
	}
	if msg.ErrorDescription != "Something went wrong" {
		t.Errorf("expected ErrorDescription 'Something went wrong', got '%s'", msg.ErrorDescription)
	}
	if !jsonEqual(msg.ErrorDetails, `{}`) {
		t.Errorf("unexpected ErrorDetails: %s", string(msg.ErrorDetails))
	}
}

func TestParseJSONMessage_InvalidMessageType(t *testing.T) {
	raw := []byte(`[99, "12345", "Authorize", {}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err == nil || msg != nil {
		t.Errorf("expected error for invalid message type, got: %v", err)
	}
}

func TestParseJSONMessage_InvalidJSON(t *testing.T) {
	raw := []byte(`{ this is not valid JSON }`)
	msg, err := core.ParseJSONMessage(raw)
	if err == nil || msg != nil {
		t.Errorf("expected error for invalid JSON, got: %v", err)
	}
}

func TestParseJSONMessage_EmptyPayload(t *testing.T) {
	raw := []byte(`[]`)
	msg, err := core.ParseJSONMessage(raw)
	if err == nil || msg != nil {
		t.Errorf("expected error for empty message, got: %v", err)
	}
}

// Helper: compares two JSON strings for semantic equality.
func jsonEqual(a json.RawMessage, b string) bool {
	var objA interface{}
	var objB interface{}
	if err := json.Unmarshal(a, &objA); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &objB); err != nil {
		return false
	}
	return stringify(objA) == stringify(objB)
}

func stringify(v interface{}) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}
