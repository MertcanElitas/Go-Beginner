package concurency

import (
	"fmt"
	"sync"
)

type Race struct {
	Val int
}

func RaceExample() {
	raceTest := &Race{}

	wg := &sync.WaitGroup{}
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go Increment(raceTest, wg)
	}

	wg.Wait()

	fmt.Println(raceTest)
}

func Increment(rt *Race, wg *sync.WaitGroup) {
	rt.Val += 1
	wg.Done()
}
