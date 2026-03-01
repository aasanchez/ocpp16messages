//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	bn "github.com/aasanchez/ocpp16messages/bootNotification"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzBootNotificationReq(f *testing.F) {
	f.Add("Vendor", "Model", false, "", false, "")
	f.Add("", "Model", false, "", false, "")
	f.Add("Vendor", "", false, "", false, "")
	f.Add("Vendor", "Model", true, "", false, "")
	f.Add("Vendor", "Model", true, "SN-123", true, "FW-1.0")

	f.Fuzz(func(
		t *testing.T,
		vendor string,
		model string,
		hasChargePointSerialNumber bool,
		chargePointSerialNumber string,
		hasFirmwareVersion bool,
		firmwareVersion string,
	) {
		if len(vendor) > maxFuzzStringLen ||
			len(model) > maxFuzzStringLen ||
			len(chargePointSerialNumber) > maxFuzzStringLen ||
			len(firmwareVersion) > maxFuzzStringLen {
			t.Skip()
		}

		var chargePointSerialNumberPtr *string
		if hasChargePointSerialNumber {
			chargePointSerialNumberPtr = &chargePointSerialNumber
		}

		var firmwareVersionPtr *string
		if hasFirmwareVersion {
			firmwareVersionPtr = &firmwareVersion
		}

		req, err := bn.Req(bn.ReqInput{
			ChargePointVendor:       vendor,
			ChargePointModel:        model,
			ChargePointSerialNumber: chargePointSerialNumberPtr,
			ChargeBoxSerialNumber:   nil,
			FirmwareVersion:         firmwareVersionPtr,
			Iccid:                   nil,
			Imsi:                    nil,
			MeterType:               nil,
			MeterSerialNumber:       nil,
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

		if vendor == "" {
			t.Fatal("Req succeeded with empty ChargePointVendor")
		}

		if model == "" {
			t.Fatal("Req succeeded with empty ChargePointModel")
		}

		if req.ChargePointVendor.String() != vendor {
			t.Fatalf(
				"ChargePointVendor = %q, want %q",
				req.ChargePointVendor.String(),
				vendor,
			)
		}

		if req.ChargePointModel.String() != model {
			t.Fatalf(
				"ChargePointModel = %q, want %q",
				req.ChargePointModel.String(),
				model,
			)
		}

		if hasChargePointSerialNumber {
			if req.ChargePointSerialNumber == nil {
				t.Fatal("ChargePointSerialNumber = nil, want non-nil")
			}
			if chargePointSerialNumber == "" {
				t.Fatal("Req succeeded with empty ChargePointSerialNumber")
			}
			if req.ChargePointSerialNumber.String() != chargePointSerialNumber {
				t.Fatalf(
					"ChargePointSerialNumber = %q, want %q",
					req.ChargePointSerialNumber.String(),
					chargePointSerialNumber,
				)
			}
		} else if req.ChargePointSerialNumber != nil {
			t.Fatal("ChargePointSerialNumber != nil, want nil")
		}

		if hasFirmwareVersion {
			if req.FirmwareVersion == nil {
				t.Fatal("FirmwareVersion = nil, want non-nil")
			}
			if firmwareVersion == "" {
				t.Fatal("Req succeeded with empty FirmwareVersion")
			}
			if req.FirmwareVersion.String() != firmwareVersion {
				t.Fatalf(
					"FirmwareVersion = %q, want %q",
					req.FirmwareVersion.String(),
					firmwareVersion,
				)
			}
		} else if req.FirmwareVersion != nil {
			t.Fatal("FirmwareVersion != nil, want nil")
		}
	})
}
