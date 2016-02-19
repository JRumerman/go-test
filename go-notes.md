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
	// i == index
	// v == value
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
	AbsDescribe(a)
	a = c
	AbsDescribe(a)
}

func AbsDescribe (i IAbs) {
	fmt.Printf("Abs Describe: (%v, %T)\n", i.Abs(), i)
}
```

### Type Assertions / Empty Interface
- the empty interface is a special interface implemented by all types. because it is an empty interface it can be used as a generic-like parameter.
- casting the empty interface to another type is done using the `var f, ok := x.(*type*)` syntax.
	- f is the value of type T if the cast was successful
	- ok is a boolean indicating if the cast was successful
- if the ok is omitted, and the cast fails, the system panics (throws an error)
	- similar to the <T>.tryParse(T val, out bool success) in .NET

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
```
- type switches are special switch statements that switch over a type rather than a value. they are used against the empty interface.
```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

### Errors
- errors are simply types that have implemented the error interface. the error interface is defined as the following:
```go
type error interface {
    Error() string
}
```

therefore, to handle errors, add a return parameter of the error interface type and create a new instance of a custom error object when an error occurs and return it.

```go
package main

import (
	"fmt"
)

type ErrNegativeInt int

func (e ErrNegativeInt) Error() string {
	return fmt.Sprintf("%v was negative.", int(e))
}

func do (v int) (int, error) {
	if v < 0 {
		return 0, ErrNegativeInt(v)
	}

	return v+1, nil
}

func main(){
	fmt.Println(do (10))
	fmt.Println(do (-1))
	fmt.Println(do (-2))
}
```

### Readers
- readers are another interface. they are used to populate a byte array from the underlying data.
- they are used a lot by the internal packages
``` go
package main

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	l := len(b)
	for i := 0; i < l; i++ {
		b[i] = byte('A')
	}

	return l, nil
}

func main() {
	reader.Validate(MyReader{})
}
```
- another reader example
``` go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func translate(v byte) byte {
	var diff int8
	if (v > 64 && v < 79) || (v > 96 && v < 109) {
		diff = 13
	}

	if (v > 78 && v < 90) || (v > 108 && v < 123) {
		diff = -13
	}

	return byte(int8(v) + diff)
}

func (a rot13Reader) Read(c []byte) (int, error) {
	n, err := a.r.Read(c)

	for i, v := range c {
		c[i] = translate(v)
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

### concurrency
- defer statements defers the execution of a function until the surrounding function returns. it acts as a stack so its lifo

#### goroutines
- lightweight thread managed by go.
- they are simply functions that are called with the `go` keyword prefix
- the calling stuff is evaluated before the method begins; within the current thread (e.g. the method and the variables)
``` go
package main

import (
	"fmt"
	"time"
)

func write(phrase string) {
	for i:=0; i<5; i++ {
		fmt.Println(phrase);	
	}
} 

func main() {
	go write("world")
	write("hello")
	time.Sleep(2000)
}
```

#### channels
- a way of synchronizing data across multiple threads.
- waiting for data before continuing.
``` go
package main

import (
	"fmt"
	"time"
)

func product(values []int, c chan int) {
	p := 1
	for _, v := range values {
		p *= v
	}

	c <- p
}

func main() {
	c := make(chan int)
	go product([]int{1, 2, 3, 4, 5, 6, 7, 8}, c)
	fmt.Println(<-c)
}
```

- channels can be buffered. buffering a channel means that sends to the channel will block if the buffer is full and receives will block if the buffer is empty.

