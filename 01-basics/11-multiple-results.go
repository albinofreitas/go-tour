// +build OMIT

package main

import (
	"fmt"
	"log"
	"errors"
)

func div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return x / y, nil
}

func main() {
	r, err := div(100, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}
