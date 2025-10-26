package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	want := []int{2, 4, 6}
	got := sliceExample(slice)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sliceExample() = %v, want %v", got, want)
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3}
	got := addElements(slice, 4)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements() = %v, want %v", got, want)
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3}
	copied := copySlice(slice)
	slice[0] = 99
	if copied[0] == 99 {
		t.Errorf("copySlice() did not copy correctly")
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	got := removeElement(slice, 1)
	want := []int{1, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("removeElement() = %v, want %v", got, want)
	}
}
