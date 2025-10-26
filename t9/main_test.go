package main

import (
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	inCh := make(chan uint8)
	outCh := make(chan float64)

	go Pipeline(inCh, outCh)

	go func() {
		for _, v := range []uint8{1, 2, 3} {
			inCh <- v
		}
		close(inCh)
	}()

	var results []float64
	for res := range outCh {
		results = append(results, res)
	}

	want := []float64{1, 8, 27}
	if !reflect.DeepEqual(results, want) {
		t.Errorf("Pipeline() = %v, want %v", results, want)
	}
}
