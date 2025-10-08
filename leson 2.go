package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Println(" Видите первое число:")
	fmt.Scan(&a)
	fmt.Println(" Видите второе число:")
	fmt.Scan(&b)
	c = a % b
	fmt.Println("Остаток от деления чисел =", c)

	fmt.Println("Конвертер температур:")
	var celsius float64
	fmt.Println("Введите температуру в градусах Цельсия:")
	fmt.Scan(&celsius)
	fahrenheit := celsius*9/5 + 32
	fmt.Println("Температура в градусах Фаренгейта:", fahrenheit)
}
