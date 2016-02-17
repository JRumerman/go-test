package main

import (
	"fmt"
)

func main() {
	// slice
	s := []byte{'a', 'b', 'c', 'd', 'e', 'f'}

	// ranges loop over slices and maps and produce two values: index and value
	for i, v := range s {
		fmt.Printf("i: %d, v: %c\n", i, v)
	}
}
