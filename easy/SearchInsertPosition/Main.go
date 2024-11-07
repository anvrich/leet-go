package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	result := SearchInsert(nums, target)
	fmt.Println(result)
}
