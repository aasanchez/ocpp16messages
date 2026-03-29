package reset_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/reset"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errType = "type"
)

func TestReq_Valid_Hard(t *testing.T) {
	t.Parallel()

	req, err := reset.Req(reset.ReqInput{Type: "Hard"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Type != types.ResetTypeHard {
		t.Errorf(
			types.ErrorMismatch,
			types.ResetTypeHard,
			req.Type,
		)
	}
}

func TestReq_Valid_Soft(t *testing.T) {
	t.Parallel()

	req, err := reset.Req(reset.ReqInput{Type: "Soft"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Type != types.ResetTypeSoft {
		t.Errorf(
			types.ErrorMismatch,
			types.ResetTypeSoft,
			req.Type,
		)
	}
}

func TestReq_EmptyType(t *testing.T) {
	t.Parallel()

	_, err := reset.Req(reset.ReqInput{Type: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Unknown(t *testing.T) {
	t.Parallel()

	_, err := reset.Req(reset.ReqInput{Type: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := reset.Req(reset.ReqInput{Type: "hard"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}
