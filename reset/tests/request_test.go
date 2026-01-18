package reset_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/reset"
	rt "github.com/aasanchez/ocpp16messages/reset/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errType = "type"
)

func TestReq_Valid_Hard(t *testing.T) {
	t.Parallel()

	req, err := reset.Req(reset.ReqInput{Type: "Hard"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Type != rt.ResetTypeHard {
		t.Errorf(
			st.ErrorMismatch,
			rt.ResetTypeHard,
			req.Type,
		)
	}
}

func TestReq_Valid_Soft(t *testing.T) {
	t.Parallel()

	req, err := reset.Req(reset.ReqInput{Type: "Soft"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Type != rt.ResetTypeSoft {
		t.Errorf(
			st.ErrorMismatch,
			rt.ResetTypeSoft,
			req.Type,
		)
	}
}

func TestReq_EmptyType(t *testing.T) {
	t.Parallel()

	_, err := reset.Req(reset.ReqInput{Type: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(st.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Unknown(t *testing.T) {
	t.Parallel()

	_, err := reset.Req(reset.ReqInput{Type: "Unknown"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(st.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := reset.Req(reset.ReqInput{Type: "hard"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(st.ErrorWantContains, err, errType)
	}
}
