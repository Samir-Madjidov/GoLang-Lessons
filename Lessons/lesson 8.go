package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
	Number  string
}

// Метод для пополнения счета
func (acc *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		acc.Balance += amount
		fmt.Printf("Счет пополнен на %.2f. Новый баланс: %.2f\n", amount, acc.Balance)
	} else {
		fmt.Println("Ошибка: сумма должна быть положительной")
	}
}

// Метод для снятия денег
func (acc *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= acc.Balance {
		acc.Balance -= amount
		fmt.Printf("Со счета снято %.2f. Новый баланс: %.2f\n", amount, acc.Balance)
		return true
	} else {
		fmt.Println("Ошибка: недостаточно средств или неверная сумма")
		return false
	}
}

// Метод для перевода денег на другой счет
func (acc *BankAccount) Transfer(amount float64, recipient *BankAccount) bool {
	if acc.Withdraw(amount) {
		recipient.Deposit(amount)
		fmt.Printf("Перевод %.2f на счет %s выполнен успешно\n", amount, recipient.Owner)
		return true
	}
	return false
}

// Метод для вывода информации о счете
func (acc BankAccount) Display() {
	fmt.Printf("Владелец: %s\n", acc.Owner)
	fmt.Printf("Номер счета: %s\n", acc.Number)
	fmt.Printf("Баланс: %.2f\n", acc.Balance)
}

func main() {
	// Создаем счета
	account1 := BankAccount{
		Owner:   "Анна",
		Balance: 1000,
		Number:  "1234567890",
	}

	account2 := BankAccount{
		Owner:   "Петр",
		Balance: 500,
		Number:  "0987654321",
	}

	fmt.Println("Начальное состояние:")
	account1.Display()
	fmt.Println()
	account2.Display()

	fmt.Println("\n--- Операции ---")
	account1.Deposit(200)
	account1.Withdraw(150)
	account1.Transfer(300, &account2)

	fmt.Println("\nФинальное состояние:")
	account1.Display()
	fmt.Println()
	account2.Display()
}
