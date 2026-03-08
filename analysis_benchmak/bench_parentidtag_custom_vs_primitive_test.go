//go:build bench

package benchmark

import (
	"errors"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

const parentIdTagSample = "PARENT-TAG-001"

type primitiveIdTagInfo struct {
	ParentIdTag *string
}

var (
	errPrimitiveEmpty    = errors.New("parentIdTag cannot be empty")
	errPrimitiveTooLong  = errors.New("parentIdTag exceeds max length")
	errPrimitiveBadASCII = errors.New("parentIdTag contains invalid ASCII")

	sinkParentIdTagInfoCustom st.IdTagInfo
	sinkParentIdTagInfoPrim   primitiveIdTagInfo
	sinkParentIdTagValue      string
)

func BenchmarkParentIdTag_CustomChain(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		idTagInfo, err := buildCustomParentIdTagInfo(parentIdTagSample)
		if err != nil {
			b.Fatal(err)
		}

		sinkParentIdTagInfoCustom = idTagInfo
	}
}

func BenchmarkParentIdTag_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		parentIdTag := parentIdTagSample
		idTagInfo := primitiveIdTagInfo{ParentIdTag: &parentIdTag}
		sinkParentIdTagInfoPrim = idTagInfo
	}
}

func BenchmarkParentIdTag_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		parentIdTag := parentIdTagSample
		if err := validatePrimitiveParentIdTag(parentIdTag); err != nil {
			b.Fatal(err)
		}

		idTagInfo := primitiveIdTagInfo{ParentIdTag: &parentIdTag}
		sinkParentIdTagInfoPrim = idTagInfo
	}
}

func BenchmarkParentIdTag_CustomReadPrimitive(b *testing.B) {
	b.ReportAllocs()

	idTagInfo, err := buildCustomParentIdTagInfo(parentIdTagSample)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkParentIdTagValue = idTagInfo.ParentIdTag.Value().Value()
	}
}

func BenchmarkParentIdTag_PrimitiveRead(b *testing.B) {
	b.ReportAllocs()

	parentIdTag := parentIdTagSample
	idTagInfo := primitiveIdTagInfo{ParentIdTag: &parentIdTag}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkParentIdTagValue = *idTagInfo.ParentIdTag
	}
}

func buildCustomParentIdTagInfo(parentIdTag string) (st.IdTagInfo, error) {
	ciString, err := st.NewCiString20Type(parentIdTag)
	if err != nil {
		return st.IdTagInfo{}, err
	}

	idToken := st.NewIdToken(ciString)

	idTagInfo, err := st.NewIdTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		return st.IdTagInfo{}, err
	}

	return idTagInfo.WithParentIdTag(idToken), nil
}

func validatePrimitiveParentIdTag(parentIdTag string) error {
	if parentIdTag == "" {
		return errPrimitiveEmpty
	}

	if len(parentIdTag) > st.CiString20Max {
		return errPrimitiveTooLong
	}

	for index := 0; index < len(parentIdTag); index++ {
		char := parentIdTag[index]
		if char < 32 || char > 126 {
			return errPrimitiveBadASCII
		}
	}

	return nil
}
