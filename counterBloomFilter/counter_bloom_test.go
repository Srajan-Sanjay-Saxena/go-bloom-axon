package counterbloomfilter

import "testing"

func TestAdd_And_Contains(t *testing.T) {
	cbf := New(100)
	cbf.Add([]byte("hello"))
	if !cbf.Contains([]byte("hello")) {
		t.Error("expected 'hello' to be present after Add")
	}
}

func TestContains_EmptyFilter(t *testing.T) {
	cbf := New(100)
	if cbf.Contains([]byte("hello")) {
		t.Error("empty filter should not contain any item")
	}
}

func TestRemove_AfterAdd(t *testing.T) {
	cbf := New(100)
	cbf.Add([]byte("hello"))
	if err := cbf.Remove([]byte("hello")); err != nil {
		t.Fatalf("unexpected error on Remove: %v", err)
	}
	if cbf.Contains([]byte("hello")) {
		t.Error("expected 'hello' to be absent after Remove")
	}
}

func TestRemove_WithoutAdd(t *testing.T) {
	cbf := New(100)
	if err := cbf.Remove([]byte("ghost")); err != ErrUnderflow {
		t.Errorf("expected ErrUnderflow, got %v", err)
	}
}

func TestAdd_MultipleItems_IndependentRemoval(t *testing.T) {
	cbf := New(100)
	cbf.Add([]byte("foo"))
	cbf.Add([]byte("bar"))

	cbf.Remove([]byte("foo"))

	if cbf.Contains([]byte("foo")) {
		t.Error("expected 'foo' to be absent after Remove")
	}
	if !cbf.Contains([]byte("bar")) {
		t.Error("expected 'bar' to still be present")
	}
}

func TestAdd_SameItemTwice_RequiresTwoRemoves(t *testing.T) {
	cbf := New(100)
	cbf.Add([]byte("dup"))
	cbf.Add([]byte("dup"))

	cbf.Remove([]byte("dup"))
	if !cbf.Contains([]byte("dup")) {
		t.Error("expected 'dup' to still be present after one Remove (added twice)")
	}

	cbf.Remove([]byte("dup"))
	if cbf.Contains([]byte("dup")) {
		t.Error("expected 'dup' to be absent after two Removes")
	}
}

func TestAdd_CounterOverflow(t *testing.T) {
	cbf := New(100)
	for i := 0; i < 255; i++ {
		if err := cbf.Add([]byte("flood")); err != nil {
			t.Fatalf("unexpected error at iteration %d: %v", i, err)
		}
	}
	if err := cbf.Add([]byte("flood")); err != ErrCounterMaxed {
		t.Errorf("expected ErrCounterMaxed, got %v", err)
	}
}
