package main

import (
	"fmt"
)

func main(){
	var decimalNum float32

	fmt.Printf("Enter a decimal number")
	fmt.Scanf("%f", &decimalNum)

	fmt.Printf("The truncated value of %f, is %d\n", decimalNum, int(decimalNum))
}