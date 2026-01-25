package updateFirmware_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
	"github.com/aasanchez/ocpp16messages/updateFirmware"
)

func TestConf_Success(t *testing.T) {
	t.Parallel()

	input := updateFirmware.ConfInput{}

	_, err := updateFirmware.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestConf_ReturnsEmptyMessage(t *testing.T) {
	t.Parallel()

	input := updateFirmware.ConfInput{}

	conf, err := updateFirmware.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expected := updateFirmware.ConfMessage{}
	if conf != expected {
		t.Errorf(st.ErrorMismatchValue, expected, conf)
	}
}
