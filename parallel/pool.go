package parallel

// type stuct

type Parallel struct {
	maxCount int
	sem      chan int
}

func NewParallel(max int) *Parallel {
	out := &Parallel{}
	out.maxCount = max
	out.sem = make(chan int, max)
	return out
}

func (t *Parallel) Apply() {
	t.sem <- 1
}

func (t *Parallel) Done() {
	<-t.sem
}

func (t *Parallel) DoFunc(handle func()) {
	t.Apply()
	go t.parallelDo(handle)
}

func (t *Parallel) parallelDo(handle func()) {
	handle()
	t.Done()
}
