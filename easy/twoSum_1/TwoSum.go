package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, num := range nums {
		res := target - num
		fmt.Printf("Итерация %d: Число = %d, Ищем = %d\n", i, num, res)

		if index, ok := hashTable[res]; ok {
			fmt.Printf("Найдено! Индексы: %d и %d\n", index, i)
			return []int{index, i}
		}

		fmt.Printf("Добавляем в хэш-таблицу: %d -> %d\n", num, i)
		hashTable[num] = i
	}
	return []int{}
}
