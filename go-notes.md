[online playground](http://play.golang.org/)

### go commands
#### compilation
- `go install x.go`
	- compiles and moves output to bin folder (defined by $GOBIN)
- `go run x.go`
	- compiles and runs output in place
- `go build x.go`
	- creates a package if folder is provided as input
	- creates an executable if .go files are provided as input
- `go test`
	- runs test that are defined in a particular way.
``` go
import "testing"

func TestReverse(t *testing.T) {}
```
	- file is named with suffix "_test"
#### getting packages
- 'go get github.com/mailgun/godebug'
	- installs the godebug library
	- will install godebug's src and its bin
	
### file layout
``` go
package name

import (
	"package1"
	"package2/sub"
)
```

### function definitions
- return multiple values from a single function (typeA, typeB)
``` go
func fn (x, y float64) (float64, int) {

}

// x, y both float64
// returns two values float64, int

// named return values (only use in obvious situations)
func fn2 (x float64, y int) (a, b int) {
	a = 1
	b = 2
	return
}
```
- pass a variable by value; default behavior `(v Vertex)`
- pass a variable by reference; pass in a pointer `(v *Vertex)`

### variables
- variables can be declared at the function level or the package level.
- variables can be initialized upon declaration
- multiple variables can be declared at once if they are of the same type
- initializers are applied based upon position
- if a variable is not initialized, it is the default value "zero" value for that type.
- a variable's var and type can be omitted if the type is obvious based upon the declaration (k below will be an int)
- one variable can be cast to another type, but only explicitly (u and f below)

``` go
package me

var x, y, z int = 1, 2, 3

func fn (a int) (int) {
	k := 3
	u := uint(k)
	f := float64(u)

	return x + y + z + a
}
```

### concurrency
- defer statements defers the execution of a function until the surrounding function returns. it acts as a stack so its lifo


### looping
- no parenthese around for loops or if statements
- manadatory brackets {}
- no while loops, use fors

``` go
for {
	// infinite for (while (true))
}
```

### if / thens
- no parenthese around if statements
- manadatory brackets {}
- like we're used to in for loops, there can be a little short statement at the beginning of an if statement or a switch statement. 
- the values of those little short statements are available throughout the if ladder

``` go
func fn (x int) {
	if k:=3+x; k<10 {
		return k
	} else {
		return k - 10
	}
}
```

there are fucking pointers...

### fmt library
* used to print stuff out
* import "fmt"


### Debugging
* use mailgun's godebug. introduced in 1/2015 so pretty new. before that it was gdb.
	- https://github.com/mailgun/godebug
* go get github.com/mailgun/godebug
	- installs the godebug library
* insert a '_ = "breakpoint"' to get a breakpoint going
* godebug run x.go
	- starts the debugger for x; stopping at the first breakpoint

### Arrays
* Arrays are fixed sized structs. They are declared with a fixed size and that size value cannot be changed. They are not used directly very often, but are used as the storage mechanism for slices.
``` go
a := [2]int {1, 2}
b := [...]int {1,2}
c := [2]int
c[0] = 1
c[1] = 2
```
### Slices
* Slices point to a portion (or all of an array). They are declared like arrays, but without the size component. Internally, they point to an array, have a length, and have a capacity. They are descriptor objects.
	- ways to create slices
	``` go
		a := []int {1,2}
		b := make([]int, 2, 2)
		b[0] = 1
		b[1] = 2
	 	c := make([]int, 2)
		c[0] = 1
		c[1] = 2		
	```
	- slices can be a portion of an array
	``` go
		b := []byte{'r', 'u', 'm', 'e', 'r', 'm', 'a', 'n'}
		c := b[1:3] // 'u', 'm' -- elements 1 through 2
			// c would have ptr to b[1], length 2, capacity 7 
		c[0] == 'u' // positions are 0 and 1
		c[1] == 'm'
		d := b[:3] // positions 0, 1, 2
		e := b[2:] // positions 2,3,4,5,6,7
		c = c[:cap(c)] // positions 1,2,3,4,5,6,7 (umerman)
			- regrow the array to the capacity of the underlying array.
			- can't get back to zero though. 
	```
	- arrays can't be grown and since they are the underlying storage mechanism of slices, slices can't be grown either. Instead we must make a new slice and copy the elements from the old to the new one.
		- the new size of the slice is up to the situation. given the situation, we may want to double it, triple it, square it, etc.

### Ranges
* ranges allow for easy iteration over a slice or a map
	``` go
	for i, v := range m {
		i == index
		v == value
	}
	```

### Maps
* maps are key/value pairs

### Functions as variables
- pass them around just like in JavaScript. Sweet! Something I recognize.
- they can be returned from another function. awesome! just like JS.
- they are assigned to variables
	``` go
	x := func (a, b int) int {
		return a+b
	}
	```
	-
### Methods
- Go does not have classes but you can still define methods on types. Methods are attached to a type using a special argument called the Receiver argument.
- A receiver argument can be a value or can be a pointer. 
	- Value
		- Receives a copy of the receiver
	- Pointer
		- Receives a pointer to the receiver and therefore can modify the receiver directly
- If a receiver argument method has a value or a pointer type specified, Go fixes incorrectly called methods such that the right one is calld.
	- `(&v).Abs()`
	- `v.Abs()`
		- same thing regardless if the receiver is a pointer or a value.
- For consistency, all methods on a given type should have pointer or value receivers. We shouldn't mix and match


``` go
package main

import (
	"fmt"
	"math"
)

type Vertex struct { 
	x, y float64 
}

func (v Vertex) Distance (u Vertex) float64 {
	return math.Sqrt(math.Pow((v.x - u.x), 2) + math.Pow((v.y - u.y), 2))
}

func main () {
	a, b := Vertex {-2, -3 }, Vertex{-4, 4}
	fmt.Println(a.Distance(b))
}

```

### Interfaces
- An interface is a type. 

``` go
type IAbs interface {
	Abs() float64
}
```

- interfaces are implemented implicitly. There is no code that says type X implements interface A. This means it is very easy to implement an interface

``` go
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
	fmt.Println(a.Abs())
	a = c
	fmt.Println(a.Abs())
	// a = b
	// fmt.Println(a.Abs())
}
```