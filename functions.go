package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func filter(list []int, fn func(int) bool) []int {
	ret := []int{}
	for _, v := range list {
		if fn(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func main() {
	fmt.Println(add(42, 13))

	area := func(x, y int) int {
		return x * y
	}

	isEven := func(val int) bool {
		return val%2 == 0
	}

	fmt.Println(area(10, 20))

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	list = filter(list, isEven)
	fmt.Println(list)
}
