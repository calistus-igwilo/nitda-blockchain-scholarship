package main

import (
	"fmt"
)

func main() {

	var x int = 1
	y := 1	
	
	if true {
		x := 2
		y = 2
		fmt.Println(x)
		fmt.Println(y)
	}
	
	fmt.Println(x)
	fmt.Println(y)
}