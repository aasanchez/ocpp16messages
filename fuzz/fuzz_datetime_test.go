//go:build fuzz

package fuzz

import (
    "testing"
    "time"

    st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewDateTime(f *testing.F) {
    f.Add("2025-01-15T10:30:00Z")
    f.Add("not-a-time")
    f.Add("")
    f.Add("2025-01-15T10:30:00+02:00")

    f.Fuzz(func(t *testing.T, input string) {
        dt, err := st.NewDateTime(input)
        if err != nil {
            return
        }

        if dt.Value().Location() != time.UTC {
            t.Fatalf("location = %v, want UTC", dt.Value().Location())
        }

        roundTrip := dt.String()
        if roundTrip == "" {
            t.Fatalf("String() returned empty for input %q", input)
        }

        if _, parseErr := time.Parse(time.RFC3339Nano, roundTrip); parseErr != nil {
            t.Fatalf("String() not RFC3339Nano: %v", parseErr)
        }
    })
}
