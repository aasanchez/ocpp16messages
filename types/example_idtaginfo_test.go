package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

// Example output formatting constants.
const (
	// strError is the prefix used in example output for error messages.
	strError = "Error:"

	// strStatus is the prefix used in example output for status information.
	strStatus = "Status:"
)

// ExampleNewIdTagInfo demonstrates creating a basic IdTagInfo with an
// Accepted status.
func ExampleNewIdTagInfo() {
	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	fmt.Println(strStatus, idTagInfo.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleNewIdTagInfo_withExpiryDate demonstrates creating an IdTagInfo
// with an expiry date using the builder pattern.
func ExampleNewIdTagInfo_withExpiryDate() {
	expiryDate, _ := types.NewDateTime("2025-12-31T23:59:59Z")

	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	idTagInfo = idTagInfo.WithExpiryDate(expiryDate)
	fmt.Println(strStatus, idTagInfo.Status.String())
	fmt.Println("Has expiry date:", idTagInfo.ExpiryDate != nil)
	// Output:
	// Status: Accepted
	// Has expiry date: true
}

// ExampleNewIdTagInfo_withParentIdTag demonstrates creating an IdTagInfo
// with a parent ID tag using the builder pattern.
func ExampleNewIdTagInfo_withParentIdTag() {
	parentCiStr, _ := types.NewCiString20Type("PARENT-TAG-001")
	parentIdToken := types.NewIdToken(parentCiStr)

	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	idTagInfo = idTagInfo.WithParentIdTag(parentIdToken)
	fmt.Println(strStatus, idTagInfo.Status.String())
	fmt.Println("Has parent tag:", idTagInfo.ParentIdTag != nil)
	// Output:
	// Status: Accepted
	// Has parent tag: true
}

// ExampleNewIdTagInfo_complete demonstrates creating a complete IdTagInfo
// with all optional fields set.
func ExampleNewIdTagInfo_complete() {
	expiryDate, _ := types.NewDateTime("2025-12-31T23:59:59Z")
	parentCiStr, _ := types.NewCiString20Type("PARENT-TAG-001")
	parentIdToken := types.NewIdToken(parentCiStr)

	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	idTagInfo = idTagInfo.
		WithExpiryDate(expiryDate).
		WithParentIdTag(parentIdToken)

	fmt.Println(strStatus, idTagInfo.Status.String())
	fmt.Println("Has expiry date:", idTagInfo.ExpiryDate != nil)
	fmt.Println("Has parent tag:", idTagInfo.ParentIdTag != nil)
	// Output:
	// Status: Accepted
	// Has expiry date: true
	// Has parent tag: true
}

// ExampleNewIdTagInfo_blocked demonstrates creating an IdTagInfo with
// a Blocked status.
func ExampleNewIdTagInfo_blocked() {
	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusBlocked)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	fmt.Println(strStatus, idTagInfo.Status.String())
	// Output:
	// Status: Blocked
}

// ExampleNewIdTagInfo_expired demonstrates creating an IdTagInfo with
// an Expired status.
func ExampleNewIdTagInfo_expired() {
	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusExpired)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	fmt.Println(strStatus, idTagInfo.Status.String())
	// Output:
	// Status: Expired
}

// ExampleNewIdTagInfo_invalid demonstrates creating an IdTagInfo with
// an Invalid status.
func ExampleNewIdTagInfo_invalid() {
	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusInvalid)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	fmt.Println(strStatus, idTagInfo.Status.String())
	// Output:
	// Status: Invalid
}

// ExampleNewIdTagInfo_concurrentTx demonstrates creating an IdTagInfo with
// a ConcurrentTx status.
func ExampleNewIdTagInfo_concurrentTx() {
	idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatusConcurrentTx)
	if err != nil {
		fmt.Println(strError, err)

		return
	}

	fmt.Println(strStatus, idTagInfo.Status.String())
	// Output:
	// Status: ConcurrentTx
}
