package sharedtypes_test

import (
	"errors"
	"testing"
	"unicode/utf8"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ciStr is a minimal interface for fuzz tests: Validate and Value.
type ciStr interface {
	Validate() error
	Value() string
}

// fuzzCiString encapsulates common fuzz logic for CiString types.
// addSeeds should call f.Add(...) to register seed inputs.
// setFn should wrap st.SetCiStringXXType, returning the typed CiString.
func fuzzCiString(f *testing.F, addSeeds func(*testing.F), setFn func(string) (ciStr, error)) {
	addSeeds(f)
	f.Fuzz(func(t *testing.T, data string) {
		if !utf8.ValidString(data) {
			return
		}

		cs, err := setFn(data)
		if err != nil {
			if !errors.Is(err, st.ErrExceedsMaxLength) && !errors.Is(err, st.ErrNonPrintableASCII) {
				t.Fatalf("unexpected error type: %v", err)
			}

			return
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("Validate() returned error for input %q: %v", data, err)
		}

		if got := cs.Value(); got != data {
			t.Fatalf("round-trip failed: input %q, got %q", data, got)
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("second Validate() returned error for input %q: %v", data, err)
		}

		cs2, err := setFn(cs.Value())
		if err != nil {
			t.Fatalf("idempotent Set failed for %q: %v", cs.Value(), err)
		}

		if cs2.Value() != cs.Value() {
			t.Fatalf("idempotent Set mismatch: got %q, want %q", cs2.Value(), cs.Value())
		}
	})
}

// FuzzSetCiString20Type ensures that SetCiString20Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
func FuzzSetCiString20Type(f *testing.F) {
	add := func(f *testing.F) {
		f.Add("valid-string")
		f.Add("12345678901234567890")
		f.Add("a-string-that-is-just-over-20-chars")
		f.Add("")
		f.Add("string-with-non-printable-\t-char")
		f.Add("string-with-€-symbol")
		f.Add("\u001f")
		f.Add("\x1f")
	}
	fuzzCiString(f, add, func(s string) (ciStr, error) {
		c, err := st.SetCiString20Type(s)

		return c, err
	})
}

// FuzzSetCiString25Type ensures that SetCiString25Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data and checks
// for round-trip consistency and validation invariants.
func FuzzSetCiString25Type(f *testing.F) {
	add := func(f *testing.F) {
		f.Add("a-valid-string")
		f.Add("1234567890123456789012345")
		f.Add("a-string-that-is-just-over-25-characters")
		f.Add("")
		f.Add("string-with-non-printable-\t-char")
		f.Add("string-with-€-symbol")
		f.Add("\u001f")
		f.Add("\x1f")
	}
	fuzzCiString(f, add, func(s string) (ciStr, error) {
		c, err := st.SetCiString25Type(s)

		return c, err
	})
}

// FuzzSetCiString50Type ensures that SetCiString50Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data and checks
// for round-trip consistency and validation invariants.
func FuzzSetCiString50Type(f *testing.F) {
	// Add seed values to guide the fuzzer.
	f.Add("a-valid-string")
		f.Add("12345678901234567890123456789012345678901234567890")
		f.Add("a-string-that-is-just-a-little-bit-over-50-characters-long")
		f.Add("")
		f.Add("string-with-non-printable-\t-char")
		f.Add("string-with-€-symbol")
		f.Add("\u001f")
		f.Add("\x1f")

	f.Fuzz(func(t *testing.T, data string) {
		if !utf8.ValidString(data) {
			return
		}

		cs, err := st.SetCiString50Type(data)
		if err != nil {
			if !errors.Is(err, st.ErrExceedsMaxLength) && !errors.Is(err, st.ErrNonPrintableASCII) {
				t.Fatalf("unexpected error type: %v", err)
			}

			return
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("Validate() returned error for input %q: %v", data, err)
		}

		if got := cs.Value(); got != data {
			t.Fatalf("round-trip failed: input %q, got %q", data, got)
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("second Validate() returned error for input %q: %v", data, err)
		}

		cs2, err := st.SetCiString50Type(cs.Value())
		if err != nil {
			t.Fatalf("idempotent Set failed for %q: %v", cs.Value(), err)
		}

		if cs2.Value() != cs.Value() {
			t.Fatalf("idempotent Set mismatch: got %q, want %q", cs2.Value(), cs.Value())
		}
	})
}

// FuzzSetCiString255Type ensures that SetCiString255Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data and checks
// for round-trip consistency and validation invariants.
func FuzzSetCiString255Type(f *testing.F) {
	// Add seed values to guide the fuzzer.
	f.Add("a-valid-string")
	f.Add("a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long")
	f.Add("")
	f.Add("string-with-non-printable-	-char")
	f.Add("\u001f")
	f.Add("\x1f")

	f.Fuzz(func(t *testing.T, data string) {
		if !utf8.ValidString(data) {
			return
		}

		cs, err := st.SetCiString255Type(data)
		if err != nil {
			if !errors.Is(err, st.ErrExceedsMaxLength) && !errors.Is(err, st.ErrNonPrintableASCII) {
				t.Fatalf("unexpected error type: %v", err)
			}

			return
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("Validate() returned error for input %q: %v", data, err)
		}

		if got := cs.Value(); got != data {
			t.Fatalf("round-trip failed: input %q, got %q", data, got)
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("second Validate() returned error for input %q: %v", data, err)
		}

		cs2, err := st.SetCiString255Type(cs.Value())
		if err != nil {
			t.Fatalf("idempotent Set failed for %q: %v", cs.Value(), err)
		}

		if cs2.Value() != cs.Value() {
			t.Fatalf("idempotent Set mismatch: got %q, want %q", cs2.Value(), cs.Value())
		}
	})
}

// FuzzSetCiString500Type ensures that SetCiString500Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data and checks
// for round-trip consistency and validation invariants.
func FuzzSetCiString500Type(f *testing.F) {
	// Add seed values to guide the fuzzer.
	f.Add("a-valid-string")
	f.Add("a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long")
	f.Add("")
	f.Add("string-with-non-printable-	-char")
	f.Add("\u001f")
	f.Add("\x1f")

	f.Fuzz(func(t *testing.T, data string) {
		if !utf8.ValidString(data) {
			return
		}

		cs, err := st.SetCiString500Type(data)
		if err != nil {
			if !errors.Is(err, st.ErrExceedsMaxLength) && !errors.Is(err, st.ErrNonPrintableASCII) {
				t.Fatalf("unexpected error type: %v", err)
			}

			return
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("Validate() returned error for input %q: %v", data, err)
		}

		if got := cs.Value(); got != data {
			t.Fatalf("round-trip failed: input %q, got %q", data, got)
		}

		if err := cs.Validate(); err != nil {
			t.Fatalf("second Validate() returned error for input %q: %v", data, err)
		}

		cs2, err := st.SetCiString500Type(cs.Value())
		if err != nil {
			t.Fatalf("idempotent Set failed for %q: %v", cs.Value(), err)
		}

		if cs2.Value() != cs.Value() {
			t.Fatalf("idempotent Set mismatch: got %q, want %q", cs2.Value(), cs.Value())
		}
	})
}
