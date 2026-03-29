package stoptransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/stoptransaction"
	types "github.com/aasanchez/ocpp16types"
)

const (
	testValidIdTag          = "RFID-TAG-12345"
	testValidTimestamp      = "2025-01-15T10:30:00Z"
	testTransactionId12345  = 12345
	testMeterStop5000       = 5000
	testValueZero           = 0
	testValueNegativeOne    = -1
	testTxDataLenZero       = 0
	testExpectedTxDataLen   = 1
	errFieldTransactionId   = "transactionId"
	errFieldMeterStop       = "meterStop"
	errFieldTimestamp       = "timestamp"
	errFieldIdTag           = "idTag"
	errFieldReason          = "reason"
	errReqReasonNil         = "Req() Reason = nil, want non-nil"
	errReqIdTagNil          = "Req() IdTag = nil, want non-nil"
	errMsgExceedsMaxLen     = "exceeds maximum length"
	errMsgWantExceedsMaxLen = "Req() error = %v, want 'exceeds maximum length'"
	errMsgNonPrintable      = "non-printable ASCII"
	errMsgWantNonPrintable  = "Req() error = %v, want 'non-printable ASCII'"
	errMsgInvalidTxData     = "transactionData[0]"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	expectedTransactionId := uint16(testTransactionId12345)
	if req.TransactionId.Value() != expectedTransactionId {
		t.Errorf(
			types.ErrorMismatch,
			expectedTransactionId,
			req.TransactionId.Value(),
		)
	}

	expectedMeterStop := uint16(testMeterStop5000)
	if req.MeterStop.Value() != expectedMeterStop {
		t.Errorf(types.ErrorMismatch, expectedMeterStop, req.MeterStop.Value())
	}
}

func TestReq_ValidWithIdTag(t *testing.T) {
	t.Parallel()

	idTag := testValidIdTag

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           &idTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.IdTag == nil {
		t.Error(errReqIdTagNil)
	}

	if req.IdTag.String() != testValidIdTag {
		t.Errorf(types.ErrorMismatch, testValidIdTag, req.IdTag.String())
	}
}

func TestReq_ValidWithReasonLocal(t *testing.T) {
	t.Parallel()

	reason := "Local"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}

	if req.Reason.String() != reason {
		t.Errorf(types.ErrorMismatch, reason, req.Reason.String())
	}
}

func TestReq_ValidWithReasonRemote(t *testing.T) {
	t.Parallel()

	reason := "Remote"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}

	if req.Reason.String() != reason {
		t.Errorf(types.ErrorMismatch, reason, req.Reason.String())
	}
}

func TestReq_ValidWithReasonEVDisconnected(t *testing.T) {
	t.Parallel()

	reason := "EVDisconnected"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}

	if req.Reason.String() != reason {
		t.Errorf(types.ErrorMismatch, reason, req.Reason.String())
	}
}

