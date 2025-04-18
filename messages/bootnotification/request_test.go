package bootnotification

// import (
// 	"strings"
// 	"testing"

// 	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
// )

// // ---- Test helpers ----
// func mustCiString20(val string) sharedtypes.CiString20Type {
// 	c, err := sharedtypes.CiString20(val)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return c
// }
// func mustCiString25(val string) sharedtypes.CiString25Type {
// 	c, err := sharedtypes.CiString25(val)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return c
// }
// func mustCiString50(val string) sharedtypes.CiString50Type {
// 	c, err := sharedtypes.CiString50(val)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return c
// }
// func ptr[T any](v T) *T { return &v }

// // ---- Tests ----

// // SonarCloud false positive: underscores allowed in Go test names
// func TestRequestAllValidFields(t *testing.T) {
// 	t.Parallel()

// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "CBoxSN-123",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "CPSN-456",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "1.2.3.4-rc",
// 		Iccid:                   "8988",
// 		Imsi:                    "1234567890",
// 		MeterSerialNumber:       "MeterSN-001",
// 		MeterType:               "SmartMeter",
// 	}
// 	req, err := Request(input)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	type fieldCheck struct {
// 		name     string
// 		got      string
// 		expected string
// 	}

// 	checks := []fieldCheck{
// 		{"ChargeBoxSerialNumber", req.ChargeBoxSerialNumber.String(), "CBoxSN-123"},
// 		{"ChargePointModel", req.ChargePointModel.String(), "CPModelX"},
// 		{"ChargePointSerialNumber", req.ChargePointSerialNumber.String(), "CPSN-456"},
// 		{"ChargePointVendor", req.ChargePointVendor.String(), "CPVendorZ"},
// 		{"FirmwareVersion", req.FirmwareVersion.String(), "1.2.3.4-rc"},
// 		{"Iccid", req.Iccid.String(), "8988"},
// 		{"Imsi", req.Imsi.String(), "1234567890"},
// 		{"MeterSerialNumber", req.MeterSerialNumber.String(), "MeterSN-001"},
// 		{"MeterType", req.MeterType.String(), "SmartMeter"},
// 	}

// 	for _, check := range checks {
// 		if check.got != check.expected {
// 			t.Errorf("unexpected %s: got %q, want %q", check.name, check.got, check.expected)
// 		}
// 	}
// }

// func TestRequestOnlyRequiredFields(t *testing.T) {
// 	t.Parallel()

// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	req, err := Request(input)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if req.ChargeBoxSerialNumber != nil ||
// 		req.ChargePointSerialNumber != nil ||
// 		req.FirmwareVersion != nil ||
// 		req.Iccid != nil ||
// 		req.Imsi != nil ||
// 		req.MeterSerialNumber != nil ||
// 		req.MeterType != nil {
// 		t.Errorf("expected all optional fields to be nil, got: %+v", req)
// 	}
// }

// func TestRequestErrorsChargePointModel(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 21) // > 20
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        invalid,
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setChargePointModel") {
// 		t.Errorf("expected setChargePointModel error, got %v", err)
// 	}
// }

// func TestRequestErrorsChargePointVendor(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 21)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       invalid,
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setChargePointVendor") {
// 		t.Errorf("expected setChargePointVendor error, got %v", err)
// 	}
// }

// func TestRequestErrorsChargeBoxSerialNumber(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 26)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   invalid,
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setChargeBoxSerialNumber") {
// 		t.Errorf("expected setChargeBoxSerialNumber error, got %v", err)
// 	}
// }

// func TestRequestErrorsChargePointSerialNumber(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 26)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: invalid,
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setChargePointSerialNumber") {
// 		t.Errorf("expected setChargePointSerialNumber error, got %v", err)
// 	}
// }

// func TestRequestErrorsFirmwareVersion(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 51)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         invalid,
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setFirmwareVersion") {
// 		t.Errorf("expected setFirmwareVersion error, got %v", err)
// 	}
// }

// func TestRequestErrorsIccid(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 21)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   invalid,
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setIccid") {
// 		t.Errorf("expected setIccid error, got %v", err)
// 	}
// }

// func TestRequestErrorsImsi(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 21)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    invalid,
// 		MeterSerialNumber:       "",
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setImsi") {
// 		t.Errorf("expected setImsi error, got %v", err)
// 	}
// }

// func TestRequestErrorsMeterSerialNumber(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 26)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       invalid,
// 		MeterType:               "",
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setMeterSerialNumber") {
// 		t.Errorf("expected setMeterSerialNumber error, got %v", err)
// 	}
// }

// func TestRequestErrorsMeterType(t *testing.T) {
// 	t.Parallel()

// 	invalid := strings.Repeat("X", 26)
// 	input := bootNotificationReqInput{
// 		ChargeBoxSerialNumber:   "",
// 		ChargePointModel:        "CPModelX",
// 		ChargePointSerialNumber: "",
// 		ChargePointVendor:       "CPVendorZ",
// 		FirmwareVersion:         "",
// 		Iccid:                   "",
// 		Imsi:                    "",
// 		MeterSerialNumber:       "",
// 		MeterType:               invalid,
// 	}
// 	_, err := Request(input)

