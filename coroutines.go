package main

import (
	"fmt"
	"time"
)

func write(phrase string) {
	for i:=0; i<5; i++ {
		fmt.Println(phrase);	
	}
} 

func main() {
	go write("world")
	write("hello")
	time.Sleep(2000)
}