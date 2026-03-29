//go:build fuzz

package fuzz

import (
	"strconv"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzCiStringRoundTrip(f *testing.F) {
	f.Add("hello")
	f.Add("")
	f.Add("a")
	f.Add("12345678901234567890") // 20
	f.Add("1234567890123456789012345") // 25
	f.Add("12345678901234567890123456789012345678901234567890") // 50

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzLen {
			t.Skip()
		}

		// CiString20Type
		if ci20, err := types.NewCiString20Type(input); err == nil {
			if ci20.String() != input {
				t.Fatalf(
					"CiString20 String() = %q, want %q",
					ci20.String(), input,
				)
			}

			if ci20.Value() != input {
				t.Fatalf(
					"CiString20 Value() = %q, want %q",
					ci20.Value(), input,
				)
			}

			if ci20.String() != ci20.String() {
				t.Fatal("CiString20 String() not deterministic")
			}
		}

		// CiString25Type
		if ci25, err := types.NewCiString25Type(input); err == nil {
			if ci25.String() != input {
				t.Fatalf(
					"CiString25 String() = %q, want %q",
					ci25.String(), input,
				)
			}

			if ci25.String() != ci25.String() {
				t.Fatal("CiString25 String() not deterministic")
			}
		}

		// CiString50Type
		if ci50, err := types.NewCiString50Type(input); err == nil {
			if ci50.String() != input {
				t.Fatalf(
					"CiString50 String() = %q, want %q",
					ci50.String(), input,
				)
			}

			if ci50.String() != ci50.String() {
				t.Fatal("CiString50 String() not deterministic")
			}
		}

		// CiString255Type
		if ci255, err := types.NewCiString255Type(input); err == nil {
			if ci255.String() != input {
				t.Fatalf(
					"CiString255 String() = %q, want %q",
					ci255.String(), input,
				)
			}

			if ci255.String() != ci255.String() {
				t.Fatal(
					"CiString255 String() not deterministic",
				)
			}
		}

		// CiString500Type
		if ci500, err := types.NewCiString500Type(input); err == nil {
			if ci500.String() != input {
				t.Fatalf(
					"CiString500 String() = %q, want %q",
					ci500.String(), input,
				)
			}

			if ci500.String() != ci500.String() {
				t.Fatal(
					"CiString500 String() not deterministic",
				)
			}
		}
	})
}

func FuzzDateTimeRoundTrip(f *testing.F) {
	f.Add("2025-01-15T10:30:00Z")
	f.Add("2025-01-15T10:30:00+02:00")
	f.Add("2025-12-31T23:59:59.999999999Z")
	f.Add("")
	f.Add("not-a-time")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzLen {
			t.Skip()
		}

		dt, err := types.NewDateTime(input)
		if err != nil {
			return
		}

		// String() must be deterministic
		first := dt.String()
		second := dt.String()

		if first != second {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				first, second,
			)
		}

		// String() must parse as RFC3339Nano
		parsed, parseErr := time.Parse(
			time.RFC3339Nano, first,
		)
		if parseErr != nil {
			t.Fatalf(
				"String() %q not parseable: %v",
				first, parseErr,
			)
		}

		// Round-trip: reconstruct from String()
		dt2, rtErr := types.NewDateTime(first)
		if rtErr != nil {
			t.Fatalf(
				"round-trip failed: NewDateTime(%q) = %v",
				first, rtErr,
			)
		}

		if !dt2.Value().Equal(dt.Value()) {
			t.Fatalf(
				"round-trip value mismatch: %v vs %v",
				dt2.Value(), dt.Value(),
			)
		}

		// Must be in UTC
		if parsed.Location() != time.UTC {
			t.Fatalf(
				"parsed location = %v, want UTC",
				parsed.Location(),
			)
		}
	})
}

func FuzzIntegerRoundTrip(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(65535)
	f.Add(-1)
	f.Add(65536)

	f.Fuzz(func(t *testing.T, value int) {
		integer, err := types.NewInteger(value)
		if err != nil {
			return
		}

		// String() must be deterministic
		first := integer.String()
		second := integer.String()

		if first != second {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				first, second,
			)
		}

		// String() must parse back to same value
		parsed, parseErr := strconv.ParseUint(first, 10, 16)
		if parseErr != nil {
			t.Fatalf(
				"String() %q not parseable: %v",
				first, parseErr,
			)
		}

		if uint16(parsed) != integer.Value() {
			t.Fatalf(
				"round-trip: parsed %d != Value() %d",
				parsed, integer.Value(),
			)
		}

		// Round-trip: reconstruct from Value()
		integer2, rtErr := types.NewInteger(int(integer.Value()))
		if rtErr != nil {
			t.Fatalf(
				"round-trip NewInteger(%d) = %v",
				integer.Value(), rtErr,
			)
		}

		if integer2.Value() != integer.Value() {
			t.Fatalf(
				"round-trip value: %d != %d",
				integer2.Value(), integer.Value(),
			)
		}
	})
}

func FuzzListVersionNumberRoundTrip(f *testing.F) {
	f.Add(-1)
	f.Add(0)
	f.Add(1)
	f.Add(100)
	f.Add(2147483647) // max int32

	f.Fuzz(func(t *testing.T, value int) {
		lvn, err := types.NewListVersionNumber(value)
		if err != nil {
			return
		}

		// String() must be deterministic
		first := lvn.String()
		second := lvn.String()

		if first != second {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				first, second,
			)
		}

		// String() must parse back
		parsed, parseErr := strconv.ParseInt(first, 10, 32)
		if parseErr != nil {
			t.Fatalf(
				"String() %q not parseable: %v",
				first, parseErr,
			)
		}

		if int32(parsed) != lvn.Value() {
			t.Fatalf(
				"round-trip: parsed %d != Value() %d",
				parsed, lvn.Value(),
			)
		}

		// Semantic checks
		if lvn.Value() == types.ListVersionUnsupported &&
			!lvn.IsUnsupported() {
			t.Fatal(
				"Value() == -1 but IsUnsupported() = false",
			)
		}

		if lvn.Value() == types.ListVersionEmpty &&
			!lvn.IsEmpty() {
			t.Fatal("Value() == 0 but IsEmpty() = false")
		}

		if lvn.Value() > 0 &&
			(lvn.IsUnsupported() || lvn.IsEmpty()) {
			t.Fatalf(
				"Value() = %d but IsUnsupported/IsEmpty = true",
				lvn.Value(),
			)
		}
	})
}
