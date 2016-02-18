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
