package getDiagnostics

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a GetDiagnostics.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Optional: Name of the file with diagnostic information that will be
	// uploaded. If not present, no file is available.
	FileName *string
}

// ConfMessage represents an OCPP 1.6 GetDiagnostics.conf message.
type ConfMessage struct {
	FileName *st.CiString255Type
}

// Conf creates a GetDiagnostics.conf message from the given input.
// It validates all fields and returns an error if:
//   - FileName (if provided) exceeds 255 characters
//
// If FileName is not provided, it means no diagnostic information is available.
func Conf(input ConfInput) (ConfMessage, error) {
	fileName, err := confValidateFileName(input.FileName)
	if err != nil {
		return ConfMessage{}, err
	}

	return ConfMessage{
		FileName: fileName,
	}, nil
}

// confValidateFileName validates the optional file name field.
func confValidateFileName(fileName *string) (*st.CiString255Type, error) {
	if fileName == nil {
		return nil, nil
	}

	fn, err := st.NewCiString255Type(*fileName)
	if err != nil {
		return nil, fmt.Errorf("fileName: %w", err)
	}

	return &fn, nil
}
