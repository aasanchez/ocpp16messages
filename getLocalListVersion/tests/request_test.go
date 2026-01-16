package getLocalListVersion_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getLocalListVersion"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestReq_Valid_EmptyInput(t *testing.T) {
	t.Parallel()

	_, err := getLocalListVersion.Req(getLocalListVersion.ReqInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestReq_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	req, err := getLocalListVersion.Req(getLocalListVersion.ReqInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	// Verify the returned message is the zero value
	expected := getLocalListVersion.ReqMessage{}
	if req != expected {
		t.Errorf(st.ErrorMismatchValue, expected, req)
	}
}
