package authorize

// import (
// 	"errors"
// 	"fmt"

// 	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
// 	st "github.com/aasanchez/ocpp16messages/shared/types"
// )

// var ErrInvalidRequestIdTag = errors.New("request -> invalid idTag")

// type RequestMessage struct {
// 	IdTag mat.IdToken
// }

// func Request(input mat.RequestPayload) (RequestMessage, error) {
// 	str, err := st.CiString20Type(input.IdTag)
// 	if err != nil {
// 		wrapped := fmt.Errorf("request -> invalid idTag -> %w", err)

// 		return RequestMessage{}, wrapped
// 	}

// 	idToken, _ := mat.IdToken(str)

// 	return RequestMessage{IdTag: idToken}, nil
// }
