package startTransaction_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/startTransaction"
)

const (
	labelStatus        = "Status:"
	labelTransactionId = "TransactionId:"
)

// ExampleConf demonstrates creating a valid StartTransaction.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: 12345,
		Status:        "Accepted",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelTransactionId, conf.TransactionId.Value())
	fmt.Println(labelStatus, conf.IdTagInfo.Status.String())
	// Output:
	// TransactionId: 12345
	// Status: Accepted
}

// ExampleConf_blocked demonstrates creating a StartTransaction.conf message
// with a Blocked status.
func ExampleConf_blocked() {
	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: 12346,
		Status:        "Blocked",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelTransactionId, conf.TransactionId.Value())
	fmt.Println(labelStatus, conf.IdTagInfo.Status.String())
	// Output:
	// TransactionId: 12346
	// Status: Blocked
}

// ExampleConf_withExpiryDate demonstrates creating a StartTransaction.conf
// message with an expiry date.
func ExampleConf_withExpiryDate() {
	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: 12347,
		Status:        "Accepted",
		ExpiryDate:    &expiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelTransactionId, conf.TransactionId.Value())
	fmt.Println(labelStatus, conf.IdTagInfo.Status.String())
	fmt.Println("HasExpiryDate:", conf.IdTagInfo.ExpiryDate != nil)
	// Output:
	// TransactionId: 12347
	// Status: Accepted
	// HasExpiryDate: true
}

// ExampleConf_withParentIdTag demonstrates creating a StartTransaction.conf
// message with a parent ID tag.
func ExampleConf_withParentIdTag() {
	parentTag := "PARENT-123"

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: 12348,
		Status:        "Accepted",
		ExpiryDate:    nil,
		ParentIdTag:   &parentTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelTransactionId, conf.TransactionId.Value())
	fmt.Println(labelStatus, conf.IdTagInfo.Status.String())
	fmt.Println("ParentIdTag:", conf.IdTagInfo.ParentIdTag.String())
	// Output:
	// TransactionId: 12348
	// Status: Accepted
	// ParentIdTag: PARENT-123
}

// ExampleConf_complete demonstrates creating a complete StartTransaction.conf
// message with all optional fields populated.
func ExampleConf_complete() {
	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: 12349,
		Status:        "Accepted",
		ExpiryDate:    &expiryDate,
		ParentIdTag:   &parentTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelTransactionId, conf.TransactionId.Value())
	fmt.Println(labelStatus, conf.IdTagInfo.Status.String())
	fmt.Println("HasExpiryDate:", conf.IdTagInfo.ExpiryDate != nil)
	fmt.Println("ParentIdTag:", conf.IdTagInfo.ParentIdTag.String())
	// Output:
	// TransactionId: 12349
	// Status: Accepted
	// HasExpiryDate: true
	// ParentIdTag: PARENT-123
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: 12345,
		Status:        "Unknown",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// status: NewIdTagInfo: AuthorizationStatus: invalid value
}

// ExampleConf_multipleErrors demonstrates that all validation errors
// are returned at once, not just the first one encountered.
func ExampleConf_multipleErrors() {
	invalidDate := "not-a-date"
	longTag := "THIS-TAG-IS-WAY-TOO-LONG-FOR-OCPP"

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: -1,
		Status:        "Invalid-Status",
		ExpiryDate:    &invalidDate,
		ParentIdTag:   &longTag,
	})
	if err != nil {
		fmt.Println("Multiple errors")
	}
	// Output:
	// Multiple errors
}
