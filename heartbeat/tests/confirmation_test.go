package heartbeat_test

import (
	"strings"
	"testing"
	"time"

	"github.com/aasanchez/ocpp16messages/heartbeat"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidCurrentTime = "2025-01-15T10:30:00Z"
	errFieldCurrentTime  = "currentTime"
	testNanoseconds      = 123456789
)

func TestConf_Valid(t *testing.T) {
	t.Parallel()

	conf, err := heartbeat.Conf(heartbeat.ConfInput{
		CurrentTime: testValidCurrentTime,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expectedTime, _ := time.Parse(time.RFC3339, testValidCurrentTime)
	if !conf.CurrentTime.Value().Equal(expectedTime) {
		t.Errorf(
			st.ErrorMismatch,
			expectedTime,
			conf.CurrentTime.Value(),
		)
	}
}

func TestConf_EmptyCurrentTime(t *testing.T) {
	t.Parallel()

	_, err := heartbeat.Conf(heartbeat.ConfInput{CurrentTime: ""})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty currentTime")
	}

	if !strings.Contains(err.Error(), errFieldCurrentTime) {
		t.Errorf(st.ErrorWantContains, err, errFieldCurrentTime)
	}
}

func TestConf_InvalidCurrentTime(t *testing.T) {
	t.Parallel()

	_, err := heartbeat.Conf(heartbeat.ConfInput{CurrentTime: "not-a-date"})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid currentTime")
	}

	if !strings.Contains(err.Error(), errFieldCurrentTime) {
		t.Errorf(st.ErrorWantContains, err, errFieldCurrentTime)
	}
}

func TestConf_CurrentTimeNormalizesToUTC(t *testing.T) {
	t.Parallel()

	// Input with timezone offset
	conf, err := heartbeat.Conf(heartbeat.ConfInput{
		CurrentTime: "2025-01-15T12:30:00+02:00",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	// Should be normalized to UTC (10:30 UTC = 12:30 +02:00)
	expectedTime, _ := time.Parse(time.RFC3339, "2025-01-15T10:30:00Z")
	if !conf.CurrentTime.Value().Equal(expectedTime) {
		t.Errorf(
			st.ErrorMismatch,
			expectedTime,
			conf.CurrentTime.Value(),
		)
	}
}

func TestConf_CurrentTimeWithNanoseconds(t *testing.T) {
	t.Parallel()

	conf, err := heartbeat.Conf(heartbeat.ConfInput{
		CurrentTime: "2025-01-15T10:30:00.123456789Z",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	// Verify nanoseconds are preserved
	if conf.CurrentTime.Value().Nanosecond() != testNanoseconds {
		t.Errorf(
			"Conf() nanoseconds = %d, want %d",
			conf.CurrentTime.Value().Nanosecond(),
			testNanoseconds,
		)
	}
}
