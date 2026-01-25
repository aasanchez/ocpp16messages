package types_test

import (
	"fmt"

	ut "github.com/aasanchez/ocpp16messages/unlockConnector/types"
)

const (
	outStatusLabel = "Status:"
	outValidLabel  = "Valid:"
)

// ExampleUnlockStatusUnlocked demonstrates using the Unlocked status constant.
func ExampleUnlockStatusUnlocked() {
	status := ut.UnlockStatusUnlocked
	fmt.Println(outStatusLabel, status.String())
	fmt.Println(outValidLabel, status.IsValid())
	// Output:
	// Status: Unlocked
	// Valid: true
}

// ExampleUnlockStatusUnlockFailed demonstrates using the UnlockFailed status
// constant.
func ExampleUnlockStatusUnlockFailed() {
	status := ut.UnlockStatusUnlockFailed
	fmt.Println(outStatusLabel, status.String())
	fmt.Println(outValidLabel, status.IsValid())
	// Output:
	// Status: UnlockFailed
	// Valid: true
}

// ExampleUnlockStatusNotSupported demonstrates using the NotSupported status
// constant.
func ExampleUnlockStatusNotSupported() {
	status := ut.UnlockStatusNotSupported
	fmt.Println(outStatusLabel, status.String())
	fmt.Println(outValidLabel, status.IsValid())
	// Output:
	// Status: NotSupported
	// Valid: true
}
