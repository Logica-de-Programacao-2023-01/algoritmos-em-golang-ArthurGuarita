package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Digite o seu salário: ")
	fmt.Scan(&salario)

	var novoSalario float64

	if salario < 1000 {
		novoSalario = salario * 1.1
		fmt.Println("O novo salário é de", novoSalario)
	} else {
		novoSalario = salario * 1.05
		fmt.Println("O novo salário é de", novoSalario)
	}
}
