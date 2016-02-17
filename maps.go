package main

import (
	"fmt"
)

func main() {
	// map
	m := make(map[string]int)
	m["joel"] = 36
	m["stacey"] = 36
	m["logan"] = 5
	m["ella"] = 5
	PrintMap(m)

	// declare directly inline
	n := map[string]int {
		"keri":39,
		"howard":72,
	}

	PrintMap(n)
}

func PrintMap(m map[string]int) {
	fmt.Printf("%v\n", m)

	// range over a map produces a key/value pair
	for i,v:=range m {
		fmt.Printf("%v %v\n", i, v)
	}
}