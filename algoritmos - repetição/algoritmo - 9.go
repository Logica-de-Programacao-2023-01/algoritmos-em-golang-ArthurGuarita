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
	if len(nums) == 0 {
		fmt.Print("Nenhum número foi digitado, tente novamente.")
	} else {
		var soma int
		for _, nums := range nums {
			soma = soma + nums
		}
		var media int
		media = (soma) / (len(nums))
		fmt.Print("Média é: ", media)
	}
}
