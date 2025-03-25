package types

import (
	"testing"
)

func TestMessageTypeConstants(t *testing.T) {
	if CALL != 2 {
		t.Errorf("Expected CALL to be 2, got %d", CALL)
	}
	if CALLRESULT != 3 {
		t.Errorf("Expected CALLRESULT to be 3, got %d", CALLRESULT)
	}
	if CALLERROR != 4 {
		t.Errorf("Expected CALLERROR to be 4, got %d", CALLERROR)
	}
}
