package bootnotificationtypes

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// type CurrentTime struct {
// 	t time.Time
// }

// // Now returns the current UTC time as CurrentTime.
// func Now() CurrentTime {
// 	return CurrentTime{t: time.Now().UTC()}
// }

// // NewFromRFC3339 parses a string in RFC3339 format into a CurrentTime.
// func FromRFC3339(input string) (CurrentTime, error) {
// 	t, err := time.Parse(time.RFC3339, input)
// 	if err != nil {
// 		return CurrentTime{}, fmt.Errorf("invalid RFC3339 time: %w", err)
// 	}

// 	return CurrentTime{t: t}, nil
// }

// // NewFromEpoch converts a Unix epoch timestamp (in seconds) to CurrentTime.
// func FromEpoch(epoch int64) CurrentTime {
// 	return CurrentTime{t: time.Unix(epoch, 0).UTC()}
// }

// // ToRFC3339 returns the string representation of the time in RFC3339 format.
// func (ct CurrentTime) ToRFC3339() string {
// 	return ct.t.Format(time.RFC3339)
// }

// // ToEpoch returns the Unix timestamp (in seconds) of the CurrentTime.
// func (ct CurrentTime) ToEpoch() int64 {
// 	return ct.t.Unix()
// }

// // FromString tries to parse a general date string and convert it into CurrentTime.
// // It only accepts valid RFC3339 strings.
// func FromString(input string) (CurrentTime, error) {
// 	return FromRFC3339(input)
// }

// // ToString is an alias for ToRFC3339.
// func (ct CurrentTime) ToString() string {
// 	return ct.ToRFC3339()
// }

// // ValidateRFC3339 checks if a string is a valid RFC3339 timestamp.
// func ValidateRFC3339(input string) error {
// 	_, err := time.Parse(time.RFC3339, input)
// 	if err != nil {
// 		return errors.New("input is not a valid RFC3339 timestamp")
// 	}

// 	return nil
// }
