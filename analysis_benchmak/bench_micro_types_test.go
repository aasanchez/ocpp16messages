//go:build bench

package benchmark

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

var (
	sinkDateTimeCustom    st.DateTime
	sinkDateTimePrimitive string

	sinkParentIdTagInfoCustom st.IdTagInfo
	sinkParentIdTagInfoPrim   primitiveIdTagInfo
	sinkParentIdTagValue      string
)

func BenchmarkDateTime_CustomType(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		dt, err := st.NewDateTime(benchmarkTimestamp)
		if err != nil {
			b.Fatal(err)
		}

		sinkDateTimeCustom = dt
	}
}

func BenchmarkDateTime_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sinkDateTimePrimitive = benchmarkTimestamp
	}
}

func BenchmarkDateTime_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := validatePrimitiveTimestampUTC(benchmarkTimestamp); err != nil {
			b.Fatal(err)
		}

		sinkDateTimePrimitive = benchmarkTimestamp
	}
}

func BenchmarkParentIdTag_CustomChain(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		idTagInfo, err := buildCustomParentIdTagInfo("PARENT-TAG-001")
		if err != nil {
			b.Fatal(err)
		}

		sinkParentIdTagInfoCustom = idTagInfo
	}
}

func BenchmarkParentIdTag_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		parentIdTag := "PARENT-TAG-001"
		idTagInfo := primitiveIdTagInfo{ParentIdTag: &parentIdTag}
		sinkParentIdTagInfoPrim = idTagInfo
	}
}

func BenchmarkParentIdTag_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		parentIdTag := "PARENT-TAG-001"
		if err := validatePrimitiveCiString20(parentIdTag); err != nil {
			b.Fatal(err)
		}

		idTagInfo := primitiveIdTagInfo{ParentIdTag: &parentIdTag}
		sinkParentIdTagInfoPrim = idTagInfo
	}
}

func BenchmarkParentIdTag_CustomReadPrimitive(b *testing.B) {
	b.ReportAllocs()

	idTagInfo, err := buildCustomParentIdTagInfo("PARENT-TAG-001")
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

	parentIdTag := "PARENT-TAG-001"
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
