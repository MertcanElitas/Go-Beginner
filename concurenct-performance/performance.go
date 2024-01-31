package concurenctperformance

import (
	"fmt"
	"sync"
)

// Concurrency, async, parallelism her zaman performans kazancı demek değildir. Aşağıdaki örnekte olduğu gibi 
// performans kaybına yol açtığı case de olabilir.

func Run() {
	forLoopGoroutines()
}

func forLoopGoroutines() {
	wg := sync.WaitGroup{}
	wg.Add(10000)
	lock := sync.Mutex{}
	total := 0

	for i := 0; i < 10000; i++ {
		go func() {
			lock.Lock()
			total += 1
			lock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(total)
}

func forLoopPlain() {
	total := 0

	for i := 0; i < 10000; i++ {
		total += i
	}

	fmt.Println(total)
}
