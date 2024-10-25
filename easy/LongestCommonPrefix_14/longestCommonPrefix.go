package main

import "fmt"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	fmt.Println("Initial prefix:", prefix)

	for i := 1; i < len(strs); i++ {
		fmt.Println("\nComparing with:", strs[i])
		for len(prefix) > len(strs[i]) || strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			fmt.Println("Updated prefix:", prefix)
			if len(prefix) == 0 {
				return ""
			}
		}
	}

	fmt.Println("\nFinal prefix:", prefix)
	return prefix
}
