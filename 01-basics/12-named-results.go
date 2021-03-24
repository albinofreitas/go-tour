// +build OMIT

package main

import (
	"fmt"
	"log"
	"errors"
)

func div(x int, y int) (r int, err error) {
	if y == 0 {
		err = errors.New("cannot divide by zero")
		return
	}

	r = x / y
	return
}

func main() {
	r, err := div(100, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}
