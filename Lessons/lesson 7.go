package main

import "fmt"

func main() {
	var nameArg string
	fmt.Scanln(&nameArg)
	hello(nameArg)

}

// вывод приветсвия на экран
func hello(name string) {
	fmt.Printf("Hello,", name)
}
