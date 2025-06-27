package cancelreservationtypes_test

import (
	"testing"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func BenchmarkClearCacheStatus_Accepted(b *testing.B) {
	for range b.N {
		_, _ = cancelreservationtypes.ClearCacheStatus(cancelreservationtypes.ClearCacheAccepted)
	}
}

func BenchmarkClearCacheStatus_Rejected(b *testing.B) {
	for range b.N {
		_, _ = cancelreservationtypes.ClearCacheStatus(cancelreservationtypes.ClearCacheRejected)
	}
}

func BenchmarkClearCacheStatus_Invalid(b *testing.B) {
	for range b.N {
		_, _ = cancelreservationtypes.ClearCacheStatus("NotAStatus")
	}
}
