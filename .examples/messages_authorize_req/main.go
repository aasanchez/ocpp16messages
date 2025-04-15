package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	"github.com/aasanchez/ocpp16messages/types"
)

// FullMessage represents the full structure of a typical OCPP message
type FullMessage struct {
	MessageTypeId int                      `json:"MessageTypeId"`
	UniqueId      string                   `json:"UniqueId"`
	Action        string                   `json:"Action"`
	Payload       authorize.RequestMessage `json:"Payload"`
}
