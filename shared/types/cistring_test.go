//go:build whitebox
// +build whitebox

package sharedtypes

import (
	"testing"
)

func TestSharedCiString_Home(t *testing.T) {
	t.Parallel()

	setCiString("OCPP 16", 20)
}
