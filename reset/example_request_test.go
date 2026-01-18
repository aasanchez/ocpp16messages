package reset_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/reset"
)

// ExampleReq demonstrates creating a valid Reset.req message
// with a Hard reset type.
func ExampleReq() {
	req, err := reset.Req(reset.ReqInput{
		Type: "Hard",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Type:", req.Type.String())
	// Output:
	// Type: Hard
}

// ExampleReq_soft demonstrates creating a Reset.req message
// with a Soft reset type.
func ExampleReq_soft() {
	req, err := reset.Req(reset.ReqInput{
		Type: "Soft",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Type:", req.Type.String())
	// Output:
	// Type: Soft
}

// ExampleReq_invalidType demonstrates the error returned when
// an invalid type is provided.
func ExampleReq_invalidType() {
	_, err := reset.Req(reset.ReqInput{
		Type: "Unknown",
	})
	if err != nil {
		fmt.Println("Error: invalid type")
	}
	// Output:
	// Error: invalid type
}

// ExampleReq_emptyType demonstrates the error returned when
// an empty type is provided.
func ExampleReq_emptyType() {
	_, err := reset.Req(reset.ReqInput{
		Type: "",
	})
	if err != nil {
		fmt.Println("Error: invalid type")
	}
	// Output:
	// Error: invalid type
}
