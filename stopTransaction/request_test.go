package stoptransaction

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/metervalues/types"
)

func Test_validateTransactionData_EmptySlice(t *testing.T) {
	t.Parallel()

	const expectedLenZero = 0

	transactionData := []mt.MeterValueInput{}

	validated, errs := validateTransactionData(transactionData, nil)

	if errs != nil {
		t.Fatalf("errs = %v, want nil", errs)
	}

	if validated == nil {
		t.Fatal("validated data = nil, want empty slice")
	}

	if len(validated) != expectedLenZero {
		t.Fatalf(
			"len(validated) = %d, want %d",
			len(validated),
			expectedLenZero,
		)
	}
}

func Test_validateTransactionData_NilSlice(t *testing.T) {
	t.Parallel()

	validated, errs := validateTransactionData(nil, nil)

	if errs != nil {
		t.Fatalf("errs = %v, want nil", errs)
	}

	if validated != nil {
		t.Fatalf("validated = %v, want nil", validated)
	}
}
