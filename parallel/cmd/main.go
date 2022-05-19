package main

import (
	"fmt"
	"time"

	"github.com/Knowckx/infa"
	"github.com/Knowckx/infa/parallel"
)

// used to test parallel
func main() {
	pa := parallel.NewParallel(4)
	for i := 0; i < 5; i++ {
		v := i
		pa.Apply()
		go ParallelDo(pa, v)
	}
	infa.BeStuck()
}

func ParallelDo(pa *parallel.Parallel, i int) {
	args := fmt.Sprintf("%d", i)
	Worker(args)
	pa.Done()
}

func Worker(info string) {
	fmt.Println("Working", info)
	time.Sleep(3 * time.Second)
	fmt.Println("Working Finidsh", info)
}
