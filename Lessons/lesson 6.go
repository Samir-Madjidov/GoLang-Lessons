package main

import "fmt"

func main() {
	phonebook := make(map[string]string)
	phonebook["Jan"] = "+9924875487636"
	phonebook["Gam"] = "+9929056744553"
	phonebook["Irina"] = "+992209004765"

	var name string
	fmt.Print("Кого ещем: ")
	fmt.Scan(&name)
	if phone, exists := phonebook[name]; exists {
		fmt.Printf("Номер %s: %s\n", name, phone)
	} else {
		fmt.Printf("Контак %s не найден\n", name)
	}
}
