package sharedtypes_fuzz

import (
   "errors"
   "testing"
   "unicode/utf8"

   st "github.com/aasanchez/ocpp16messages/shared/types"
)

// FuzzSetCiString20Type ensures that SetCiString20Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data and checks
// for round-trip consistency and validation invariants.
func FuzzSetCiString20Type(f *testing.F) {
	// Add seed values to guide the fuzzer.
	// These seeds include typical valid, boundary, and invalid inputs.
	f.Add("valid-string")
		f.Add("12345678901234567890")
		f.Add("a-string-that-is-just-over-20-chars")
		f.Add("")
		f.Add("string-with-non-printable-\t-char")
		f.Add("string-with-€-symbol")
		f.Add("\u001f")
		f.Add("")

	f.Fuzz(func(t *testing.T, data string) {
		// The underlying validation logic expects ASCII characters, so we should ensure
		// the fuzzed input is valid UTF-8 before proceeding.
		if !utf8.ValidString(data) {
			return
		}

		cs, err := st.SetCiString20Type(data)
		if err != nil {
			// Contract: only known validation errors are allowed
			if !errors.Is(err, st.ErrExceedsMaxLength) && !errors.Is(err, st.ErrNonPrintableASCII) {
				t.Fatalf("unexpected error type: %v", err)
			}
			return
		}
		// Invariant: Validate should succeed for a valid instance
		if err := cs.Validate(); err != nil {
			t.Fatalf("Validate() returned error for input %q: %v", data, err)
		}
		// Round-Trip: Value must equal original input
		if got := cs.Value(); got != data {
			t.Fatalf("round-trip failed: input %q, got %q", data, got)
		}
		// Idempotence: repeated Validate should still succeed
		if err := cs.Validate(); err != nil {
			t.Fatalf("second Validate() returned error for input %q: %v", data, err)
		}
		// Idempotence: setting from Value should yield equivalent instance
		cs2, err := st.SetCiString20Type(cs.Value())
		if err != nil {
			t.Fatalf("idempotent Set failed for %q: %v", cs.Value(), err)
		}
		if cs2.Value() != cs.Value() {
			t.Fatalf("idempotent Set mismatch: got %q, want %q", cs2.Value(), cs.Value())
		}
	})
}

// FuzzSetCiString25Type ensures that SetCiString25Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data and checks
// for round-trip consistency and validation invariants.
func FuzzSetCiString25Type(f *testing.F) {
	// Add seed values to guide the fuzzer.
	f.Add("a-valid-string")
	f.Add("1234567890123456789012345")                     // 25 chars, at the limit
	f.Add("a-string-that-is-just-over-25-characters")      // over 25 chars
	f.Add("")                                              // empty string
	f.Add("string-with-non-printable-	-char") // contains tab
	f.Add("string-with-€-symbol")                          // non-ASCII
	f.Add("")                                          // non-printable control character
	f.Add("")                                          // non-printable control character

	f.Fuzz(func(t *testing.T, data string) {
		if !utf8.ValidString(data) {
			return
		}

		cs, err := st.SetCiString25Type(data)
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
		cs2, err := st.SetCiString25Type(cs.Value())
		if err != nil {
			t.Fatalf("idempotent Set failed for %q: %v", cs.Value(), err)
		}
		if cs2.Value() != cs.Value() {
			t.Fatalf("idempotent Set mismatch: got %q, want %q", cs2.Value(), cs.Value())
		}
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
		f.Add("")

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
	f.Add("")
	f.Add("")

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
	f.Add("")
	f.Add("")

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
