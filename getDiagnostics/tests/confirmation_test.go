package getDiagnostics_test

import (
	"strings"
	"testing"

	gd "github.com/aasanchez/ocpp16messages/getDiagnostics"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errFileName = "fileName"

	validFileNameValue = "diagnostics_20250101.zip"
	shortFileNameValue = "a.zip"

	longStrPart = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	fileNameNotNil = "FileName should not be nil"
)

func TestConf_Valid_NoFileName(t *testing.T) {
	t.Parallel()

	conf, err := gd.Conf(gd.ConfInput{
		FileName: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.FileName != nil {
		t.Errorf("FileName should be nil, got %v", conf.FileName)
	}
}

func TestConf_Valid_NilFileName(t *testing.T) {
	t.Parallel()

	conf, err := gd.Conf(gd.ConfInput{
		FileName: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.FileName != nil {
		t.Errorf("FileName should be nil, got %v", conf.FileName)
	}
}

func TestConf_Valid_WithFileName(t *testing.T) {
	t.Parallel()

	fileName := validFileNameValue

	conf, err := gd.Conf(gd.ConfInput{
		FileName: &fileName,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.FileName == nil {
		t.Fatal(fileNameNotNil)
	}

	if conf.FileName.Value() != validFileNameValue {
		t.Errorf(st.ErrorMismatch, validFileNameValue, conf.FileName.Value())
	}
}

func TestConf_Valid_WithShortFileName(t *testing.T) {
	t.Parallel()

	fileName := shortFileNameValue

	conf, err := gd.Conf(gd.ConfInput{
		FileName: &fileName,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.FileName == nil {
		t.Fatal(fileNameNotNil)
	}

	if conf.FileName.Value() != shortFileNameValue {
		t.Errorf(st.ErrorMismatch, shortFileNameValue, conf.FileName.Value())
	}
}

func TestConf_Invalid_EmptyFileName(t *testing.T) {
	t.Parallel()

	fileName := ""

	_, err := gd.Conf(gd.ConfInput{
		FileName: &fileName,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty FileName")
	}

	if !strings.Contains(err.Error(), errFileName) {
		t.Errorf(st.ErrorWantContains, err, errFileName)
	}
}

func TestConf_Invalid_FileNameExceedsMax(t *testing.T) {
	t.Parallel()

	// Create a string that exceeds 255 characters
	longFileName := "diagnostics_" +
		longStrPart +
		longStrPart +
		longStrPart +
		longStrPart +
		".zip"

	_, err := gd.Conf(gd.ConfInput{
		FileName: &longFileName,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "FileName exceeds max length")
	}

	if !strings.Contains(err.Error(), errFileName) {
		t.Errorf(st.ErrorWantContains, err, errFileName)
	}
}
