// Вычислить остаток от деления двух чисел
// Выполнил: Samir Madjidov
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
}
