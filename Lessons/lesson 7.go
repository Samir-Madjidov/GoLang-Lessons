package main

import "fmt"

// Обрабатывает пользовательские данные
func processUserData(name string, age int) (string, bool, error) {
	// Проверяем валидность данных
	if name == "" {
		return "", false, fmt.Errorf("имя не может быть пустым")
	}

	if age < 0 || age > 150 {
		return "", false, fmt.Errorf("некорректный возраст: %d", age)
	}

	// Формируем приветствие
	greeting := fmt.Sprintf("Привет, %s! Тебе %d лет.", name, age)

	// Проверяем совершеннолетие
	isAdult := age >= 18

	return greeting, isAdult, nil
}

func main() {
	// Тестируем функцию с разными данными
	testCases := []struct {
		name string
		age  int
	}{
		{"Анна", 25},
		{"Петр", 17},
		{"", 30},
		{"Мария", -5},
	}

	for _, tc := range testCases {
		fmt.Printf("\nТест: %s, %d лет\n", tc.name, tc.age)

		greeting, isAdult, err := processUserData(tc.name, tc.age)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println(greeting)
		if isAdult {
			fmt.Println("Совершеннолетний")
		} else {
			fmt.Println("Несовершеннолетний")
		}
	}
}
