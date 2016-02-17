### online playground
	http://play.golang.org/
	
### file layout
	package <i>name</i>

	import (
		"<i>package1</i>"
		"<i>package2/sub</i>"
	)

### function definitions
	return multiple values from a single function (typeA, typeB)

		func fn (x, y float64) (float64, int) {

		}

		x, y both float64
		returns two values float64, int

	named return values (only use in obvious situations)
		func fn2 (x float64, y int) (a, b int) {
			a = 1
			b = 2
			return
		}

### variables
	variables can be declared at the function level or the package level.
	variables can be initialized upon declaration
	multiple variables can be declared at once if they are of the same type
	initializers are applied based upon position
	if a variable is not initialized, it is the default value "zero" value for that type.
	a variable's var and type can be omitted if the type is obvious based upon the declaration (k below will be an int)
	one variable can be cast to another type, but only explicitly (u and f below)


	package me

	var x, y, z int = 1, 2, 3

	func fn (a int) (int) {
		k := 3
		u := uint(k)
		f := float64(u)

		return x + y + z + a
	}

### concurrency
	defer statements defers the execution of a function until the surrounding function returns. it acts as a stack so its lifo


### looping
	no parenthese around for loops or if statements
	manadatory brackets {}
	no while loops, use fors

	for {
		// infinite for (while (true))
	}

### if / thens
	no parenthese around if statements
	manadatory brackets {}
	like we're used to in for loops, there can be a little short statement at the beginning of an if statement or a switch statement. 
	the values of those little short statements are available throughout the if ladder

	func fn (x int) {}
		if k:=3+x; k<10 {
			return k
		} else {
			return k - 10
		}
	}

	there are fucking pointers...

### fmt library
* used to print stuff out
* import "fmt"


### Compilation
* go build x.go
	- builds x.go to x
	- run ./x

* go build x.go packages.go

### go get
* go get github.com/mailgun/godebug
	- installs the godebug library
	- will install godebug's src and its bin

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
	a := [2]int {1, 2}
	b := [...]int {1,2}
	c := [2]int
	c[0] = 1
	c[1] = 2

### Slices
* Slices point to a portion (or all of an array). They are declared like arrays, but without the size component. Internally, they point to an array, have a length, and have a capacity. They are descriptor objects.
	- ways to create slices
		a := []int {1,2}
		b := make([]int, 2, 2)
		b[0] = 1
		b[1] = 2
	 	c := make([]int, 2)
		c[0] = 1
		c[1] = 2		
	- slices can be a portion of an array
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
	- arrays can't be grown and since they are the underlying storage mechanism of slices, slices can't be grown either. Instead we must make a new slice and copy the elements from the old to the new one.
		- the new size of the slice is up to the situation. given the situation, we may want to double it, triple it, square it, etc.

### Ranges
* ranges allow for easy iteration over a slice or a map
	for i, v := range m {
		i == index
		v == value
	}

### Maps
* maps are key/value pairs










