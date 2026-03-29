package updatefirmware_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/updatefirmware"
	types "github.com/aasanchez/ocpp16types"
)

func TestConf_Success(t *testing.T) {
	t.Parallel()

	input := updatefirmware.ConfInput{}

	_, err := updatefirmware.Conf(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}
}

func TestConf_ReturnsEmptyMessage(t *testing.T) {
	t.Parallel()

	input := updatefirmware.ConfInput{}

	conf, err := updatefirmware.Conf(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	expected := updatefirmware.ConfMessage{}
	if conf != expected {
		t.Errorf(types.ErrorMismatchValue, expected, conf)
	}
}
