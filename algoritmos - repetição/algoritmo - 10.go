package main

import "fmt"

func main() {
	var nums []int
	var input int

	for {
		fmt.Print("Digite um número (digite 0 para sair): ")
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		nums = append(nums, input)
	}
	if len(nums) > 0 {
		var max int
		max = nums[0]
		for _, num := range nums {
			if num > max {
				max = num
			}
		}
		fmt.Print("O maior número é:", max)
	}
}
