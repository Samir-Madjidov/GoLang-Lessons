package main

import "fmt"

func main() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("Ты несовершеннолетний.")
	} else {
		fmt.Println("Ты взрослый.")
	}

}
