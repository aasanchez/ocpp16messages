package core

import (
	"encoding/xml"
	"errors"
	"io"
	"regexp"
	"strings"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

// ParsedSOAPMessage represents a parsed SOAP message with essential fields.
type ParsedSOAPMessage struct {
	TypeID  types.MessageType // Always CALL for SOAP requests
	Action  string
	Payload []byte
}

// Regex pattern for valid Action values (alphanumeric, underscore, dash)
var validActionPattern = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

// isValidAction checks if the Action value matches the expected format.
func isValidAction(action string) bool {
	return validActionPattern.MatchString(action)
}

// ParseSOAPMessage parses a SOAP XML message and extracts the Action and Payload.
func ParseSOAPMessage(r io.Reader) (*ParsedSOAPMessage, error) {
	decoder := xml.NewDecoder(r)

	state := SOAPParserState{
		inHeader:  false,
		inBody:    false,
		foundBody: false,
		depth:     0,
	}

	var innerXML strings.Builder
	var action string

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch el := tok.(type) {
		case xml.StartElement:
			err := handleStartElement(el, &state, decoder, &action, &innerXML)
			if err != nil {
				return nil, err
			}

		case xml.EndElement:
			handleEndElement(el, &state, &innerXML)

		case xml.CharData:
			handleCharData(el, &state, &innerXML)
		}
	}

	if action == "" {
		return nil, errors.New("missing Action in SOAP header")
	}
	if !state.foundBody || innerXML.Len() == 0 {
		return nil, errors.New("missing or empty SOAP Body")
	}

	payload := []byte(innerXML.String())

	return &ParsedSOAPMessage{
		TypeID:  types.CALL, // SOAP messages are treated as CALL type
		Action:  action,
		Payload: payload,
	}, nil
}

// SOAPParserState holds the state for parsing a SOAP message.
type SOAPParserState struct {
	inHeader  bool
	inBody    bool
	foundBody bool
	depth     int
}

// handleStartElement processes the start element of the XML and updates the state.
func handleStartElement(el xml.StartElement, state *SOAPParserState, decoder *xml.Decoder, action *string, innerXML *strings.Builder) error {
	switch {
	case el.Name.Local == "Header":
		state.inHeader = true
	case el.Name.Local == "Body":
		state.inBody = true
		state.foundBody = true
	case state.inHeader && el.Name.Local == "Action":
		return extractAction(el, decoder, action)
	case state.inBody && state.depth == 0:
		// capture the inner payload XML
		start := el
		innerXML.WriteString("<" + start.Name.Local)
		for _, attr := range start.Attr {
			innerXML.WriteString(" " + attr.Name.Local + `="` + attr.Value + `"`)
		}
		innerXML.WriteString(">")
		state.depth++
	default:
		if state.inBody && state.depth > 0 {
			// include nested elements
			raw, _ := xml.Marshal(el)
			innerXML.Write(raw)
			state.depth++
		}
	}
	return nil
}

// extractAction extracts the Action value from the XML and validates it.
func extractAction(el xml.StartElement, decoder *xml.Decoder, action *string) error {
	var value string
	if err := decoder.DecodeElement(&value, &el); err != nil {
		return err
	}
	value = strings.TrimSpace(value)
	if !isValidAction(value) {
		return errors.New("invalid Action format")
	}
	*action = value
	return nil
}

// handleEndElement processes the end element of the XML.
func handleEndElement(el xml.EndElement, state *SOAPParserState, innerXML *strings.Builder) {
	switch {
	case el.Name.Local == "Header":
		state.inHeader = false
	case el.Name.Local == "Body":
		state.inBody = false
	case state.inBody && state.depth > 0:
		innerXML.WriteString("</" + el.Name.Local + ">")
		state.depth--
	}
}

// handleCharData processes the character data inside the body of the SOAP message.
func handleCharData(el xml.CharData, state *SOAPParserState, innerXML *strings.Builder) {
	if state.inBody && state.depth > 0 {
		innerXML.Write(el)
	}
}
