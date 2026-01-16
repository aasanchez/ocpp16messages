package getLocalListVersion_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getLocalListVersion"
	mt "github.com/aasanchez/ocpp16messages/getLocalListVersion/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testVersionPositive  = 5
	testVersionUnsupport = -1
	testVersionEmpty     = 0
)

func TestConf_Valid_PositiveVersion(t *testing.T) {
	t.Parallel()

	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: testVersionPositive,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != testVersionPositive {
		t.Errorf(
			st.ErrorMismatchValue,
			testVersionPositive,
			conf.ListVersion.Value(),
		)
	}
}

func TestConf_Valid_UnsupportedVersion(t *testing.T) {
	t.Parallel()

	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: testVersionUnsupport,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != mt.ListVersionUnsupported {
		t.Errorf(
			st.ErrorMismatchValue,
			mt.ListVersionUnsupported,
			conf.ListVersion.Value(),
		)
	}

	if !conf.ListVersion.IsUnsupported() {
		t.Error("IsUnsupported() = false, want true")
	}
}

func TestConf_Valid_EmptyListVersion(t *testing.T) {
	t.Parallel()

	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: testVersionEmpty,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != mt.ListVersionEmpty {
		t.Errorf(
			st.ErrorMismatchValue,
			mt.ListVersionEmpty,
			conf.ListVersion.Value(),
		)
	}

	if !conf.ListVersion.IsEmpty() {
		t.Error("IsEmpty() = false, want true")
	}
}

func TestConf_Valid_ZeroValue(t *testing.T) {
	t.Parallel()

	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: 0,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != 0 {
		t.Errorf(st.ErrorMismatchValue, 0, conf.ListVersion.Value())
	}
}
