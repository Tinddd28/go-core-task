package main

import "fmt"

func Pipeline(inCh <-chan uint8, outCh chan<- float64) {
	for v := range inCh {
		outCh <- float64(v) * float64(v) * float64(v)
	}
	close(outCh)
}

func main() {
	inCh := make(chan uint8)
	outCh := make(chan float64)

	go Pipeline(inCh, outCh)

	go func() {
		for _, v := range []uint8{2, 3, 4, 5} {
			inCh <- v
		}
		close(inCh)
	}()

	for res := range outCh {
		fmt.Println(res)
	}
}
