package core_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestParseSOAPMessage_Valid(t *testing.T) {
	soap := `
		<Envelope>
			<Header>
				<Action>Authorize</Action>
			</Header>
			<Body>
				<AuthorizeRequest>
					<idTag>ABC123</idTag>
				</AuthorizeRequest>
			</Body>
		</Envelope>`

	msg, err := core.ParseSOAPMessage(strings.NewReader(soap))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALL {
		t.Errorf("expected TypeID CALL, got %v", msg.TypeID)
	}
	if msg.Action != "Authorize" {
		t.Errorf("expected Action 'Authorize', got '%s'", msg.Action)
	}
	if !strings.Contains(string(msg.Payload), "<AuthorizeRequest>") {
		t.Errorf("expected payload to contain AuthorizeRequest, got: %s", msg.Payload)
	}
}

func TestParseSOAPMessage_ValidWithAttributes(t *testing.T) {
	soap := `
		<Envelope>
			<Header>
				<Action>Authorize</Action>
			</Header>
			<Body>
				<AuthorizeRequest id="123" type="request">
					<idTag>ABC123</idTag>
				</AuthorizeRequest>
			</Body>
		</Envelope>`

	msg, err := core.ParseSOAPMessage(strings.NewReader(soap))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(msg.Payload), `id="123"`) {
		t.Errorf("expected payload to contain id attribute, got: %s", msg.Payload)
	}
	if !strings.Contains(string(msg.Payload), `type="request"`) {
		t.Errorf("expected payload to contain type attribute, got: %s", msg.Payload)
	}
}

func TestParseSOAPMessage_InvalidActionElement(t *testing.T) {
	soap := `
		<Envelope>
			<Header>
				<Action><![CDATA[Invalid Action content]]></Action>
			</Header>
			<Body>
				<AuthorizeRequest>
					<idTag>ABC123</idTag>
				</AuthorizeRequest>
			</Body>
		</Envelope>`

	_, err := core.ParseSOAPMessage(strings.NewReader(soap))
	if err == nil {
		t.Error("expected error for invalid Action element")
	}
}

func TestParseSOAPMessage_InvalidActionElementWithInvalidUTF8(t *testing.T) {
	soap := `
		<Envelope>
			<Header>
				<Action>Invalid UTF-8: ` + string([]byte{0xFF, 0xFE, 0xFD}) + `</Action>
			</Header>
			<Body>
				<AuthorizeRequest>
					<idTag>ABC123</idTag>
				</AuthorizeRequest>
			</Body>
		</Envelope>`

	_, err := core.ParseSOAPMessage(strings.NewReader(soap))
	if err == nil {
		t.Error("expected error for invalid UTF-8 in Action element")
	}
}

func TestParseSOAPMessage_InvalidActionElementWithNestedContent(t *testing.T) {
	soap := `
		<Envelope>
			<Header>
				<Action>
					<nested>Invalid nested content</nested>
				</Action>
			</Header>
			<Body>
				<AuthorizeRequest>
					<idTag>ABC123</idTag>
				</AuthorizeRequest>
			</Body>
		</Envelope>`

	_, err := core.ParseSOAPMessage(strings.NewReader(soap))
	if err == nil {
		t.Error("expected error for invalid Action element with nested content")
	}
}

func TestParseSOAPMessage_MissingAction(t *testing.T) {
	soap := `
		<Envelope>
			<Header>
			</Header>
			<Body>
				<AuthorizeRequest>
					<idTag>ABC123</idTag>
				</AuthorizeRequest>
			</Body>
		</Envelope>`

	_, err := core.ParseSOAPMessage(strings.NewReader(soap))
	if err == nil || err.Error() != "missing Action in SOAP header" {
		t.Errorf("expected missing Action error, got: %v", err)
	}
}

func TestParseSOAPMessage_MissingBody(t *testing.T) {
	soap := `
		<Envelope>
			<Header>
				<Action>Authorize</Action>
			</Header>
		</Envelope>`

	_, err := core.ParseSOAPMessage(strings.NewReader(soap))
	if err == nil || err.Error() != "missing or empty SOAP Body" {
		t.Errorf("expected missing or empty SOAP Body error, got: %v", err)
	}
}

func TestParseSOAPMessage_InvalidXML(t *testing.T) {
	invalid := `<Envelope><Header><Action>Authorize</Action></Header><Body><AuthorizeRequest><idTag>ABC123</idTag></AuthorizeRequest>`

	_, err := core.ParseSOAPMessage(strings.NewReader(invalid))
	if err == nil {
		t.Error("expected error for invalid XML")
	}
}
