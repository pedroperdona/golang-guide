package main

import "fmt"

func main() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for number := range numbers {
		fmt.Println(number, "is "+evenOrOdd(number))
	}
}

func evenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
}
