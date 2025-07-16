package sharedtypes_fuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// FuzzSetCiString20Type ensures that SetCiString20Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data.
func FuzzSetCiString20Type(f *testing.F) {
	// Add some seed values to guide the fuzzer.
	// These seeds include typical valid and invalid inputs.
	f.Add("valid-string")
	f.Add("a-string-that-is-just-over-20-chars")
	f.Add("")
	f.Add("string-with-non-printable-	-char")

	// Run the fuzzer.
	f.Fuzz(func(t *testing.T, data string) {
		// The function being fuzzed is SetCiString20Type.
		// We don't need to check the returned values because the goal
		// of this fuzz test is to find inputs that cause panics or crashes.
		// The existing unit tests already cover the expected outputs.
		_, _ = st.SetCiString20Type(data)
	})
}

// FuzzSetCiString25Type ensures that SetCiString25Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data.
func FuzzSetCiString25Type(f *testing.F) {
	// Add some seed values to guide the fuzzer.
	f.Add("a-valid-string")
	f.Add("a-string-that-is-just-over-25-characters")
	f.Add("")
	f.Add("string-with-non-printable-	-char")

	// Run the fuzzer.
	f.Fuzz(func(t *testing.T, data string) {
		// The function being fuzzed is SetCiString25Type.
		// We are looking for inputs that cause the program to panic.
		_, _ = st.SetCiString25Type(data)
	})
}

// FuzzSetCiString50Type ensures that SetCiString50Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data.
func FuzzSetCiString50Type(f *testing.F) {
	// Add some seed values to guide the fuzzer.
	f.Add("a-valid-string")
	f.Add("a-string-that-is-just-a-little-bit-over-50-characters-long")
	f.Add("")
	f.Add("string-with-non-printable-	-char")

	// Run the fuzzer.
	f.Fuzz(func(t *testing.T, data string) {
		// The function being fuzzed is SetCiString50Type.
		// We are looking for inputs that cause the program to panic.
		_, _ = st.SetCiString50Type(data)
	})
}

// FuzzSetCiString255Type ensures that SetCiString255Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data.
func FuzzSetCiString255Type(f *testing.F) {
	// Add some seed values to guide the fuzzer.
	f.Add("a-valid-string")
	f.Add("a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long-a-string-that-is-just-a-little-bit-over-255-characters-long")
	f.Add("")
	f.Add("string-with-non-printable-	-char")

	// Run the fuzzer.
	f.Fuzz(func(t *testing.T, data string) {
		// The function being fuzzed is SetCiString255Type.
		// We are looking for inputs that cause the program to panic.
		_, _ = st.SetCiString255Type(data)
	})
}

// FuzzSetCiString500Type ensures that SetCiString500Type can handle a wide
// range of string inputs without crashing or producing unexpected behavior.
// It fuzzes the function by providing various random string data.
func FuzzSetCiString500Type(f *testing.F) {
	// Add some seed values to guide the fuzzer.
	f.Add("a-valid-string")
	f.Add("a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long-a-string-that-is-just-a-little-bit-over-500-characters-long")
	f.Add("")
	f.Add("string-with-non-printable-	-char")

	// Run the fuzzer.
	f.Fuzz(func(t *testing.T, data string) {
		// The function being fuzzed is SetCiString500Type.
		// We are looking for inputs that cause the program to panic.
		_, _ = st.SetCiString500Type(data)
	})
}
