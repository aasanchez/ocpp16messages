package core_test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestMessageTypeConstants(t *testing.T) {
	if core.CALL != 2 {
		t.Errorf("Expected CALL to be 2, got %d", core.CALL)
	}
	if core.CALLRESULT != 3 {
		t.Errorf("Expected CALLRESULT to be 3, got %d", core.CALLRESULT)
	}
	if core.CALLERROR != 4 {
		t.Errorf("Expected CALLERROR to be 4, got %d", core.CALLERROR)
	}
}
