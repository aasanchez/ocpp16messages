package core

import (
	"encoding/xml"
	"errors"
	"io"
	"regexp"
	"strings"
)

// ParsedSOAPMessage represents a parsed SOAP message with essential fields.
type ParsedSOAPMessage struct {
	TypeID  MessageType // Always CALL for SOAP requests
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

	var (
		action    string
		payload   []byte
		innerXML  strings.Builder
		inHeader  bool
		inBody    bool
		depth     int
		foundBody bool
	)

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
			err := handleStartElement(el, &inHeader, &inBody, &foundBody, &depth, &innerXML, decoder, &action)
			if err != nil {
				return nil, err
			}

		case xml.EndElement:
			handleEndElement(el, &inHeader, &inBody, &depth, &innerXML)

		case xml.CharData:
			handleCharData(el, &inBody, &depth, &innerXML)
		}
	}

	if action == "" {
		return nil, errors.New("missing Action in SOAP header")
	}
	if !foundBody || innerXML.Len() == 0 {
		return nil, errors.New("missing or empty SOAP Body")
	}

	payload = []byte(innerXML.String())

	return &ParsedSOAPMessage{
		TypeID:  CALL, // SOAP messages are treated as CALL type
		Action:  action,
		Payload: payload,
	}, nil
}

// handleStartElement processes the start element of the XML and updates the state.
func handleStartElement(el xml.StartElement, inHeader, inBody, foundBody *bool, depth *int, innerXML *strings.Builder, decoder *xml.Decoder, action *string) error {
	switch {
	case el.Name.Local == "Header":
		*inHeader = true
	case el.Name.Local == "Body":
		*inBody = true
		*foundBody = true
	case *inHeader && el.Name.Local == "Action":
		return extractAction(el, decoder, action)
	case *inBody && *depth == 0:
		// capture the inner payload XML
		start := el
		innerXML.WriteString("<" + start.Name.Local)
		for _, attr := range start.Attr {
			innerXML.WriteString(" " + attr.Name.Local + `="` + attr.Value + `"`)
		}
		innerXML.WriteString(">")
		*depth++
	default:
		if *inBody && *depth > 0 {
			// include nested elements
			raw, _ := xml.Marshal(el)
			innerXML.Write(raw)
			*depth++
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
func handleEndElement(el xml.EndElement, inHeader, inBody *bool, depth *int, innerXML *strings.Builder) {
	switch {
	case el.Name.Local == "Header":
		*inHeader = false
	case el.Name.Local == "Body":
		*inBody = false
	case *inBody && *depth > 0:
		innerXML.WriteString("</" + el.Name.Local + ">")
		*depth--
	}
}

// handleCharData processes the character data inside the body of the SOAP message.
func handleCharData(el xml.CharData, inBody *bool, depth *int, innerXML *strings.Builder) {
	if *inBody && *depth > 0 {
		innerXML.Write(el)
	}
}
