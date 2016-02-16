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

	// slice
	s := []byte{ 'a', 'b', 'c', 'd', 'e', 'f' }
	printSlice("default", s)

	// slices off the first entry in the array and cannot be gotten to ever again, thus reducing the capacity of the slice
	// sets length to 2 [1:3)
	s = s[1:3];
	printSlice("1:3", s)

	// re-slice from 0 position to capacity, but with same length as before 
	s = s[0:]
	printSlice("0:", s)

	// re-slice from 0 position to capacity, resetting length to capacity
	s = s[:cap(s)]
	printSlice("0:cap(s)", s)

	// reslice from position 0 to 4, 
	s = s[:4]
	printSlice(":4", s)

	// create a full slice from s
	e := s[:cap(s)]
	printSlice("e creation", e)

	e[0] = 'z';
	fmt.Printf("\ne[0] - %c\n", e[0])
	fmt.Printf("s[0] - %c\n", s[0])
}

func printSlice(comment string, s []byte){
	fmt.Printf("\n")
	fmt.Printf("%s\n", comment)
	fmt.Printf("length %d\n", len(s))
	fmt.Printf("capacity %d\n", cap(s))
	for i:=0; i<len(s); i++ {
		fmt.Printf("s[%d] - %c\n", i, s[i])
	}	
}