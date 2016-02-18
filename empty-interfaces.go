package main

import (
	"fmt"
	"math"
)

type IAbs interface {
	Abs() float64
}

type Vertex struct { 
	X, Y float64 
}

type MyFloat float64

func (v *Vertex) Abs () float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if (f >= 0) {
		return float64(f)
	}

	return float64(-f)
}

func main () {
	var a IAbs
	b, c := Vertex {-2, -3 }, MyFloat(-58.548)
	a = &b
	EmptyDescribe(a)
	a = c
	EmptyDescribe(a)

	var i interface {} = "string"
	s := i.(string)
	fmt.Printf("%s\n", s)

	s, ok := i.(string)
	fmt.Printf("Ok: %t, val: %s\n", ok, s)

	f, ok := i.(float64)
	fmt.Printf("Ok: %t, val: %f\n", ok, f)
}

func EmptyDescribe (i interface{}) {
	// can't call Abs here on i
	fmt.Printf("Empty Describe: (%v, %T)\n", i, i)
}