package main

import "fmt"

func calculate(a int, b int, operation string) int {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b

	}
	return 0

}

func main() {
	var a, b int
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&op)

	result := calculate(a, b, op)
	fmt.Println("Результат:", result)

}
