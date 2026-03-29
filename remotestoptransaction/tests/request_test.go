package remotestoptransaction_test

import (
	"strings"
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remotestoptransaction"
	types "github.com/aasanchez/ocpp16types"
)

const (
	testTransactionIdValid = 12345
	testTransactionIdZero  = 0
	testTransactionIdMax   = 65535
	testTransactionIdOver  = 65536
	testTransactionIdNeg   = -1
	errTransactionId       = "transactionId"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionId: testTransactionIdValid})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionId.Value() != uint16(testTransactionIdValid) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testTransactionIdValid),
			req.TransactionId.Value(),
		)
	}
}

func TestReq_Valid_Zero(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionId: testTransactionIdZero})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionId.Value() != uint16(testTransactionIdZero) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testTransactionIdZero),
			req.TransactionId.Value(),
		)
	}
}

func TestReq_Valid_Max(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionId: testTransactionIdMax})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionId.Value() != uint16(testTransactionIdMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testTransactionIdMax),
			req.TransactionId.Value(),
		)
	}
}

func TestReq_TransactionIdNegative(t *testing.T) {
	t.Parallel()

	_, err := rst.Req(rst.ReqInput{TransactionId: testTransactionIdNeg})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative transactionId")
	}

	if !strings.Contains(err.Error(), errTransactionId) {
		t.Errorf(types.ErrorWantContains, err, errTransactionId)
	}
}

func TestReq_TransactionIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := rst.Req(rst.ReqInput{TransactionId: testTransactionIdOver})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "transactionId exceeds max")
	}

	if !strings.Contains(err.Error(), errTransactionId) {
		t.Errorf(types.ErrorWantContains, err, errTransactionId)
	}
}