func TestReq_ValidWithTransactionData(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId: testTransactionId12345,
		IdTag:         nil,
		MeterStop:     testMeterStop5000,
		Timestamp:     testValidTimestamp,
		Reason:        nil,
		TransactionData: []types.MeterValueInput{
			{
				Timestamp: testValidTimestamp,
				SampledValue: []types.SampledValueInput{
					{
						Value:     "5000",
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if len(req.TransactionData) != testExpectedTxDataLen {
		t.Errorf(
			types.ErrorMismatch,
			testExpectedTxDataLen,
			len(req.TransactionData),
		)
	}
}

func TestReq_ValidWithEmptyTransactionData(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: []types.MeterValueInput{},
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionData == nil {
		t.Error("TransactionData = nil, want empty slice")
	}

	if len(req.TransactionData) != testTxDataLenZero {
		t.Errorf(
			types.ErrorMismatch, testTxDataLenZero, len(req.TransactionData),
		)
	}
}

func TestReq_TransactionIdZero(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testValueZero,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionId.Value() != testValueZero {
		t.Errorf(types.ErrorMismatch, testValueZero, req.TransactionId.Value())
	}
}

func TestReq_TransactionIdNegative(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testValueNegativeOne,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative transactionId")
	}

	if !strings.Contains(err.Error(), errFieldTransactionId) {
		t.Errorf(types.ErrorWantContains, err, errFieldTransactionId)
	}
}

func TestReq_MeterStopZero(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testValueZero,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.MeterStop.Value() != testValueZero {
		t.Errorf(types.ErrorMismatch, testValueZero, req.MeterStop.Value())
	}
}

func TestReq_MeterStopNegative(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testValueNegativeOne,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative meterStop")
	}

	if !strings.Contains(err.Error(), errFieldMeterStop) {
		t.Errorf(types.ErrorWantContains, err, errFieldMeterStop)
	}
}

func TestReq_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       "not-a-timestamp",
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_EmptyTimestamp(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       "",
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_EmptyIdTag(t *testing.T) {
	t.Parallel()

	emptyTag := ""

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           &emptyTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty idTag")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_IdTagTooLong(t *testing.T) {
	t.Parallel()

	longTag := "RFID-ABC123456789012345" // 23 chars, max is 20

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           &longTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for IdTag too long")
	}

	if !strings.Contains(err.Error(), errMsgExceedsMaxLen) {
		t.Errorf(errMsgWantExceedsMaxLen, err)
	}
}

func TestReq_IdTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	invalidTag := "RFID\x00ABC"

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           &invalidTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for non-printable chars")
	}

	if !strings.Contains(err.Error(), errMsgNonPrintable) {
		t.Errorf(errMsgWantNonPrintable, err)
	}
}

func TestReq_InvalidReason(t *testing.T) {
	t.Parallel()

	invalidReason := "InvalidReason"

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &invalidReason,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid reason")
	}

	if !strings.Contains(err.Error(), errFieldReason) {
		t.Errorf(types.ErrorWantContains, err, errFieldReason)
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestReq_InvalidTransactionData(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId: testTransactionId12345,
		IdTag:         nil,
		MeterStop:     testMeterStop5000,
		Timestamp:     testValidTimestamp,
		Reason:        nil,
		TransactionData: []types.MeterValueInput{
			{
				Timestamp:    "invalid-timestamp",
				SampledValue: []types.SampledValueInput{},
			},
		},
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid transactionData")
	}

	if !strings.Contains(err.Error(), errMsgInvalidTxData) {
		t.Errorf(types.ErrorWantContains, err, errMsgInvalidTxData)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	emptyTag := ""
	invalidReason := "BadReason"

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testValueNegativeOne,
		IdTag:           &emptyTag,
		MeterStop:       testValueNegativeOne,
		Timestamp:       "invalid",
		Reason:          &invalidReason,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldTransactionId) {
		t.Errorf(types.ErrorWantContains, err, errFieldTransactionId)
	}

	if !strings.Contains(errStr, errFieldIdTag) {
		t.Errorf(types.ErrorWantContains, err, errFieldIdTag)
	}

	if !strings.Contains(errStr, errFieldMeterStop) {
		t.Errorf(types.ErrorWantContains, err, errFieldMeterStop)
	}

	if !strings.Contains(errStr, errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}

	if !strings.Contains(errStr, errFieldReason) {
		t.Errorf(types.ErrorWantContains, err, errFieldReason)
	}
}

func TestReq_Complete(t *testing.T) {
	t.Parallel()

	idTag := testValidIdTag
	reason := "Local"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   testTransactionId12345,
		IdTag:           &idTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	expectedTxId := uint16(testTransactionId12345)
	if req.TransactionId.Value() != expectedTxId {
		t.Errorf(types.ErrorMismatch, expectedTxId, req.TransactionId.Value())
	}

	if req.IdTag == nil {
		t.Error(errReqIdTagNil)
	}

	expectedMeterStop := uint16(testMeterStop5000)
	if req.MeterStop.Value() != expectedMeterStop {
		t.Errorf(types.ErrorMismatch, expectedMeterStop, req.MeterStop.Value())
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}
}
