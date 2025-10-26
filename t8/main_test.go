package t8

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()
	counter := 0

	wg.Add(3)
	go func() {
		time.Sleep(10 * time.Millisecond)
		counter++
		wg.Done()
	}()
	go func() {
		time.Sleep(20 * time.Millisecond)
		counter++
		wg.Done()
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		counter++
		wg.Done()
	}()

	wg.Wait()
	if counter != 3 {
		t.Errorf("Expected counter to be 3, got %d", counter)
	}
}

func TestCustomWaitGroupZero(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(0)
	go func() {
		wg.Wait()
	}()
}
