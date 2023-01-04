package main

import "fmt"

func factorialNumber(anyNumber float64) float64 {
	if anyNumber < 0 {
		fmt.Println("The factorial of a negative number is undefined")
		return 0
	} else if anyNumber == 0 {
		return 1
	} else {
		return anyNumber * factorialNumber(anyNumber-1)
	}
}

func greatest(first int, second int, third int, fourth int) int {

	if first > second && first > third && first > fourth {
		return first
	} else if second > third && second > fourth {
		return second
	} else if third > fourth {
		return third
	} else {
		return fourth
	}
}

func truncate(numberFloat float64) (numberInt int) {
	numberInt = int(numberFloat)
	return
}

func wordNumber(word string, wordNum int) string {
	if wordNum <= 0 {
		return ("Please Enter a number above zero")
	}
	result := ""
	for i := 0; i < wordNum; i++ {
		result += word
	}
	return result

}

func main() {
	var fname string
	var chosenNumber int
	fmt.Println("Hi, Welcome! Whats your name?")
	fmt.Scanf("%s", &fname)
	fmt.Printf("Hi %s, Welcome\n", fname)
	fmt.Println("1. Factorial")
	fmt.Println("2. Find the greatest of three numbers")
	fmt.Println("3. Truncate a decimal")
	fmt.Println("4. Print HI")
	fmt.Println("Choose the number from the list above")
	fmt.Scanf("%d", &chosenNumber)

	if chosenNumber == 1 {
		var anyNumber float64
		fmt.Println("Please enter your number here")
		fmt.Scanf("%f", anyNumber)
		result := factorialNumber(anyNumber)
		fmt.Printf("The factorial of %f is %f", anyNumber, result)
	} else if chosenNumber == 2 {
		var first int
		var second int
		var third int
		var fourth int
		fmt.Println("Please enter your first number")
		fmt.Scanf("%d", &first)
		fmt.Println("Please enter your second number")
		fmt.Scanf("%d", &second)
		fmt.Println("Please enter your third Number")
		fmt.Scanf("%d", &third)
		fmt.Println("Please enter your fourth number")
		fmt.Scanf("%d", &fourth)
		result := greatest(first, second, third, fourth)
		fmt.Printf("The greatest of the four numbers is %d", result)
	} else if chosenNumber == 3 {
		var numberFloat float64
		fmt.Println("Please input your decimal number")
		fmt.Scanf("%f", &numberFloat)
		result := truncate(numberFloat)
		fmt.Printf("The truncated value of %f is %d", numberFloat, result)
	} else if chosenNumber == 4 {
		var word string
		var wordNum int
		fmt.Println("Enter the word you wish to print")
		fmt.Scanf("%s", &word)
		fmt.Println("Enter the number of times to be printed")
		fmt.Scanf("%d", &wordNum)
		fmt.Println(wordNumber(word, wordNum))
	} else {
		fmt.Println("Invalid choice. Try again")
	}
}
