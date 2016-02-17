package main

import (
	"fmt"
)

func main() {
	_ = "breakpoint"

	fmt.Printf("playing with formatters.\n\n")
	i := 47
	var p *int
	p = &i // assigns p to the address of i

	fmt.Printf("memory address (%%p) %p\n", p) // %p
	fmt.Printf("types (%%T) %T\n", i)          // %p

	fmt.Printf("default value (%%v) %v\n", i)
	fmt.Printf("integer base 10 (%%d) %d\n", *p) // %d
	fmt.Printf("integer base 2 (%%b) %b\n", i)   // %b
	fmt.Printf("bool (%%t) %t\n", true)
	fmt.Printf("float (%%f) %f\n", 5.0)
	fmt.Printf("scientifc (%%e) %e\n", 5000000.0)
	fmt.Printf("scientifc (%%E) %E\n", 5000000.0)
	fmt.Printf("strings (%%s) %s\n", "Yo! This is a string.")
	fmt.Printf("quoted strings (%%q) %q\n", "Yo! This is a quoted string.")
	fmt.Printf("width specific integers (%%6d) |%6d|%6d|\n", 12, 345)
	fmt.Printf("width specific floats (%%6.2f) |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width specific floats with left justifications (%%-6.2f) |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width specific strings (%%6s) |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width specific strings with left justifications (%%-6s) |%-6s|%-6s|\n", "foo", "b")
}
