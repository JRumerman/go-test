package main

import (
	"fmt"
)

func main() {
	fmt.Printf("playing with pointers.\n")
	i := 47
	p := &i // assigns p to the address of i

	fmt.Printf("memory address %p\n", p) // prints the memory address
	fmt.Printf("value %d\n", *p) // dereferences the pointer to get the value.
	fmt.Printf("value %b\n", *p) // dereferences the pointer to get the value.
	fmt.Printf("memory address of the pointer %p\n", &p) // 
}