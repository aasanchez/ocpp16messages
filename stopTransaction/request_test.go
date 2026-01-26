package stopTransaction

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
)

func Test_validateTransactionData_EmptySlice(t *testing.T) {
	t.Parallel()

	transactionData := []mt.MeterValueInput{}

	validated, errs := validateTransactionData(transactionData, nil)

	if errs != nil {
		t.Fatalf("errs = %v, want nil", errs)
	}

	if validated == nil {
		t.Fatal("validated data = nil, want empty slice")
	}

	if len(validated) != 0 {
		t.Fatalf("len(validated) = %d, want 0", len(validated))
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
