package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	p := Point { 1, 2 }
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%v\n", p.x)
	fmt.Printf("%v\n", p.y)

	l := &p
	l.x = 100
	fmt.Printf("%#v\n", l)

	x := Point { x: 8 }
	y := Point {}
	z := &Point { 2,3 }
	fmt.Printf("%v %v %v\n", x,y,z)
}