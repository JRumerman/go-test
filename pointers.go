package main

import (
	"fmt"
	"unsafe"
)

func createSomething() map[string][]string {
	x := make(map[string][]string)
	fmt.Printf("memory address of x in createSomething: %p\n", unsafe.Pointer(&x))
	return x
}

func passSomething(data map[string][]string) {
	fmt.Printf("memory address of data in passSomething: %p\n", unsafe.Pointer(&data))
}

func main() {
	// fmt.Printf("playing with pointers.\n")
	// i := 47
	// p := &i // assigns p to the address of i

	// fmt.Printf("memory address %p\n", p)                 // prints the memory address
	// fmt.Printf("value %d\n", *p)                         // dereferences the pointer to get the value.
	// fmt.Printf("memory address of the pointer %p\n", &p) //

	y := createSomething()
	passSomething(y)
	fmt.Printf("memory address of y: %p\n", unsafe.Pointer(&y))

	// data := make([]byte, 10)
	// data[0] = 1
	// data[1] = 2
	// fmt.Printf("data[0] %v\n", data[0])
	// fmt.Printf("memory address of data: %p\n", data)
	// passSomething(data)
	// fmt.Printf("data[0] %v\n", data[0])
}
