package getlocallistversion_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getlocallistversion"
	types "github.com/aasanchez/ocpp16types"
)

func TestReq_Valid_EmptyInput(t *testing.T) {
	t.Parallel()

	_, err := getlocallistversion.Req(getlocallistversion.ReqInput{})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}
}

func TestReq_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	req, err := getlocallistversion.Req(getlocallistversion.ReqInput{})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	// Verify the returned message is the zero value
	expected := getlocallistversion.ReqMessage{}
	if req != expected {
		t.Errorf(types.ErrorMismatchValue, expected, req)
	}
}
