package main

import "fmt"

func main(){
	/* var x [5]int

	x[2] = 6

	fmt.Printf("%d\n", x[2])
	*/
	
	// var x [5]int = [5]{1, 2, 3, 4, 5}

	/* x := [...]int {1, 2, 3, 4}

	for i, v := range x {
		fmt.Printf("Index: %d  Value: %d\n", i, v)
	}
	*/

	arr := [...]string {"a", "b", "c", "d", "e", "f"}
	s1 := arr[1:3]
	s2 := arr[2:4]

	fmt.Printf("s1%s \n",s1)
	fmt.Printf("s2%s \n", s2)

	s1[2] = "z"
	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)
	fmt.Printf("%s\n", arr)


}