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
			_, err := types.SetInteger(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetInteger() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}
