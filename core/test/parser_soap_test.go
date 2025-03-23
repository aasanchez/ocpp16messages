package core_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestParseSOAPMessage_Valid(t *testing.T) {
	xml := `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
		<soap:Header>
			<Action>Authorize</Action>
		</soap:Header>
		<soap:Body>
			<idTag>XYZ789</idTag>
		</soap:Body>
	</soap:Envelope>
	`

	msg, err := core.ParseSOAPMessage(bytes.NewReader([]byte(xml)))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALL {
		t.Errorf("expected TypeID CALL, got %v", msg.TypeID)
	}
	if msg.Action != "Authorize" {
		t.Errorf("expected Action 'Authorize', got '%s'", msg.Action)
	}
	if strings.TrimSpace(string(msg.Payload)) != "<idTag>XYZ789</idTag>" {
		t.Errorf("unexpected Payload: %s", string(msg.Payload))
	}
}

func TestParseSOAPMessage_MissingAction(t *testing.T) {
	xml := `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
		<soap:Header></soap:Header>
		<soap:Body>
			<idTag>XYZ789</idTag>
		</soap:Body>
	</soap:Envelope>
	`

	_, err := core.ParseSOAPMessage(bytes.NewReader([]byte(xml)))
	if err == nil {
		t.Error("expected error for missing Action")
	}
}

func TestParseSOAPMessage_MissingBody(t *testing.T) {
	xml := `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
		<soap:Header>
			<Action>Authorize</Action>
		</soap:Header>
	</soap:Envelope>
	`

	_, err := core.ParseSOAPMessage(bytes.NewReader([]byte(xml)))
	if err == nil {
		t.Error("expected error for missing Body")
	}
}

func TestParseSOAPMessage_InvalidXML(t *testing.T) {
	xml := `<Envelope><Header></Header><Body></Body>` // Invalid XML structure
	_, err := core.ParseSOAPMessage(bytes.NewReader([]byte(xml)))
	if err == nil {
		t.Error("expected error for invalid XML")
	}
}
