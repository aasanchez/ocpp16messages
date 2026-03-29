package getlocallistversion_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getlocallistversion"
	types "github.com/aasanchez/ocpp16types"
)

const (
	testVersionPositive  = 5
	testVersionUnsupport = -1
	testVersionEmpty     = 0
	testVersionOverflow  = 2147483648
)

func TestConf_Valid_PositiveVersion(t *testing.T) {
	t.Parallel()

	conf, err := getlocallistversion.Conf(getlocallistversion.ConfInput{
		ListVersion: testVersionPositive,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != testVersionPositive {
		t.Errorf(
			types.ErrorMismatchValue,
			testVersionPositive,
			conf.ListVersion.Value(),
		)
	}
}

func TestConf_Valid_UnsupportedVersion(t *testing.T) {
	t.Parallel()

	conf, err := getlocallistversion.Conf(getlocallistversion.ConfInput{
		ListVersion: testVersionUnsupport,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != types.ListVersionUnsupported {
		t.Errorf(
			types.ErrorMismatchValue,
			types.ListVersionUnsupported,
			conf.ListVersion.Value(),
		)
	}

	if !conf.ListVersion.IsUnsupported() {
		t.Error("IsUnsupported() = false, want true")
	}
}

func TestConf_Valid_EmptyListVersion(t *testing.T) {
	t.Parallel()

	conf, err := getlocallistversion.Conf(getlocallistversion.ConfInput{
		ListVersion: testVersionEmpty,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != types.ListVersionEmpty {
		t.Errorf(
			types.ErrorMismatchValue,
			types.ListVersionEmpty,
			conf.ListVersion.Value(),
		)
	}

	if !conf.ListVersion.IsEmpty() {
		t.Error("IsEmpty() = false, want true")
	}
}

func TestConf_Valid_ZeroValue(t *testing.T) {
	t.Parallel()

	conf, err := getlocallistversion.Conf(getlocallistversion.ConfInput{
		ListVersion: testVersionEmpty,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ListVersion.Value() != testVersionEmpty {
		t.Errorf(
			types.ErrorMismatchValue,
			testVersionEmpty,
			conf.ListVersion.Value(),
		)
	}
}

func TestConf_InvalidListVersion_ExceedsInt32(t *testing.T) {
	t.Parallel()

	// Value exceeds int32 max (2147483647), should cause overflow error
	_, err := getlocallistversion.Conf(getlocallistversion.ConfInput{
		ListVersion: testVersionOverflow,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for int32 overflow")
	}
}
