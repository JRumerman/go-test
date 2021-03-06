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
	AbsDescribe(a)
	a = c
	AbsDescribe(a)
}

func AbsDescribe (i IAbs) {
	fmt.Printf("Abs Describe: (%v, %T)\n", i.Abs(), i)
}