// 	if err == nil || !strings.Contains(err.Error(), "setMeterType") {
// 		t.Errorf("expected setMeterType error, got %v", err)
// 	}
// }

// func TestRequestMessageValidateOk(t *testing.T) {
// 	t.Parallel()

// 	msg := RequestMessage{
// 		ChargeBoxSerialNumber:   nil,
// 		ChargePointModel:        mustCiString20("OnlyModel"),
// 		ChargePointSerialNumber: nil,
// 		ChargePointVendor:       mustCiString20("OnlyVendor"),
// 		FirmwareVersion:         nil,
// 		Iccid:                   nil,
// 		Imsi:                    nil,
// 		MeterSerialNumber:       nil,
// 		MeterType:               nil,
// 	}

// 	if err := msg.Validate(); err != nil {
// 		t.Errorf("unexpected Validate() error: %v", err)
// 	}
// }

// func TestRequestMessageValidateErrorChargePointModel(t *testing.T) {
// 	t.Parallel()

// 	invalidModel, _ := sharedtypes.CiString20(strings.Repeat("X", 21))
// 	msg := RequestMessage{
// 		ChargeBoxSerialNumber:   nil,
// 		ChargePointModel:        invalidModel,
// 		ChargePointSerialNumber: nil,
// 		ChargePointVendor:       mustCiString20("validVendor"),
// 		FirmwareVersion:         nil,
// 		Iccid:                   nil,
// 		Imsi:                    nil,
// 		MeterSerialNumber:       nil,
// 		MeterType:               nil,
// 	}
// 	err := msg.Validate()

// 	if err == nil || !strings.Contains(err.Error(), "invalid ChargePointModel") {
// 		t.Errorf("expected invalid ChargePointModel error, got: %v", err)
// 	}
// }

// func TestRequestMessageValidateErrorChargePointVendor(t *testing.T) {
// 	t.Parallel()

// 	invalidVendor, _ := sharedtypes.CiString20(strings.Repeat("X", 21))
// 	msg := RequestMessage{
// 		ChargeBoxSerialNumber:   nil,
// 		ChargePointModel:        mustCiString20("validModel"),
// 		ChargePointSerialNumber: nil,
// 		ChargePointVendor:       invalidVendor,
// 		FirmwareVersion:         nil,
// 		Iccid:                   nil,
// 		Imsi:                    nil,
// 		MeterSerialNumber:       nil,
// 		MeterType:               nil,
// 	}
// 	err := msg.Validate()

// 	if err == nil || !strings.Contains(err.Error(), "invalid ChargePointVendor") {
// 		t.Errorf("expected invalid ChargePointVendor error, got: %v", err)
// 	}
// }

// func TestRequestMessageStringAllOptionalsSet(t *testing.T) {
// 	t.Parallel()

// 	msg := RequestMessage{
// 		ChargeBoxSerialNumber:   ptr(mustCiString25("BoxSN")),
// 		ChargePointModel:        mustCiString20("ModelX"),
// 		ChargePointSerialNumber: ptr(mustCiString25("CPSN")),
// 		ChargePointVendor:       mustCiString20("VendorY"),
// 		FirmwareVersion:         ptr(mustCiString50("0.2")),
// 		Iccid:                   ptr(mustCiString20("ICCID")),
// 		Imsi:                    ptr(mustCiString20("IMSI")),
// 		MeterSerialNumber:       ptr(mustCiString25("MeterSN")),
// 		MeterType:               ptr(mustCiString25("TypeX")),
// 	}
// 	out := msg.String()

// 	for _, want := range []string{"BoxSN", "ModelX", "CPSN", "VendorY", "0.2", "ICCID", "IMSI", "MeterSN", "TypeX"} {
// 		if !strings.Contains(out, want) {
// 			t.Errorf("expected %q in String(), got: %s", want, out)
// 		}
// 	}
// }

// func TestRequestMessageStringNoOptionalsSet(t *testing.T) {
// 	t.Parallel()

// 	msg := RequestMessage{
// 		ChargeBoxSerialNumber:   nil,
// 		ChargePointModel:        mustCiString20("OnlyModel"),
// 		ChargePointSerialNumber: nil,
// 		ChargePointVendor:       mustCiString20("OnlyVendor"),
// 		FirmwareVersion:         nil,
// 		Iccid:                   nil,
// 		Imsi:                    nil,
// 		MeterSerialNumber:       nil,
// 		MeterType:               nil,
// 	}
// 	out := msg.String()

// 	for _, notWant := range []string{
// 		"chargeBoxSerialNumber",
// 		"chargePointSerialNumber",
// 		"firmwareVersion",
// 		"iccid",
// 		"imsi",
// 		"meterSerialNumber",
// 		"meterType",
// 	} {
// 		if strings.Contains(out, notWant) {
// 			t.Errorf("did not expect %q in output: %s", notWant, out)
// 		}
// 	}
// }
