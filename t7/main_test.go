package t7

import (
	"testing"
)

func TestMergeIntChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()
	go func() {
		ch2 <- 3
		close(ch2)
	}()
	go func() {
		ch3 <- 4
		ch3 <- 5
		ch3 <- 6
		close(ch3)
	}()

	out := MergeIntChannels(ch1, ch2, ch3)
	var result []int
	for v := range out {
		result = append(result, v)
	}

	// Проверяем, что все значения присутствуют (порядок не гарантируется)
	want := []int{1, 2, 3, 4, 5, 6}
	if !equalUnordered(result, want) {
		t.Errorf("MergeIntChannels() = %v, want %v (unordered)", result, want)
	}
}

// equalUnordered сравнивает два среза без учёта порядка
func equalUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[int]int)
	for _, v := range a {
		counts[v]++
	}
	for _, v := range b {
		counts[v]--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}
