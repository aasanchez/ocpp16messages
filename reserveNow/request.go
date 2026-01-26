package reserveNow

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a ReserveNow.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: Unique identifier for this reservation.
	ReservationId int
	// Required: The connector to reserve. 0 means reserve any available.
	ConnectorId int
	// Required: The identifier to use for the reservation.
	IdTag string
	// Required: The date and time when the reservation expires (RFC3339).
	ExpiryDate string
	// Optional: The parent identifier to use for authorization.
	ParentIdTag *string
}

// ReqMessage represents an OCPP 1.6 ReserveNow.req message.
type ReqMessage struct {
	ReservationId st.Integer
	ConnectorId   st.Integer
	IdTag         st.CiString20Type
	ExpiryDate    st.DateTime
	ParentIdTag   *st.CiString20Type
}

// Req creates a ReserveNow.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ReservationId is negative or exceeds uint16 max value (65535)
//   - ConnectorId is negative or exceeds uint16 max value (65535)
//   - IdTag is empty
//   - IdTag exceeds 20 characters
//   - IdTag contains non-printable ASCII characters
//   - ExpiryDate is not a valid RFC3339 timestamp
//   - ParentIdTag (if provided) exceeds 20 characters
//   - ParentIdTag (if provided) contains non-printable ASCII characters
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	reservationId, err := st.NewIntegerFromInt(input.ReservationId)
	if err != nil {
		errs = append(errs, fmt.Errorf("reservationId: %w", err))
	}

	connectorId, err := st.NewIntegerFromInt(input.ConnectorId)
	if err != nil {
		errs = append(errs, fmt.Errorf("connectorId: %w", err))
	}

	idTag, err := st.NewCiString20Type(input.IdTag)
	if err != nil {
		errs = append(errs, fmt.Errorf("idTag: %w", err))
	}

	expiryDate, err := st.NewDateTime(input.ExpiryDate)
	if err != nil {
		errs = append(errs, fmt.Errorf("expiryDate: %w", err))
	}

	var parentIdTag *st.CiString20Type

	if input.ParentIdTag != nil {
		parentIdTag, errs = validateParentIdTag(*input.ParentIdTag, errs)
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		ReservationId: reservationId,
		ConnectorId:   connectorId,
		IdTag:         idTag,
		ExpiryDate:    expiryDate,
		ParentIdTag:   parentIdTag,
	}, nil
}

// validateParentIdTag validates the optional parentIdTag field.
func validateParentIdTag(
	parentIdTag string,
	errs []error,
) (*st.CiString20Type, []error) {
	val, err := st.NewCiString20Type(parentIdTag)
	if err != nil {
		return nil, append(errs, fmt.Errorf("parentIdTag: %w", err))
	}

	return &val, errs
}
