package bootNotification_test

import (
	"strings"
	"testing"

	bn "github.com/aasanchez/ocpp16messages/messages/bootNotification"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	vendorABC    = "VendorABC"
	modelXYZ     = "ModelXYZ"
	errVendor    = "chargePointVendor"
	errModel     = "chargePointModel"
	errSerial    = "chargePointSerialNumber"
	errFirmware  = "firmwareVersion"
	errIccid     = "iccid"
	longString21 = "123456789012345678901"
	longString26 = "12345678901234567890123456"
	longString51 = "123456789012345678901234567890123456789012345678901"
)

func TestReq_Valid_RequiredOnly(t *testing.T) {
	t.Parallel()

	req, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       vendorABC,
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ChargePointVendor.Value() != vendorABC {
		t.Errorf(st.ErrorMismatch, vendorABC, req.ChargePointVendor.Value())
	}

	if req.ChargePointModel.Value() != modelXYZ {
		t.Errorf(st.ErrorMismatch, modelXYZ, req.ChargePointModel.Value())
	}
}

func TestReq_Valid_AllFields(t *testing.T) {
	t.Parallel()

	serial := "SN12345"
	chargeBox := "CB12345"
	firmware := "1.0.0"
	iccid := "89012345678901234567"
	imsi := "310150123456789"
	meterType := "ABB"
	meterSerial := "MS12345"

	req, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       vendorABC,
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: &serial,
		ChargeBoxSerialNumber:   &chargeBox,
		FirmwareVersion:         &firmware,
		Iccid:                   &iccid,
		Imsi:                    &imsi,
		MeterType:               &meterType,
		MeterSerialNumber:       &meterSerial,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	verifyAllFieldsNotNil(t, req)
}

func verifyAllFieldsNotNil(t *testing.T, req bn.ReqMessage) {
	t.Helper()

	if req.ChargePointSerialNumber == nil {
		t.Errorf(st.ErrorWantNonNil, "ChargePointSerialNumber")
	}

	if req.ChargeBoxSerialNumber == nil {
		t.Errorf(st.ErrorWantNonNil, "ChargeBoxSerialNumber")
	}

	if req.FirmwareVersion == nil {
		t.Errorf(st.ErrorWantNonNil, "FirmwareVersion")
	}

	if req.Iccid == nil {
		t.Errorf(st.ErrorWantNonNil, "Iccid")
	}

	if req.Imsi == nil {
		t.Errorf(st.ErrorWantNonNil, "Imsi")
	}

	if req.MeterType == nil {
		t.Errorf(st.ErrorWantNonNil, "MeterType")
	}

	if req.MeterSerialNumber == nil {
		t.Errorf(st.ErrorWantNonNil, "MeterSerialNumber")
	}
}

func TestReq_EmptyVendor(t *testing.T) {
	t.Parallel()

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       "",
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty vendor")
	}

	if !strings.Contains(err.Error(), errVendor) {
		t.Errorf(st.ErrorWantContains, err, errVendor)
	}
}

func TestReq_EmptyModel(t *testing.T) {
	t.Parallel()

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       vendorABC,
		ChargePointModel:        "",
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty model")
	}

	if !strings.Contains(err.Error(), errModel) {
		t.Errorf(st.ErrorWantContains, err, errModel)
	}
}

func TestReq_VendorTooLong(t *testing.T) {
	t.Parallel()

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       longString21,
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "vendor too long")
	}

	if !strings.Contains(err.Error(), errVendor) {
		t.Errorf(st.ErrorWantContains, err, errVendor)
	}
}

func TestReq_ModelTooLong(t *testing.T) {
	t.Parallel()

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       vendorABC,
		ChargePointModel:        longString21,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "model too long")
	}

	if !strings.Contains(err.Error(), errModel) {
		t.Errorf(st.ErrorWantContains, err, errModel)
	}
}

func TestReq_SerialNumberTooLong(t *testing.T) {
	t.Parallel()

	serial := longString26

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       vendorABC,
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: &serial,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "serial number too long")
	}

	if !strings.Contains(err.Error(), errSerial) {
		t.Errorf(st.ErrorWantContains, err, errSerial)
	}
}

func TestReq_FirmwareVersionTooLong(t *testing.T) {
	t.Parallel()

	firmware := longString51

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       vendorABC,
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         &firmware,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "firmware version too long")
	}

	if !strings.Contains(err.Error(), errFirmware) {
		t.Errorf(st.ErrorWantContains, err, errFirmware)
	}
}

func TestReq_IccidTooLong(t *testing.T) {
	t.Parallel()

	iccid := longString21

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       vendorABC,
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   &iccid,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "iccid too long")
	}

	if !strings.Contains(err.Error(), errIccid) {
		t.Errorf(st.ErrorWantContains, err, errIccid)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	serial := longString26
	iccid := longString21

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       "",
		ChargePointModel:        "",
		ChargePointSerialNumber: &serial,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   &iccid,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want multiple errors")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errVendor) {
		t.Errorf(st.ErrorWantContains, err, errVendor)
	}

	if !strings.Contains(errStr, errModel) {
		t.Errorf(st.ErrorWantContains, err, errModel)
	}

	if !strings.Contains(errStr, errSerial) {
		t.Errorf(st.ErrorWantContains, err, errSerial)
	}

	if !strings.Contains(errStr, errIccid) {
		t.Errorf(st.ErrorWantContains, err, errIccid)
	}
}

func TestReq_InvalidCharacters(t *testing.T) {
	t.Parallel()

	_, err := bn.Req(bn.ReqInput{
		ChargePointVendor:       "Vendor\x00ABC",
		ChargePointModel:        modelXYZ,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid characters in vendor")
	}

	if !strings.Contains(err.Error(), errVendor) {
		t.Errorf(st.ErrorWantContains, err, errVendor)
	}
}
