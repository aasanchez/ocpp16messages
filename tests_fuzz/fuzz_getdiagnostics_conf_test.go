//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"unicode/utf8"

	"github.com/aasanchez/ocpp16messages/getdiagnostics"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzGetDiagnosticsConf(f *testing.F) {
	f.Add(false, "")
	f.Add(true, "diag.log")
	f.Add(true, "")
	f.Add(true, "\x00")
	f.Add(true, "a")

	f.Fuzz(func(t *testing.T, hasFileName bool, fileName string) {
		if len(fileName) > maxFuzzStringLen {
			t.Skip()
		}

		if !utf8.ValidString(fileName) && len(fileName) > 256 {
			t.Skip()
		}

		var fileNamePtr *string
		if hasFileName {
			fileNamePtr = &fileName
		}

		conf, err := getdiagnostics.Conf(getdiagnostics.ConfInput{
			FileName: fileNamePtr,
		})
		if err != nil {
			if !errors.Is(err, st.ErrEmptyValue) && !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if hasFileName {
			if conf.FileName == nil {
				t.Fatal("FileName = nil, want non-nil")
			}
			if fileName == "" {
				t.Fatal("Conf succeeded with empty FileName")
			}
			if conf.FileName.String() != fileName {
				t.Fatalf("FileName = %q, want %q", conf.FileName.String(), fileName)
			}
		} else if conf.FileName != nil {
			t.Fatal("FileName != nil, want nil")
		}
	})
}
