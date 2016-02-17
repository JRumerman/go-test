package main

import (
	"fmt"
)

func main() {
	// slice
	s := []byte{'a', 'b', 'c', 'd', 'e', 'f'}

	// complete copy for later
	u := s[:]

	printSlice("default", s)

	// slices off the first entry in the array and cannot be gotten to ever again, thus reducing the capacity of the slice
	// sets length to 2 [1:3)
	s = s[1:3]
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

	// a slice points to the same underlying array, no matter how many times it is resliced.
	e[0] = 'z'
	fmt.Printf("\ne[0] - %c\n", e[0])
	fmt.Printf("s[0] - %c\n", s[0])

	// increase the capacity of by making a copy and reassigning
	z := make([]byte, len(u), (cap(u)+1)*2)
	copy(z, u)
	u = z
	printSlice("copy of u with increased capacity", u)

	for j := 0; j < 10; j++ {
		u = append(u, 'z')
	}
	printSlice("append u + z - 10 times", u)
}

func printSlice(comment string, s []byte) {
	fmt.Printf("\n")
	fmt.Printf("%s len: %d cap: %d\n", comment, len(s), cap(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] - %c\n", i, s[i])
	}
}
