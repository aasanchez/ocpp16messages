package sharedtypes

import (
	"testing"
)

const testIntegerValue = 73

func Test_sharedtypes_SetInteger(t *testing.T) {
	t.Parallel()

	i, err := SetInteger("73")
	if err != nil {
		t.Fatalf("unexpected error from SetInteger: %v", err)
	}

	got := i.Value()
	if got != testIntegerValue {
		t.Errorf("Value() = %d; want %d", got, testIntegerValue)
	}
}

func Test_sharedtypes_SetInteger_Overflow(t *testing.T) {
	t.Parallel()

	_, err := SetInteger("4294967296")
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetInteger_Negative(t *testing.T) {
	t.Parallel()

	_, err := SetInteger("-10")
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetInteger_Alphanumeric(t *testing.T) {
	t.Parallel()

	_, err := SetInteger("abc")
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}
