package main

import "fmt"

func main() {
	fmt.Println(EvenOrOdd(10))
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}