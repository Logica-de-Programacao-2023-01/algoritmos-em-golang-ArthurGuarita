package main

import "fmt"

func main() {
	var valorProduto float64
	fmt.Print("Digite o valor do produto: ")
	fmt.Scan(&valorProduto)

	var desconto float64
	desconto = valorProduto * 0.9
	fmt.Print("O valor com desconto de 10% é de R$", desconto)
}
