package getlocallistversion_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getlocallistversion"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestReq_Valid_EmptyInput(t *testing.T) {
	t.Parallel()

	_, err := getlocallistversion.Req(getlocallistversion.ReqInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestReq_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	req, err := getlocallistversion.Req(getlocallistversion.ReqInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	// Verify the returned message is the zero value
	expected := getlocallistversion.ReqMessage{}
	if req != expected {
		t.Errorf(st.ErrorMismatchValue, expected, req)
	}
}
