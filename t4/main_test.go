package t4

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	want := []string{"apple", "cherry", "43", "lead", "gno1"}
	got := Difference(slice1, slice2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Difference() = %v, want %v", got, want)
	}
}

func TestDifference_EmptySecond(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{}
	want := []string{"a", "b"}
	got := Difference(slice1, slice2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Difference() = %v, want %v", got, want)
	}
}

func TestDifference_AllOverlap(t *testing.T) {
	slice1 := []string{"x", "y"}
	slice2 := []string{"x", "y"}
	want := []string{}
	got := Difference(slice1, slice2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Difference() = %v, want %v", got, want)
	}
}
