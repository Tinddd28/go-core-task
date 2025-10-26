package t3

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("foo", 42)
	val, ok := m.Get("foo")
	if !ok || val != 42 {
		t.Errorf("Add/Get failed: got (%v, %v), want (42, true)", val, ok)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("bar", 7)
	m.Remove("bar")
	_, ok := m.Get("bar")
	if ok {
		t.Errorf("Remove failed: key 'bar' should not exist")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)
	copied := m.Copy()
	if !reflect.DeepEqual(copied, map[string]int{"a": 1, "b": 2}) {
		t.Errorf("Copy failed: got %v, want %v", copied, map[string]int{"a": 1, "b": 2})
	}
	m.Add("c", 3)
	if _, exists := copied["c"]; exists {
		t.Errorf("Copy is not independent: copied map changed after original")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 100)
	if !m.Exists("key") {
		t.Errorf("Exists failed: key should exist")
	}
	if m.Exists("nope") {
		t.Errorf("Exists failed: key should not exist")
	}
}
