//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	dt "github.com/aasanchez/ocpp16messages/dataTransfer"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzDataTransferReq(f *testing.F) {
	f.Add("VendorId", false, "", false, "")
	f.Add("", false, "", false, "")
	f.Add("VendorId", true, "MessageId", true, "payload")
	f.Add("VendorId", true, "", false, "")

	f.Fuzz(func(
		t *testing.T,
		vendorId string,
		hasMessageId bool,
		messageId string,
		hasData bool,
		data string,
	) {
		if len(vendorId) > maxFuzzStringLen ||
			len(messageId) > maxFuzzStringLen ||
			len(data) > maxFuzzStringLen {
			t.Skip()
		}

		var messageIdPtr *string
		if hasMessageId {
			messageIdPtr = &messageId
		}

		var dataPtr *string
		if hasData {
			dataPtr = &data
		}

		req, err := dt.Req(dt.ReqInput{
			VendorId:  vendorId,
			MessageId: messageIdPtr,
			Data:      dataPtr,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if vendorId == "" {
			t.Fatal("Req succeeded with empty VendorId")
		}

		if req.VendorId.String() != vendorId {
			t.Fatalf("VendorId = %q, want %q", req.VendorId.String(), vendorId)
		}

		if hasMessageId {
			if req.MessageId == nil {
				t.Fatal("MessageId = nil, want non-nil")
			}
			if messageId == "" {
				t.Fatal("Req succeeded with empty MessageId")
			}
			if req.MessageId.String() != messageId {
				t.Fatalf(
					"MessageId = %q, want %q",
					req.MessageId.String(),
					messageId,
				)
			}
		} else if req.MessageId != nil {
			t.Fatal("MessageId != nil, want nil")
		}

		if hasData {
			if req.Data == nil {
				t.Fatal("Data = nil, want non-nil")
			}
			if *req.Data != data {
				t.Fatalf("Data = %q, want %q", *req.Data, data)
			}
		} else if req.Data != nil {
			t.Fatal("Data != nil, want nil")
		}
	})
}
