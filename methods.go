package main

import (
	"fmt"
	"math"
)

type Vertex struct { 
	x, y float64 
}

type MyFloat float64

func (v Vertex) Distance (u Vertex) float64 {
	return math.Sqrt(math.Pow((v.x - u.x), 2) + math.Pow((v.y - u.y), 2))
}

func (f MyFloat) Abs() float64 {
	if (f >= 0) {
		return float64(f)
	}

	return float64(-f)
}

func (v *Vertex) Invert () {
	v.x = -v.x
	v.y = -v.y
}

func main () {
	a, b, c := Vertex {-2, -3 }, Vertex{-4, 4}, MyFloat(-58.548)
	fmt.Println(a.Distance(b))
	a.Invert()
	fmt.Println(a)
	fmt.Println(c.Abs())
}
