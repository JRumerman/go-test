package main

import (
	"fmt"
)

func main() {
	// array
	var x [2] int
	x[0] = 100
	x[1] = 200

	fmt.Printf ("%v\n", x)
	fmt.Printf ("%v %v\n", x[0], x[1])

	for i:=0; i<len(x); i++ {
		fmt.Printf("x[%d] - %d\n", i, x[i])
	}

	// array
	var y = [...]string{"Ella", "Logan", "Joel", "Stacey"}
	for i:=0; i<len(y); i++ {
		fmt.Printf("y[%d] - %s\n", i, y[i])
	}
}