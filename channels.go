package main

import (
	"fmt"
	"time"
)

func product(values []int, c chan int) {
	p := 1
	for _, v := range values {
		p *= v
	}

	time.Sleep(5 * time.Second)

	c <- p
}

func main() {
	c := make(chan int)
	go product([]int{1, 2, 3, 4, 5, 6, 7, 8}, c)
	fmt.Println(<-c)
}
