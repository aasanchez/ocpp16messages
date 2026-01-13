package clearChargingProfile_test

import (
	"strings"
	"testing"

	ccp "github.com/aasanchez/ocpp16messages/clearChargingProfile"
	ct "github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errId                     = "id"
	errConnectorId            = "connectorId"
	errChargingProfilePurpose = "chargingProfilePurpose"
	errStackLevel             = "stackLevel"

	purposeNotNil = "ChargingProfilePurpose should not be nil"

	valueZero       = 0
	valueOne        = 1
	valueTwo        = 2
	valueThree      = 3
	valueFive       = 5
	valueId         = 123
	valueNegative   = -1
	valueExceedsMax = 65536
)

func intPtr(v int) *int {
	return &v
}

func strPtr(v string) *string {
	return &v
}

func TestReq_Valid_NoFields(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Id != nil {
		t.Errorf("Id should be nil, got %v", req.Id)
	}

	if req.ConnectorId != nil {
		t.Errorf("ConnectorId should be nil, got %v", req.ConnectorId)
	}

	if req.ChargingProfilePurpose != nil {
		t.Errorf("ChargingProfilePurpose should be nil, got %v",
			req.ChargingProfilePurpose)
	}

	if req.StackLevel != nil {
		t.Errorf("StackLevel should be nil, got %v", req.StackLevel)
	}
}

func TestReq_Valid_WithId(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueId),
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Id == nil {
		t.Fatal("Id should not be nil")
	}

	if req.Id.Value() != valueId {
		t.Errorf(st.ErrorMismatchValue, valueId, req.Id.Value())
	}
}

func TestReq_Valid_WithIdZero(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueZero),
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Id == nil {
		t.Fatal("Id should not be nil")
	}

	if req.Id.Value() != valueZero {
		t.Errorf(st.ErrorMismatchValue, valueZero, req.Id.Value())
	}
}

func TestReq_Valid_WithConnectorId(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            intPtr(valueOne),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil {
		t.Fatal("ConnectorId should not be nil")
	}

	if req.ConnectorId.Value() != valueOne {
		t.Errorf(st.ErrorMismatchValue, valueOne, req.ConnectorId.Value())
	}
}

func TestReq_Valid_WithConnectorIdZero(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            intPtr(valueZero),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil {
		t.Fatal("ConnectorId should not be nil")
	}

	if req.ConnectorId.Value() != valueZero {
		t.Errorf(st.ErrorMismatchValue, valueZero, req.ConnectorId.Value())
	}
}

func TestReq_Valid_WithPurpose_ChargePointMaxProfile(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: strPtr("ChargePointMaxProfile"),
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	expected := ct.ChargePointMaxProfile
	if *req.ChargingProfilePurpose != expected {
		t.Errorf(st.ErrorMismatch, expected, *req.ChargingProfilePurpose)
	}
}

func TestReq_Valid_WithPurpose_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: strPtr("TxDefaultProfile"),
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	expected := ct.TxDefaultProfile
	if *req.ChargingProfilePurpose != expected {
		t.Errorf(st.ErrorMismatch, expected, *req.ChargingProfilePurpose)
	}
}

func TestReq_Valid_WithPurpose_TxProfile(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	expected := ct.TxProfile
	if *req.ChargingProfilePurpose != expected {
		t.Errorf(st.ErrorMismatch, expected, *req.ChargingProfilePurpose)
	}
}

func TestReq_Valid_WithStackLevel(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             intPtr(valueFive),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.StackLevel == nil {
		t.Fatal("StackLevel should not be nil")
	}

	if req.StackLevel.Value() != valueFive {
		t.Errorf(st.ErrorMismatchValue, valueFive, req.StackLevel.Value())
	}
}

func TestReq_Valid_AllFields_Id(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorId:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Id == nil || req.Id.Value() != valueOne {
		t.Error("Id mismatch")
	}
}

func TestReq_Valid_AllFields_ConnectorId(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorId:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil || req.ConnectorId.Value() != valueTwo {
		t.Error("ConnectorId mismatch")
	}
}

func TestReq_Valid_AllFields_Purpose(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorId:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	if *req.ChargingProfilePurpose != ct.TxProfile {
		t.Error("ChargingProfilePurpose mismatch")
	}
}

func TestReq_Valid_AllFields_StackLevel(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorId:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.StackLevel == nil || req.StackLevel.Value() != valueThree {
		t.Error("StackLevel mismatch")
	}
}

func TestReq_Invalid_NegativeId(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueNegative),
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative Id")
	}

	if !strings.Contains(err.Error(), errId) {
		t.Errorf(st.ErrorWantContains, err, errId)
	}
}

func TestReq_Invalid_IdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueExceedsMax),
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Id exceeds max")
	}

	if !strings.Contains(err.Error(), errId) {
		t.Errorf(st.ErrorWantContains, err, errId)
	}
}

func TestReq_Invalid_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            intPtr(valueNegative),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative ConnectorId")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_ConnectorIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            intPtr(valueExceedsMax),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "ConnectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_Purpose(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: strPtr("Invalid"),
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid ChargingProfilePurpose")
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(st.ErrorWantContains, err, errChargingProfilePurpose)
	}
}

func TestReq_Invalid_PurposeEmpty(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: strPtr(""),
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty ChargingProfilePurpose")
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(st.ErrorWantContains, err, errChargingProfilePurpose)
	}
}

func TestReq_Invalid_PurposeLowercase(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: strPtr("txprofile"),
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase ChargingProfilePurpose")
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(st.ErrorWantContains, err, errChargingProfilePurpose)
	}
}

func TestReq_Invalid_NegativeStackLevel(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             intPtr(valueNegative),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative StackLevel")
	}

	if !strings.Contains(err.Error(), errStackLevel) {
		t.Errorf(st.ErrorWantContains, err, errStackLevel)
	}
}

func TestReq_Invalid_StackLevelExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             intPtr(valueExceedsMax),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "StackLevel exceeds max")
	}

	if !strings.Contains(err.Error(), errStackLevel) {
		t.Errorf(st.ErrorWantContains, err, errStackLevel)
	}
}

func TestReq_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueNegative),
		ConnectorId:            intPtr(valueNegative),
		ChargingProfilePurpose: strPtr("Invalid"),
		StackLevel:             intPtr(valueNegative),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errId) {
		t.Errorf(st.ErrorWantContains, err, errId)
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(st.ErrorWantContains, err, errChargingProfilePurpose)
	}

	if !strings.Contains(err.Error(), errStackLevel) {
		t.Errorf(st.ErrorWantContains, err, errStackLevel)
	}
}
