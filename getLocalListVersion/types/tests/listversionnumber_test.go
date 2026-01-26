package types_test

import (
	"errors"
	"math"
	"strconv"
	"testing"

	mt "github.com/aasanchez/ocpp16messages/getLocalListVersion/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValuePositive    = 5
	testValueZero        = 0
	testValueNegativeOne = -1
	testValueNegativeTen = -10
	testValueOverflow    = math.MaxInt32 + 1

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
		t.Errorf(
			st.ErrorMismatchValue,
			expectedValuePositive,
			listVersion.Value(),
		)
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

	_, err := mt.NewListVersionNumber(testValueOverflow)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "overflow value")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
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

	if listVersion.String() != strconv.Itoa(testValuePositive) {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ListVersionNumber.String()",
			listVersion.String(),
			strconv.Itoa(testValuePositive),
		)
	}
}

func TestListVersionNumber_String_NegativeOne(t *testing.T) {
	t.Parallel()

	listVersion, err := mt.NewListVersionNumber(testValueNegativeOne)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if listVersion.String() != strconv.Itoa(testValueNegativeOne) {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ListVersionNumber.String()",
			listVersion.String(),
			strconv.Itoa(testValueNegativeOne),
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
