package meterValues_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/meterValues"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestConf_Success(t *testing.T) {
	t.Parallel()

	input := meterValues.ConfInput{}

	_, err := meterValues.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestConf_ReturnsEmptyMessage(t *testing.T) {
	t.Parallel()

	input := meterValues.ConfInput{}

	conf, err := meterValues.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expected := meterValues.ConfMessage{}
	if conf != expected {
		t.Errorf(st.ErrorMismatchValue, expected, conf)
	}
}
