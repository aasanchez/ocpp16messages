package sharedtypes_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// CiString20Type (max 20)
// OCPP 1.6: IdTag, short tokens, compact labels.
// Valid exact-bound length is accepted.
func Test_sharedtypes_SetCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 20)

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Empty is allowed by the base type. Final
// allowance depends on the message schema.
func Test_sharedtypes_SetCiString20Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Overflow (>20) is rejected with an error.
func Test_sharedtypes_SetCiString20Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 21)

	_, err := st.SetCiString20Type(val)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Value() round-trip returns the original input.
func Test_sharedtypes_SetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, _ := st.SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrorValueMismatch, input, ciStr.Value())
	}
}

// CiString25Type (max 25)
// OCPP 1.6: vendor/model-like descriptors.
// Valid exact-bound length is accepted.
func Test_sharedtypes_SetCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 25)

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Empty allowed at base type; schema may forbid.
func Test_sharedtypes_SetCiString25Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Overflow (>25) is rejected with an error.
func Test_sharedtypes_SetCiString25Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 26)

	_, err := st.SetCiString25Type(val)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Value() round-trip returns the original input.
func Test_sharedtypes_SetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, _ := st.SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrorValueMismatch, input, ciStr.Value())
	}
}

// CiString50Type (max 50)
// OCPP 1.6: firmware version, serials, labels.
// Valid exact-bound length is accepted.
func Test_sharedtypes_SetCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 50)

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Empty allowed at base type; schema may forbid.
func Test_sharedtypes_SetCiString50Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Overflow (>50) is rejected with an error.
func Test_sharedtypes_SetCiString50Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 51)

	_, err := st.SetCiString50Type(val)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Value() round-trip returns the original input.
func Test_sharedtypes_SetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, _ := st.SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrorValueMismatch, input, ciStr.Value())
	}
}

// CiString255Type (max 255)
// OCPP 1.6: descriptions, diagnostics, vendor text.
// Valid exact-bound length is accepted.
func Test_sharedtypes_SetCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 255)

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Empty allowed at base type; schema may forbid.
func Test_sharedtypes_SetCiString255Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Overflow (>255) is rejected with an error.
func Test_sharedtypes_SetCiString255Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 256)

	_, err := st.SetCiString255Type(val)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Value() round-trip returns the original input.
func Test_sharedtypes_SetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, _ := st.SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrorValueMismatch, input, ciStr.Value())
	}
}

// CiString500Type (max 500)
// OCPP 1.6: extended diagnostics, long ASCII notes.
// Valid exact-bound length is accepted.
func Test_sharedtypes_SetCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 500)

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Empty allowed at base type; schema may forbid.
func Test_sharedtypes_SetCiString500Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Overflow (>500) is rejected with an error.
func Test_sharedtypes_SetCiString500Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 501)

	_, err := st.SetCiString500Type(val)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Value() round-trip returns the original input.
func Test_sharedtypes_SetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, _ := st.SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrorValueMismatch, input, ciStr.Value())
	}
}
