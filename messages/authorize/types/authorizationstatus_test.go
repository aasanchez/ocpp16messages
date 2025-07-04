package authorizetypes

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func Test_getStatusSet_containsAllConstants(t *testing.T) {
	t.Parallel()

	m := getStatusSet()
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	want := []string{Accepted, Blocked, ConcurrentTx, Expired, Invalid}
	sort.Strings(want)
	if !reflect.DeepEqual(keys, want) {
		t.Errorf("getStatusSet keys = %v; want %v", keys, want)
	}
}

func TestSetAuthorizationStatus_invalidWrapsErrInvalidAuthorizationStatus(t *testing.T) {
	t.Parallel()

	input := "NotAStatus"
	_, err := SetAuthorizationStatus(input)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, st.ErrInvalidAuthorizationStatus) {
		t.Errorf("error %v; want to wrap st.ErrInvalidAuthorizationStatus", err)
	}

	wantMsg := fmt.Sprintf("%s: %q", st.ErrInvalidAuthorizationStatus, input)
	if err.Error() != wantMsg {
		t.Errorf("error message = %q; want %q", err.Error(), wantMsg)
	}
}
