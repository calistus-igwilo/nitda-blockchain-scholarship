package main

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi %v Welcome!\n", name)

	return message
}

func main() {
	fmt.Printf(Hello("Calistus"))
}