// +build OMIT

package main

import "fmt"

func main() {
	sum := 1
	//init and post statements are optional.
	for ; sum < 1000 ; {
		sum += sum
	}
	fmt.Println(sum)
}
