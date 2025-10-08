package main

import "fmt"

func main() {
	var massa float64
	var rost float64
	fmt.Println("Введите массу тела в кг:")
	fmt.Scan(&massa)
	fmt.Println("Введите рост в см:")
	fmt.Scan(&rost)
	rost = rost / 100
	ИМТ := massa / (rost * rost)
	fmt.Printf("Ваш ИМТ: %.2f\n", ИМТ)
	if ИМТ < 18.5 {
		fmt.Println("Недостаточная масса тела")
	} else if ИМТ < 24.9 {
		fmt.Println("Нормальная масса тела")
	} else {
		fmt.Println("Избыточная масса тела")
	}
}



// by Samir Madjidov 
// github.com/samir-madjidov
