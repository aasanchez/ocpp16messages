package sendlocallist

import (
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func Test_validateReqAuthorizationList_EmptySlice(t *testing.T) {
	t.Parallel()

	const expectedLenZero = 0

	authList := []types.AuthorizationDataInput{}

	validated, errs := validateReqAuthorizationList(authList, nil)

	if errs != nil {
		t.Fatalf("errs = %v, want nil", errs)
	}

	if validated == nil {
		t.Fatal("validated list = nil, want empty slice")
	}

	if len(validated) != expectedLenZero {
		t.Fatalf(
			"len(validated) = %d, want %d",
			len(validated),
			expectedLenZero,
		)
	}
}
