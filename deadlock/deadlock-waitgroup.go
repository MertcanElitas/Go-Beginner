package deadlock

import (
	"fmt"
	"sync"
)

// Waitgroup'a bekliyeceğinden daha çok sayıda değer verirsek deadlock oluşur.

func DeadlockWaitgroup() {
	//waitgroup for synchronization
	wg := sync.WaitGroup{}
	//wrong number of goroutine to wait
	wg.Add(2)

	go func() {
		fmt.Println("hello from go routine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("hello concurrency")
}
