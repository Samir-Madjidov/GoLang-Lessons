package main

import "fmt"

// Функция возвращает результат и остаток от деления
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func main() {
	q, r := divide(10, 3)
	fmt.Printf("10 / 3 = %d (остаток %d)\n", q, r) // 10 / 3 = 3 (остаток 1)
}
