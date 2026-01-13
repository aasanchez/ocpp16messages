package clearCache_test

import (
	"testing"

	cc "github.com/aasanchez/ocpp16messages/messages/clearCache"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestReq_Valid_EmptyInput(t *testing.T) {
	t.Parallel()

	_, err := cc.Req(cc.ReqInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestReq_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	req, err := cc.Req(cc.ReqInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	// Verify the returned message is the zero value
	expected := cc.ReqMessage{}
	if req != expected {
		t.Errorf(st.ErrorMismatchValue, expected, req)
	}
}
