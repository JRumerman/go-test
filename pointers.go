package main

import (
	"fmt"
)

func createSomething() map[string][]string {
	x := make(map[string][]string)
	fmt.Printf("memory address of x in createSomething: %p\n", x)
	return x
}

func main() {
	fmt.Printf("playing with pointers.\n")
	i := 47
	p := &i // assigns p to the address of i

	fmt.Printf("memory address %p\n", p)                 // prints the memory address
	fmt.Printf("value %d\n", *p)                         // dereferences the pointer to get the value.
	fmt.Printf("memory address of the pointer %p\n", &p) //

	y := createSomething()
	fmt.Printf("memory address of y: %p\n", y)
}
