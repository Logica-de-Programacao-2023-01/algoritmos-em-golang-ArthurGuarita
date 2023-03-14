package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Digite o valor do salário: R$")
	fmt.Scan(&salario)
	var salarioNovo float64
	salarioNovo = salario * 1.15
	fmt.Print("O salário novo é de R$", salarioNovo)
}
