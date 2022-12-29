package main

import "fmt"

var x int = 5
func f(){
	fmt.Printf("%d", x)
}

func g(){
	fmt.Printf("%d", x)
}
func main(){
	//fmt.Printf(g)
}