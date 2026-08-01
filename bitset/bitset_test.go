package bitset

import (
	bse "go-bloom-axon/bitset_error"
	"testing"
)

func TestNew(t *testing.T) {
	bs := New(100)
	if bs.Size() != 100 {
		t.Errorf("expected size 100, got %d", bs.Size())
	}
}

func TestSet_And_Get(t *testing.T) {
	bs := New(100)

	if err := bs.Set(50); err != nil {
		t.Fatalf("unexpected error on Set: %v", err)
	}

	ok, err := bs.Get(50)
	if err != nil || !ok {
		t.Errorf("expected bit 50 to be set")
	}
}

func TestGet_UnsetBit(t *testing.T) {
	bs := New(100)
	ok, err := bs.Get(50)
	if err != nil || ok {
		t.Errorf("expected bit 50 to be unset")
	}
}

func TestSet_BoundaryBit(t *testing.T) {
	bs := New(64)
	if err := bs.Set(63); err != nil {
		t.Fatalf("unexpected error on Set(63): %v", err)
	}
	ok, _ := bs.Get(63)
	if !ok {
		t.Error("expected bit 63 to be set")
	}
}

func TestSet_OutOfBounds(t *testing.T) {
	bs := New(100)
	err := bs.Set(100)
	if err != bse.ErrIndexOutOfBounds {
		t.Errorf("expected ErrIndexOutOfBounds, got %v", err)
	}
}

func TestGet_OutOfBounds(t *testing.T) {
	bs := New(100)
	_, err := bs.Get(100)
	if err != bse.ErrIndexOutOfBounds {
		t.Errorf("expected ErrIndexOutOfBounds, got %v", err)
	}
}

func TestSet_MultipleBits(t *testing.T) {
	bs := New(200)
	bits := []uint64{0, 63, 64, 127, 199}
	for _, b := range bits {
		if err := bs.Set(b); err != nil {
			t.Fatalf("unexpected error setting bit %d: %v", b, err)
		}
	}
	for _, b := range bits {
		ok, err := bs.Get(b)
		if err != nil || !ok {
			t.Errorf("expected bit %d to be set", b)
		}
	}
}
