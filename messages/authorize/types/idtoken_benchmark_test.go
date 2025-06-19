package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func BenchmarkIdToken_CreateValid(b *testing.B) {
	raw := "B85A-50CBE9678EC6" // valid CiString20
	for i := 0; i < b.N; i++ {
		ci, err := sharedtypes.CiString20(raw)
		if err != nil {
			b.Fatal(err)
		}
		_, err = authorizetypes.IdToken(ci)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIdToken_Value(b *testing.B) {
	ci, err := sharedtypes.CiString20("B85A-50CBE9678EC6")
	if err != nil {
		b.Fatal(err)
	}
	token, err := authorizetypes.IdToken(ci)
	if err != nil {
		b.Fatal(err)
	}

	var val string
	for i := 0; i < b.N; i++ {
		val = token.Value()
	}
	_ = val
}
