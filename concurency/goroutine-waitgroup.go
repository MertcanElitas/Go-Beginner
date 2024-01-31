package concurency

import (
	"fmt"
	"sync"
)

func PrintGoroutineWithWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("hello from goroutine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("hello from main")
}
