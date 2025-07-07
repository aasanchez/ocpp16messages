package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/shared/types"
)

func TestSetInteger(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		value         string
		want          uint32
		wantErr       bool
	}{
		{"ValidInteger", "123", 123, false},
		{"Zero", "0", 0, false},
		{"MaxUint32", "4294967295", 4294967295, false},
		{"NegativeInteger", "-1", 0, true},
		{"NotANumber", "abc", 0, true},
		{"EmptyString", "", 0, true},
		{"TooLarge", "4294967296", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			integer, err := types.SetInteger(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetInteger() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !tt.wantErr && integer.Value() != tt.want {
				t.Errorf("SetInteger() = %v, want %v", integer.Value(), tt.want)
			}
		})
	}
}

func TestInteger_Value(t *testing.T) {
	t.Parallel()
	integer, _ := types.SetInteger("42")
	if integer.Value() != 42 {
		t.Errorf("Value() = %v, want 42", integer.Value())
	}
}

func TestInteger_Validate(t *testing.T) {
	t.Parallel()
	integer, _ := types.SetInteger("123")
	if err := integer.Validate(); err != nil {
		t.Errorf("Validate() error = %v, want nil", err)
	}
}
