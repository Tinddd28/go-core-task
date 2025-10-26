package t8

type CustomWaitGroup struct {
	sem   chan struct{}
	count int
	done  chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		sem:  make(chan struct{}, 10000), // достаточно большой буфер
		done: make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	for i := 0; i < delta; i++ {
		wg.sem <- struct{}{}
		wg.count++
	}
}

func (wg *CustomWaitGroup) Done() {
	<-wg.sem
	wg.count--
	if wg.count == 0 {
		close(wg.done)
	}
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.done
}
