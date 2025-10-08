package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число:")
	fmt.Scan(&number)
	if number%3 == 0 {
		fmt.Println(number, "Fizz")
	} else if number%5 == 0 {
		fmt.Println(number, "Buzz")
	} else {
		fmt.Println(number, "FizzBuzz")
	}
}

// by Samir Madjidov
// github.com/samir-madjidov
