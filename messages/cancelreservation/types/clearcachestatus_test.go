package cancelreservationtypes

import (
	"errors"
	"strings"
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestClearCacheStatus_Accepted(t *testing.T) {
	t.Parallel()

	result, err := ClearCacheStatus(ClearCacheAccepted)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Value() != ClearCacheAccepted {
		t.Errorf("expected value %q, got %q", ClearCacheAccepted, result.Value())
	}
}

func TestClearCacheStatus_Rejected(t *testing.T) {
	t.Parallel()

	result, err := ClearCacheStatus(ClearCacheRejected)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Value() != ClearCacheRejected {
		t.Errorf("expected value %q, got %q", ClearCacheRejected, result.Value())
	}
}

func TestClearCacheStatus_Invalid(t *testing.T) {
	t.Parallel()

	_, err := ClearCacheStatus("InvalidValue")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, sharedtypes.ErrInvalidClearCacheStatus) {
		t.Errorf("error is not ErrInvalidClearCacheStatus: %v", err)
	}

	if !strings.Contains(err.Error(), "InvalidValue") {
		t.Errorf("error does not mention input: %v", err)
	}
}
