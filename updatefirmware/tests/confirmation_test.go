package updatefirmware_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
	"github.com/aasanchez/ocpp16messages/updatefirmware"
)

func TestConf_Success(t *testing.T) {
	t.Parallel()

	input := updatefirmware.ConfInput{}

	_, err := updatefirmware.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestConf_ReturnsEmptyMessage(t *testing.T) {
	t.Parallel()

	input := updatefirmware.ConfInput{}

	conf, err := updatefirmware.Conf(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expected := updatefirmware.ConfMessage{}
	if conf != expected {
		t.Errorf(st.ErrorMismatchValue, expected, conf)
	}
}
