package t7

import (
	"testing"
)

func TestRandGenerator(t *testing.T) {
	ch, stop := RandGenerator()
	defer stop()

	nums := make(map[int]struct{})
	for i := 0; i < 5; i++ {
		n, ok := <-ch
		if !ok {
			t.Fatalf("Channel closed unexpectedly")
		}
		nums[n] = struct{}{}
	}

	if len(nums) < 2 {
		t.Errorf("RandGenerator produced too few unique numbers: %v", nums)
	}

	stop()
	_, ok := <-ch
	if ok {
		t.Errorf("Channel should be closed after stop()")
	}
}
