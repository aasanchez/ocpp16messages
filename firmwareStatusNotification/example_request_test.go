package firmwareStatusNotification_test

import (
	"fmt"

	fsn "github.com/aasanchez/ocpp16messages/firmwareStatusNotification"
)

const labelStatus = "Status:"

// ExampleReq demonstrates creating a valid FirmwareStatusNotification.req
// message with an Idle status.
func ExampleReq() {
	req, err := fsn.Req(fsn.ReqInput{Status: "Idle"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Idle
}

// ExampleReq_downloading demonstrates creating a FirmwareStatusNotification.req
// message with a Downloading status.
func ExampleReq_downloading() {
	req, err := fsn.Req(fsn.ReqInput{Status: "Downloading"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Downloading
}

// ExampleReq_downloaded demonstrates creating a FirmwareStatusNotification.req
// message with a Downloaded status.
func ExampleReq_downloaded() {
	req, err := fsn.Req(fsn.ReqInput{Status: "Downloaded"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Downloaded
}

// ExampleReq_installing demonstrates creating a FirmwareStatusNotification.req
// message with an Installing status.
func ExampleReq_installing() {
	req, err := fsn.Req(fsn.ReqInput{Status: "Installing"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Installing
}

// ExampleReq_installed demonstrates creating a FirmwareStatusNotification.req
// message with an Installed status.
func ExampleReq_installed() {
	req, err := fsn.Req(fsn.ReqInput{Status: "Installed"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Installed
}

// ExampleReq_downloadFailed demonstrates creating a
// FirmwareStatusNotification.req message with a DownloadFailed status.
func ExampleReq_downloadFailed() {
	req, err := fsn.Req(fsn.ReqInput{Status: "DownloadFailed"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: DownloadFailed
}

// ExampleReq_installationFailed demonstrates creating a
// FirmwareStatusNotification.req message with an InstallationFailed status.
func ExampleReq_installationFailed() {
	req, err := fsn.Req(fsn.ReqInput{Status: "InstallationFailed"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: InstallationFailed
}

// ExampleReq_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleReq_invalidStatus() {
	_, err := fsn.Req(fsn.ReqInput{Status: "InvalidStatus"})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// status: invalid value
}
