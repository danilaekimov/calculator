package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Калькулятор")

	// Получение аргументов из командной строки
	if len(os.Args) != 4 {
		fmt.Println("Использование: calculator число1 операция число2")
		os.Exit(1)
	}

	// Преобразование строки в число
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Ошибка: первое число должно быть числом")
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Ошибка: второе число должно быть числом")
		os.Exit(1)
	}

	// Получение опреации
	operation := os.Args[2]

	// Выполнение вычислений
	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Ошибка: неизвестная операция")
		os.Exit(1)
	}
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operation, num2, result)
}
