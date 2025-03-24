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
		inHeader  bool
		inBody    bool
		action    string
		payload   []byte
		innerXML  strings.Builder
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
			switch {
			case el.Name.Local == "Header":
				inHeader = true
			case el.Name.Local == "Body":
				inBody = true
				foundBody = true
			case inHeader && el.Name.Local == "Action":
				var value string
				if err := decoder.DecodeElement(&value, &el); err != nil {
					return nil, err
				}
				value = strings.TrimSpace(value)
				if !isValidAction(value) {
					return nil, errors.New("invalid Action format")
				}
				action = value
				inHeader = false
			case inBody && depth == 0:
				// capture the inner payload XML
				start := el
				innerXML.WriteString("<" + start.Name.Local)
				for _, attr := range start.Attr {
					innerXML.WriteString(" " + attr.Name.Local + `="` + attr.Value + `"`)
				}
				innerXML.WriteString(">")
				depth++
			default:
				if inBody && depth > 0 {
					// include nested elements
					raw, _ := xml.Marshal(el)
					innerXML.Write(raw)
					depth++
				}
			}
		case xml.EndElement:
			switch {
			case el.Name.Local == "Header":
				inHeader = false
			case el.Name.Local == "Body":
				inBody = false
			case inBody && depth > 0:
				innerXML.WriteString("</" + el.Name.Local + ">")
				depth--
			}
		case xml.CharData:
			if inBody && depth > 0 {
				innerXML.Write(el)
			}
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
