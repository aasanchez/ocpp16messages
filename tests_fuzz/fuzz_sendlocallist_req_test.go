//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	sl "github.com/aasanchez/ocpp16messages/sendlocallist"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzSendLocalListReq(f *testing.F) {
	f.Add(0, types.UpdateTypeFull.String(), uint8(0), "")
	f.Add(1, types.UpdateTypeDifferential.String(), uint8(1), "")
	f.Add(1, types.UpdateTypeFull.String(), uint8(2), "RFID-ABC123")
	f.Add(-1, types.UpdateTypeFull.String(), uint8(0), "")
	f.Add(1, "invalid-update-type", uint8(0), "")
	f.Add(1, types.UpdateTypeFull.String(), uint8(2), "")

	f.Fuzz(func(
		t *testing.T,
		listVersion int,
		updateType string,
		listMode uint8,
		idTag string,
	) {
		if len(updateType) > maxFuzzStringLen || len(idTag) > maxFuzzStringLen {
			t.Skip()
		}

		var localAuthorizationList []types.AuthorizationDataInput

		switch listMode % 3 {
		case 0:
			localAuthorizationList = nil
		case 1:
			localAuthorizationList = []types.AuthorizationDataInput{}
		default:
			localAuthorizationList = []types.AuthorizationDataInput{
				{
					IdTag:     idTag,
					IdTagInfo: nil,
				},
			}
		}

		req, err := sl.Req(sl.ReqInput{
			ListVersion:            listVersion,
			LocalAuthorizationList: localAuthorizationList,
			UpdateType:             updateType,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if listVersion < 0 || listVersion > math.MaxUint16 {
			t.Fatalf("Req succeeded with listVersion=%d", listVersion)
		}

		if got := req.ListVersion.Value(); got != uint16(listVersion) {
			t.Fatalf("ListVersion = %d, want %d", got, listVersion)
		}

		if !req.UpdateType.IsValid() {
			t.Fatalf("UpdateType = %q, want valid", req.UpdateType.String())
		}

		switch listMode % 3 {
		case 0:
			if req.LocalAuthorizationList != nil {
				t.Fatal("LocalAuthorizationList != nil, want nil")
			}
		case 1:
			if req.LocalAuthorizationList == nil {
				t.Fatal("LocalAuthorizationList = nil, want empty slice")
			}
			if len(req.LocalAuthorizationList) != 0 {
				t.Fatalf(
					"len(LocalAuthorizationList) = %d, want 0",
					len(req.LocalAuthorizationList),
				)
			}
		default:
			if len(req.LocalAuthorizationList) == 0 {
				t.Fatal("LocalAuthorizationList is empty, want at least one")
			}
		}
	})
}
