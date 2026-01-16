package types_test

import (
	"strings"
	"testing"

	mt "github.com/aasanchez/ocpp16messages/getLocalListVersion/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValuePositive    = "5"
	testValueZero        = "0"
	testValueNegativeOne = "-1"
	testValueNegativeTen = "-10"
	testValueInvalid     = "invalid"
	testValueOverflow    = "9999999999999999999"

	expectedValuePositive    = int32(5)
	expectedValueZero        = int32(0)
	expectedValueNegativeOne = int32(-1)
	expectedValueNegativeTen = int32(-10)
)

func TestNewListVersionNumber_Valid_Positive(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValuePositive)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.Value() != expectedValuePositive {
		t.Errorf(st.ErrorMismatchValue, expectedValuePositive, listVersion.Value())
	}
}

func TestNewListVersionNumber_Valid_Zero(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValueZero)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.Value() != expectedValueZero {
		t.Errorf(st.ErrorMismatchValue, expectedValueZero, listVersion.Value())
	}
}

func TestNewListVersionNumber_Valid_NegativeOne(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValueNegativeOne)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.Value() != expectedValueNegativeOne {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedValueNegativeOne,
			listVersion.Value(),
		)
	}
}

func TestNewListVersionNumber_Valid_NegativeTen(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValueNegativeTen)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.Value() != expectedValueNegativeTen {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedValueNegativeTen,
			listVersion.Value(),
		)
	}
}

func TestNewListVersionNumber_Invalid_NonNumeric(t *testing.T) {
	t.Parallel()

	_, err := mt.NewListVersionNumber(testValueInvalid)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "non-numeric value")
	}
}

func TestNewListVersionNumber_Invalid_Overflow(t *testing.T) {
	t.Parallel()

	_, err := mt.NewListVersionNumber(testValueOverflow)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "overflow value")
	}

	if !strings.Contains(err.Error(), "value out of range") {
		t.Errorf(st.ErrorWantContains, err, "value out of range")
	}
}

func TestListVersionNumber_IsUnsupported_True(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValueNegativeOne)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if !listVersion.IsUnsupported() {
		t.Error("IsUnsupported() = false, want true")
	}
}

func TestListVersionNumber_IsUnsupported_False(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValuePositive)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.IsUnsupported() {
		t.Error("IsUnsupported() = true, want false")
	}
}

func TestListVersionNumber_IsEmpty_True(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValueZero)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if !listVersion.IsEmpty() {
		t.Error("IsEmpty() = false, want true")
	}
}

func TestListVersionNumber_IsEmpty_False(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValuePositive)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.IsEmpty() {
		t.Error("IsEmpty() = true, want false")
	}
}

func TestListVersionNumber_String_Positive(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValuePositive)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.String() != testValuePositive {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ListVersionNumber.String()",
			listVersion.String(),
			testValuePositive,
		)
	}
}

func TestListVersionNumber_String_NegativeOne(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValueNegativeOne)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.String() != testValueNegativeOne {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ListVersionNumber.String()",
			listVersion.String(),
			testValueNegativeOne,
		)
	}
}

func TestListVersionNumber_Value_ConstantUnsupported(t *testing.T) {
	t.Parallel()

	if mt.ListVersionUnsupported != expectedValueNegativeOne {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedValueNegativeOne,
			mt.ListVersionUnsupported,
		)
	}
}

func TestListVersionNumber_Value_ConstantEmpty(t *testing.T) {
	t.Parallel()

	if mt.ListVersionEmpty != expectedValueZero {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedValueZero,
			mt.ListVersionEmpty,
		)
	}
}
