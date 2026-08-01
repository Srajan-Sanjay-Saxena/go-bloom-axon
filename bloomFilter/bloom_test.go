package bloomfilter

import "testing"

func TestAdd_And_Contains(t *testing.T) {
	bf := New(100)
	bf.Add([]byte("hello"))
	if !bf.Contains([]byte("hello")) {
		t.Error("expected 'hello' to be in the filter")
	}
}

func TestContains_AbsentItem(t *testing.T) {
	bf := New(100)
	bf.Add([]byte("hello"))
	if bf.Contains([]byte("world")) {
		t.Log("false positive for 'world' (acceptable but worth noting)")
	}
}

func TestContains_EmptyFilter(t *testing.T) {
	bf := New(100)
	if bf.Contains([]byte("anything")) {
		t.Error("empty filter should not contain any item")
	}
}

func TestAdd_MultipleItems(t *testing.T) {
	bf := New(100)
	items := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	for _, item := range items {
		bf.Add(item)
	}
	for _, item := range items {
		if !bf.Contains(item) {
			t.Errorf("expected %q to be in the filter", item)
		}
	}
}

func TestAdd_EmptyBytes(t *testing.T) {
	bf := New(100)
	bf.Add([]byte(""))
	if !bf.Contains([]byte("")) {
		t.Error("expected empty byte slice to be in the filter after Add")
	}
}
