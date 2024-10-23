package main

import "fmt"

func RomanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanMap[s[i]]
		fmt.Printf("Обрабатываем символ: %c, его значение: %d\n", s[i], currentValue)

		if currentValue < prevValue {
			total -= currentValue
			fmt.Printf("Вычитаем: %d, текущий total: %d\n", currentValue, total)
		} else {
			total += currentValue
			fmt.Printf("Добавляем: %d, текущий total: %d\n", currentValue, total)
		}

		prevValue = currentValue
		fmt.Printf("Предыдущее значение обновлено на: %d\n\n", prevValue)
	}

	return total
}
