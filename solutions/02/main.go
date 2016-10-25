// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

}
