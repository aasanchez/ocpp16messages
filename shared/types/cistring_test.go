package sharedtypes

import (
	"testing"
)

func TestInternal_SharedTypes_cistring_cistring(t *testing.T) {
	t.Parallel()

	_, _ = setCiString("OCPP 16", 20)
}
