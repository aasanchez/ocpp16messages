package sharedtypes_test

import (
	"errors"
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// TestSetDateTime_Valid exercises valid ISO-8601 date-time strings.
func TestSetDateTime_Valid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string // expected canonical output via String()
	}{
		{"2013-02-01T20:00:00.000Z", "2013-02-01T20:00:00Z"},
		{"2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		{"2006-01-02T15:04:05.123Z", "2006-01-02T15:04:05.123Z"},
		{"2006-01-02T15:04:05.123456789Z", "2006-01-02T15:04:05.123456789Z"},
		{"2006-01-02T15:04:05+07:00", "2006-01-02T15:04:05+07:00"},
		{"2006-01-02T15:04:05.123+02:00", "2006-01-02T15:04:05.123+02:00"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()
			dt, err := st.SetDateTime(tc.input)
			if err != nil {
				t.Fatalf("SetDateTime(%q) returned error: %v", tc.input, err)
			}
			out := dt.String()
			if out != tc.expected {
				t.Errorf("String() = %q; want %q", out, tc.expected)
			}
		})
	}
}

func TestSetDateTime_Invalid_Blank(t *testing.T) {
	t.Parallel()

	_, err := st.SetDateTime("")
	if err == nil {
		t.Fatal("expected error for blank input, got nil")
	}
}

func TestSetDateTime_Invalid_Format(t *testing.T) {
	t.Parallel()

	cases := []string{"not-a-date", "2006-01-02 15:04:05Z"}
	for _, s := range cases {
		t.Run(s, func(t *testing.T) {
			t.Parallel()
		var err error
		_, err = st.SetDateTime(s)
		if err == nil {
			t.Fatalf("expected error for %q", s)
		}
		})
	}
}

func TestSetDateTime_Invalid_OutOfRange(t *testing.T) {
	t.Parallel()

	cases := []string{"2006-01-02T25:00:00Z", "2006-13-02T15:04:05Z", "2006-01-32T15:04:05Z"}
	for _, s := range cases {
		t.Run(s, func(t *testing.T) {
			t.Parallel()
			_, err := st.SetDateTime(s)
			if err == nil {
				t.Fatalf("expected range error for %q", s)
			}
			var pe *time.ParseError
			if !errors.As(err, &pe) {
				t.Fatalf("expected ParseError for %q; got %T", s, err)
			}
		})
	}
}
