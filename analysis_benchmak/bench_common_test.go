//go:build bench

package benchmark

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

const (
	benchmarkTimestamp = "2025-01-02T15:00:00Z"

	scaleOne         = 1
	scaleTwentyFive  = 25
	scaleOneHundred  = 100
	scaleTwoHundred  = 250
	scaleFiveHundred = 500
	scaleOneThousand = 1000

	maxUint16 = 65535
)

var (
	errPrimitiveEmpty        = errors.New("parentIdTag cannot be empty")
	errPrimitiveTooLong      = errors.New("parentIdTag exceeds max length")
	errPrimitiveBadASCII     = errors.New("parentIdTag contains invalid ASCII")
	errPrimitiveInvalidRange = errors.New("value out of uint16 range")
)

type primitiveIdTagInfo struct {
	ParentIdTag *string
}

type primitiveStartTransactionReq struct {
	ConnectorId   int
	IdTag         string
	MeterStart    int
	Timestamp     string
	ReservationId *int
}

type primitiveAuthorizationDataInput struct {
	IdTag string
}

type primitiveSendLocalListReq struct {
	ListVersion            int
	LocalAuthorizationList []primitiveAuthorizationDataInput
	UpdateType             string
}

type primitiveGetConfigurationReq struct {
	Key []string
}

func makeAuthorizationInputs(size int) []types.AuthorizationDataInput {
	entries := make([]types.AuthorizationDataInput, 0, size)

	for index := 0; index < size; index++ {
		entries = append(entries, types.AuthorizationDataInput{
			IdTag: "TAG-" + strconv.Itoa(index),
		})
	}

	return entries
}

func makePrimitiveAuthorizationInputs(size int) []primitiveAuthorizationDataInput {
	entries := make([]primitiveAuthorizationDataInput, 0, size)

	for index := 0; index < size; index++ {
		entries = append(entries, primitiveAuthorizationDataInput{
			IdTag: "TAG-" + strconv.Itoa(index),
		})
	}

	return entries
}

func makeConfigurationKeys(size int) []string {
	keys := make([]string, 0, size)

	for index := 0; index < size; index++ {
		keys = append(keys, "HeartbeatInterval"+strconv.Itoa(index%10))
	}

	return keys
}

func validatePrimitiveIntegerRange(value int) error {
	if value < 0 || value > maxUint16 {
		return errPrimitiveInvalidRange
	}

	return nil
}

func validatePrimitiveCiString20(value string) error {
	if value == "" {
		return errPrimitiveEmpty
	}

	if len(value) > types.CiString20Max {
		return errPrimitiveTooLong
	}

	for index := 0; index < len(value); index++ {
		char := value[index]
		if char < 32 || char > 126 {
			return errPrimitiveBadASCII
		}
	}

	return nil
}

func validatePrimitiveCiString50(value string) error {
	if value == "" {
		return errPrimitiveEmpty
	}

	if len(value) > types.CiString50Max {
		return errPrimitiveTooLong
	}

	for index := 0; index < len(value); index++ {
		char := value[index]
		if char < 32 || char > 126 {
			return errPrimitiveBadASCII
		}
	}

	return nil
}

func validatePrimitiveTimestampUTC(value string) error {
	if value == "" {
		return errPrimitiveEmpty
	}

	timestamp, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrInvalidValue, err)
	}

	_, offset := timestamp.Zone()
	if offset != 0 {
		return fmt.Errorf(
			"%w: expected UTC offset, got %s",
			types.ErrInvalidValue,
			timestamp.Format("Z07:00"),
		)
	}

	return nil
}

func validatePrimitiveStartTransactionReq(input primitiveStartTransactionReq) error {
	if err := validatePrimitiveIntegerRange(input.ConnectorId); err != nil {
		return fmt.Errorf("connectorId: %w", err)
	}

	if err := validatePrimitiveCiString20(input.IdTag); err != nil {
		return fmt.Errorf("idTag: %w", err)
	}

	if err := validatePrimitiveIntegerRange(input.MeterStart); err != nil {
		return fmt.Errorf("meterStart: %w", err)
	}

	if err := validatePrimitiveTimestampUTC(input.Timestamp); err != nil {
		return fmt.Errorf("timestamp: %w", err)
	}

	if input.ReservationId != nil {
		if err := validatePrimitiveIntegerRange(*input.ReservationId); err != nil {
			return fmt.Errorf("reservationId: %w", err)
		}
	}

	return nil
}

func validatePrimitiveSendLocalListReq(input primitiveSendLocalListReq) error {
	if err := validatePrimitiveIntegerRange(input.ListVersion); err != nil {
		return fmt.Errorf("listVersion: %w", err)
	}

	if input.UpdateType != types.UpdateTypeFull.String() &&
		input.UpdateType != types.UpdateTypeDifferential.String() {
		return fmt.Errorf("updateType: %w", types.ErrInvalidValue)
	}

	if input.LocalAuthorizationList == nil {
		return nil
	}

	for index, entry := range input.LocalAuthorizationList {
		if err := validatePrimitiveCiString20(entry.IdTag); err != nil {
			return fmt.Errorf("localAuthorizationList[%d]: %w", index, err)
		}
	}

	return nil
}

func validatePrimitiveGetConfigurationReq(input primitiveGetConfigurationReq) error {
	if len(input.Key) == 0 {
		return nil
	}

	for index, key := range input.Key {
		if err := validatePrimitiveCiString50(key); err != nil {
			return fmt.Errorf("key[%d]: %w", index, err)
		}
	}

	return nil
}
