package sendlocallist_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/sendlocallist"
	slt "github.com/aasanchez/ocpp16messages/sendlocallist/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errListVersion            = "ListVersion"
	errUpdateType             = "UpdateType"
	errLocalAuthorizationList = "localAuthorizationList"
	validIdTag                = "RFID12345"
	validStatus               = "Accepted"
	listVersionZero           = 0
	listVersionOne            = 1
	listVersionNegative       = -1
	expectedLenZero           = 0
	expectedLenOne            = 1
	expectedLenThree          = 3
)

func TestReq_Valid_Full_EmptyList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.UpdateType != slt.UpdateTypeFull {
		t.Errorf(
			st.ErrorMismatch,
			slt.UpdateTypeFull,
			req.UpdateType,
		)
	}
}

func TestReq_Valid_Differential_EmptyList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "Differential",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.UpdateType != slt.UpdateTypeDifferential {
		t.Errorf(
			st.ErrorMismatch,
			slt.UpdateTypeDifferential,
			req.UpdateType,
		)
	}
}

func TestReq_Valid_ListVersionZero(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionZero,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ListVersion.Value() != listVersionZero {
		t.Errorf(
			st.ErrorMismatchValue,
			listVersionZero,
			req.ListVersion.Value(),
		)
	}
}

func TestReq_Valid_WithAuthorizationList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []slt.AuthorizationDataInput{
			{
				IdTag: validIdTag,
				IdTagInfo: &slt.IdTagInfoInput{
					Status:      validStatus,
					ExpiryDate:  nil,
					ParentIdTag: nil,
				},
			},
		},
		UpdateType: "Full",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(req.LocalAuthorizationList) != expectedLenOne {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedLenOne,
			len(req.LocalAuthorizationList),
		)
	}
}

func TestReq_Valid_WithMultipleAuthorizationEntries(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []slt.AuthorizationDataInput{
			{
				IdTag: "RFID001",
				IdTagInfo: &slt.IdTagInfoInput{
					Status:      "Accepted",
					ExpiryDate:  nil,
					ParentIdTag: nil,
				},
			},
			{
				IdTag: "RFID002",
				IdTagInfo: &slt.IdTagInfoInput{
					Status:      "Blocked",
					ExpiryDate:  nil,
					ParentIdTag: nil,
				},
			},
			{
				IdTag:     "RFID003",
				IdTagInfo: nil,
			},
		},
		UpdateType: "Differential",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(req.LocalAuthorizationList) != expectedLenThree {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedLenThree,
			len(req.LocalAuthorizationList),
		)
	}
}

func TestReq_Valid_WithEmptyAuthorizationList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: []slt.AuthorizationDataInput{},
		UpdateType:             "Full",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.LocalAuthorizationList == nil {
		t.Error("LocalAuthorizationList = nil, want empty slice")
	}

	if len(req.LocalAuthorizationList) != expectedLenZero {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedLenZero,
			len(req.LocalAuthorizationList),
		)
	}
}

func TestReq_InvalidListVersion_Negative(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionNegative,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative list version")
	}

	if !strings.Contains(err.Error(), errListVersion) {
		t.Errorf(st.ErrorWantContains, err, errListVersion)
	}
}

func TestReq_InvalidUpdateType_Empty(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty update type")
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(st.ErrorWantContains, err, errUpdateType)
	}
}

func TestReq_InvalidUpdateType_Unknown(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "Unknown",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown update type")
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(st.ErrorWantContains, err, errUpdateType)
	}
}

func TestReq_InvalidUpdateType_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "full",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase update type")
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(st.ErrorWantContains, err, errUpdateType)
	}
}

func TestReq_InvalidAuthorizationList_EmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []slt.AuthorizationDataInput{
			{IdTag: "", IdTagInfo: nil},
		},
		UpdateType: "Full",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty idTag in list")
	}

	if !strings.Contains(err.Error(), errLocalAuthorizationList) {
		t.Errorf(st.ErrorWantContains, err, errLocalAuthorizationList)
	}
}

func TestReq_InvalidAuthorizationList_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []slt.AuthorizationDataInput{
			{
				IdTag: validIdTag,
				IdTagInfo: &slt.IdTagInfoInput{
					Status:      "InvalidStatus",
					ExpiryDate:  nil,
					ParentIdTag: nil,
				},
			},
		},
		UpdateType: "Full",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid status in list")
	}

	if !strings.Contains(err.Error(), errLocalAuthorizationList) {
		t.Errorf(st.ErrorWantContains, err, errLocalAuthorizationList)
	}
}

func TestReq_MultipleErrors_Accumulated(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionNegative,
		LocalAuthorizationList: nil,
		UpdateType:             "Invalid",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), errListVersion) {
		t.Errorf(st.ErrorWantContains, err, errListVersion)
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(st.ErrorWantContains, err, errUpdateType)
	}
}
