package main

import "fmt"

func main() {
	var celsius float64
	fmt.Println("Введите температуру в градусах Цельсия:")
	fmt.Scan(&celsius)
	fahrenheit := celsius*9/5 + 32
	fmt.Println("Температура в градусах Фаренгейта:", fahrenheit)
}
