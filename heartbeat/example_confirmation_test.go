package heartbeat_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/heartbeat"
)

// ExampleConf demonstrates creating a valid Heartbeat.conf message
// with the Central System's current time.
func ExampleConf() {
	conf, err := heartbeat.Conf(heartbeat.ConfInput{
		CurrentTime: "2025-01-15T10:30:00Z",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("CurrentTime:", conf.CurrentTime.Value().UTC())
	// Output:
	// CurrentTime: 2025-01-15 10:30:00 +0000 UTC
}

// ExampleConf_emptyCurrentTime demonstrates the error returned when
// an empty current time is provided.
func ExampleConf_emptyCurrentTime() {
	_, err := heartbeat.Conf(heartbeat.ConfInput{CurrentTime: ""})
	if err != nil {
		fmt.Println("Error: empty currentTime")
	}
	// Output:
	// Error: empty currentTime
}

// ExampleConf_invalidCurrentTime demonstrates the error returned when
// an invalid date format is provided.
func ExampleConf_invalidCurrentTime() {
	_, err := heartbeat.Conf(heartbeat.ConfInput{CurrentTime: "not-a-date"})
	if err != nil {
		fmt.Println("Error: invalid currentTime format")
	}
	// Output:
	// Error: invalid currentTime format
}
