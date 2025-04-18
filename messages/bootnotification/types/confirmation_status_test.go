package bootnotificationtypes

// import "testing"

// const (
// 	errExpectedOutput       = "expected string output, got %s"
// 	errExpectedValidate     = "expected Validate to return nil, got %v"
// 	errExpectedInvalid      = "expected IsValid() to return false for %v, but got true" // For invalid status checks
// 	errExpectedValid        = "expected IsValid() to return true for %v, but got false" // For valid status checks
// 	errExpectedInvalidNoArg = "expected IsValid() to return false, but got true"        // For cases with no specific arg
// )

// func TestBootNotificationStatus_AcceptedIsValid(t *testing.T) {
// 	t.Parallel()

// 	status := Accepted
// 	if !status.IsValid() {
// 		t.Errorf(errExpectedValid, status)
// 	}
// }

// func TestBootNotificationStatus_PendingIsValid(t *testing.T) {
// 	t.Parallel()

// 	status := Pending
// 	if !status.IsValid() {
// 		t.Errorf(errExpectedValid, status)
// 	}
// }

// func TestBootNotificationStatus_RejectedIsValid(t *testing.T) {
// 	t.Parallel()

// 	status := Rejected
// 	if !status.IsValid() {
// 		t.Errorf(errExpectedValid, status)
// 	}
// }

// // Test for invalid BootNotificationStatus values.
// func TestBootNotificationStatus_InvalidStatus(t *testing.T) {
// 	t.Parallel()

// 	invalidStatus := BootNotificationStatus("UnknownStatus")
// 	if invalidStatus.IsValid() {
// 		t.Errorf(errExpectedInvalid, invalidStatus)
// 	}
// }

// func TestBootNotificationStatus_EmptyStatus(t *testing.T) {
// 	t.Parallel()

// 	emptyStatus := BootNotificationStatus("")
// 	if emptyStatus.IsValid() {
// 		t.Error(errExpectedInvalidNoArg)
// 	}
// }

// func TestBootNotificationStatus_NullEquivalent(t *testing.T) {
// 	t.Parallel()

// 	nullStatus := BootNotificationStatus("null")
// 	if nullStatus.IsValid() {
// 		t.Errorf(errExpectedInvalid, nullStatus)
// 	}
// }

// func TestBootNotificationStatusAcceptedConstant(t *testing.T) {
// 	t.Parallel()

// 	if string(Accepted) != "Accepted" {
// 		t.Errorf(errExpectedOutput, Accepted)
// 	}
// }

// func TestBootNotificationStatusPendingConstant(t *testing.T) {
// 	t.Parallel()

// 	if string(Pending) != "Pending" {
// 		t.Errorf(errExpectedOutput, Pending)
// 	}
// }

// func TestBootNotificationStatusRejectedConstant(t *testing.T) {
// 	t.Parallel()

// 	if string(Rejected) != "Rejected" {
// 		t.Errorf(errExpectedOutput, Rejected)
// 	}
// }
