package bitset_error

import "testing"

func TestErrIndexOutOfBounds_Error(t *testing.T) {
	expected := "status : 400 , message : Index out of bounds"
	if got := ErrIndexOutOfBounds.Error(); got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}
