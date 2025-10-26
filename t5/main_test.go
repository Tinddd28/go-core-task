package t5

import (
	"reflect"
	"testing"
)

func TestIntersectInts(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	wantBool := true
	wantSlice := []int{64, 3}
	gotBool, gotSlice := IntersectInts(a, b)
	if gotBool != wantBool || !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Errorf("IntersectInts() = (%v, %v), want (%v, %v)", gotBool, gotSlice, wantBool, wantSlice)
	}
}

func TestIntersectInts_NoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	wantBool := false
	wantSlice := []int{}
	gotBool, gotSlice := IntersectInts(a, b)
	if gotBool != wantBool || !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Errorf("IntersectInts() = (%v, %v), want (%v, %v)", gotBool, gotSlice, wantBool, wantSlice)
	}
}

func TestIntersectInts_Duplicates(t *testing.T) {
	a := []int{1, 2, 2, 3}
	b := []int{2, 2, 3, 3}
	wantBool := true
	wantSlice := []int{2, 3}
	gotBool, gotSlice := IntersectInts(a, b)
	if gotBool != wantBool || !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Errorf("IntersectInts() = (%v, %v), want (%v, %v)", gotBool, gotSlice, wantBool, wantSlice)
	}
}
