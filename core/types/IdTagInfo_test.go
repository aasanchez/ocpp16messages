package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

func TestNewIdTagInfo(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		expectErr bool
		expected  *types.IdToken
		errorMsg  string
	}{
		{
			name:      "Valid IdTag",
			value:     "validIdTag123",
			expectErr: false,
			expected: &types.IdToken{
				CiString20: types.CiString20("validIdTag123"),
			},
			errorMsg: "",
		},
		{
			name:      "IdTag exceeds 20 characters",
			value:     "thisIsAVeryLongIdTagThatExceeds20Characters",
			expectErr: true,
			expected:  nil,
			errorMsg:  "invalid field 'idToken': idToken exceeds 20 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idTagInfo, err := types.NewIdTagInfo(tt.value)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("expected error message to be %v, got %v", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if idTagInfo == nil || idTagInfo.CiString20 != tt.expected.CiString20 {
					t.Errorf("expected %v, got %v", tt.expected, idTagInfo)
				}
			}
		})
	}
}

func TestIdTagInfoFields(t *testing.T) {
	// Valid IdTagInfo instance
	idTagInfo := types.IdTagInfo{
		ExpiryDate:  nil,
		ParentIdTag: nil,
		Status:      "Active",
	}

	// Check Status field
	if idTagInfo.Status != "Active" {
		t.Errorf("expected status to be 'Active', got %v", idTagInfo.Status)
	}

	// Check ExpiryDate and ParentIdTag fields (both should be nil)
	if idTagInfo.ExpiryDate != nil {
		t.Errorf("expected ExpiryDate to be nil, got %v", idTagInfo.ExpiryDate)
	}
	if idTagInfo.ParentIdTag != nil {
		t.Errorf("expected ParentIdTag to be nil, got %v", idTagInfo.ParentIdTag)
	}

	// Check if the struct is properly initialized and non-empty
	if idTagInfo == (types.IdTagInfo{}) {
		t.Error("expected IdTagInfo to be non-empty")
	}
}
