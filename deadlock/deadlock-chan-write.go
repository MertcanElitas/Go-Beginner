package deadlock

import (
	"fmt"
)

// Unbuffered bir channel'a hiç bir subscriber olmadığı sürece değer yazmaya çalışırsak ve tek bir goroutine varsa
// bu bir deadlock oluşturur. 

func main() {

	//unbuffered channel
	c1 := make(chan int)
	//writes are blocking operations
	c1 <- 1

	fmt.Println("hello concurrency")

}
