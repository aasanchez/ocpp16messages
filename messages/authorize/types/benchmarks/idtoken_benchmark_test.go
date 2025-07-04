package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func BenchmarkIdToken_Create(b *testing.B) {
	raw := "B85A-50CBE9678EC6" // valid CiString20

	for range b.N {
		str, err := sharedtypes.SetCiString20(raw)
		if err != nil {
			b.Fatal(err)
		}

		_, err = authorizetypes.IdToken(str)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIdToken_Value(b *testing.B) {
	str, err := sharedtypes.SetCiString20("B85A-50CBE9678EC6")
	if err != nil {
		b.Fatal(err)
	}

	token, err := authorizetypes.IdToken(str)
	if err != nil {
		b.Fatal(err)
	}

	var val string

	for range b.N {
		val = token.Value()
	}

	_ = val
}
