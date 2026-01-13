package bootNotification_test

import (
	"errors"
	"strings"
	"testing"

	bn "github.com/aasanchez/ocpp16messages/messages/bootNotification"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	statusAccepted    = "Accepted"
	statusPending     = "Pending"
	statusRejected    = "Rejected"
	validTime         = "2025-01-15T12:00:00Z"
	validInterval     = 300
	errStatus         = "status"
	errCurrentTime    = "currentTime"
	errInterval       = "interval"
	zeroInterval      = 0
	utcNormalizedHour = 7
)

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	conf, err := bn.Conf(bn.ConfInput{
		Status:      statusAccepted,
		CurrentTime: validTime,
		Interval:    validInterval,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusAccepted {
		t.Errorf(st.ErrorMismatch, statusAccepted, conf.Status.String())
	}
}

func TestConf_ValidPending(t *testing.T) {
	t.Parallel()

	conf, err := bn.Conf(bn.ConfInput{
		Status:      statusPending,
		CurrentTime: validTime,
		Interval:    validInterval,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusPending {
		t.Errorf(st.ErrorMismatch, statusPending, conf.Status.String())
	}
}

func TestConf_ValidRejected(t *testing.T) {
	t.Parallel()

	conf, err := bn.Conf(bn.ConfInput{
		Status:      statusRejected,
		CurrentTime: validTime,
		Interval:    validInterval,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusRejected {
		t.Errorf(st.ErrorMismatch, statusRejected, conf.Status.String())
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := bn.Conf(bn.ConfInput{
		Status:      "Unknown",
		CurrentTime: validTime,
		Interval:    validInterval,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := bn.Conf(bn.ConfInput{
		Status:      "",
		CurrentTime: validTime,
		Interval:    validInterval,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_InvalidCurrentTime(t *testing.T) {
	t.Parallel()

	_, err := bn.Conf(bn.ConfInput{
		Status:      statusAccepted,
		CurrentTime: "not-a-date",
		Interval:    validInterval,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid currentTime")
	}

	if !strings.Contains(err.Error(), errCurrentTime) {
		t.Errorf(st.ErrorWantContains, err, errCurrentTime)
	}
}

func TestConf_EmptyCurrentTime(t *testing.T) {
	t.Parallel()

	_, err := bn.Conf(bn.ConfInput{
		Status:      statusAccepted,
		CurrentTime: "",
		Interval:    validInterval,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty currentTime")
	}

	if !strings.Contains(err.Error(), errCurrentTime) {
		t.Errorf(st.ErrorWantContains, err, errCurrentTime)
	}
}

func TestConf_NegativeInterval(t *testing.T) {
	t.Parallel()

	_, err := bn.Conf(bn.ConfInput{
		Status:      statusAccepted,
		CurrentTime: validTime,
		Interval:    -1,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative interval")
	}

	if !strings.Contains(err.Error(), errInterval) {
		t.Errorf(st.ErrorWantContains, err, errInterval)
	}
}

func TestConf_IntervalTooLarge(t *testing.T) {
	t.Parallel()

	_, err := bn.Conf(bn.ConfInput{
		Status:      statusAccepted,
		CurrentTime: validTime,
		Interval:    70000, // exceeds uint16 max (65535)
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "interval too large")
	}

	if !strings.Contains(err.Error(), errInterval) {
		t.Errorf(st.ErrorWantContains, err, errInterval)
	}
}

func TestConf_ZeroInterval(t *testing.T) {
	t.Parallel()

	conf, err := bn.Conf(bn.ConfInput{
		Status:      statusAccepted,
		CurrentTime: validTime,
		Interval:    0,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Interval.Value() != zeroInterval {
		t.Errorf(
			"Conf() Interval = %d, want %d",
			conf.Interval.Value(),
			zeroInterval,
		)
	}
}

func TestConf_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := bn.Conf(bn.ConfInput{
		Status:      "Invalid-Status",
		CurrentTime: "not-a-date",
		Interval:    -1,
	})
	if err == nil {
		t.Error("Conf() error = nil, want multiple errors")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}

	if !strings.Contains(errStr, errCurrentTime) {
		t.Errorf(st.ErrorWantContains, err, errCurrentTime)
	}

	if !strings.Contains(errStr, errInterval) {
		t.Errorf(st.ErrorWantContains, err, errInterval)
	}
}

func TestConf_CurrentTimeNormalized(t *testing.T) {
	t.Parallel()

	// Test that time is normalized to UTC
	conf, err := bn.Conf(bn.ConfInput{
		Status:      statusAccepted,
		CurrentTime: "2025-01-15T12:00:00+05:00",
		Interval:    validInterval,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	// Time should be normalized to UTC (07:00:00Z)
	got := conf.CurrentTime.Value().Hour()

	if got != utcNormalizedHour {
		t.Errorf(
			"Conf() CurrentTime hour = %d, want %d (UTC)",
			got,
			utcNormalizedHour,
		)
	}
}
