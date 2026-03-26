package metervalues_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/metervalues"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestConf_Success(t *testing.T) {
	t.Parallel()

	input := metervalues.ConfInput{}

	_, err := metervalues.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestConf_ReturnsEmptyMessage(t *testing.T) {
	t.Parallel()

	input := metervalues.ConfInput{}

	conf, err := metervalues.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expected := metervalues.ConfMessage{}
	if conf != expected {
		t.Errorf(st.ErrorMismatchValue, expected, conf)
	}
}
