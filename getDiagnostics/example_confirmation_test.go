package getDiagnostics_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getDiagnostics"
)

const (
	exampleFileNameValue = "diagnostics_20250101.zip"

	longStrPart = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
)

// ExampleConf demonstrates creating a GetDiagnostics.conf message
// with no file name (no diagnostics available).
func ExampleConf() {
	conf, err := getDiagnostics.Conf(getDiagnostics.ConfInput{
		FileName: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	if conf.FileName == nil {
		fmt.Println("FileName: nil (no diagnostics available)")
	}
	// Output:
	// FileName: nil (no diagnostics available)
}

// ExampleConf_withFileName demonstrates creating a GetDiagnostics.conf message
// with a file name indicating diagnostics are available.
func ExampleConf_withFileName() {
	fileName := exampleFileNameValue

	conf, err := getDiagnostics.Conf(getDiagnostics.ConfInput{
		FileName: &fileName,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("FileName:", conf.FileName.Value())
	// Output:
	// FileName: diagnostics_20250101.zip
}

// ExampleConf_invalidFileName demonstrates the error returned when
// a file name exceeding 255 characters is provided.
func ExampleConf_invalidFileName() {
	longFileName := "diagnostics_" +
		longStrPart +
		longStrPart +
		longStrPart +
		longStrPart +
		".zip"

	_, err := getDiagnostics.Conf(getDiagnostics.ConfInput{
		FileName: &longFileName,
	})
	if err != nil {
		fmt.Println("Error: file name too long")
	}
	// Output:
	// Error: file name too long
}
