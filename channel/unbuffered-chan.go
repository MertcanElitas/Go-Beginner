package channel

import (
	"fmt"
	"time"
)

func UnBufferedChannel() {
	//channel initialization
	unbufferedChan := make(chan int)

	//channel declaration
	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2) //nil

	//only can read from channel
	go func(unbufChan <-chan int) {
		//blocking code until setted data to unbufChan
		value := <-unbufChan
		fmt.Println(value)

		//unbufChan <- 5 this is not work here
	}(unbufferedChan)

	time.Sleep(time.Second * 5)

	unbufferedChan <- 1

	time.Sleep(time.Second * 2)
}
