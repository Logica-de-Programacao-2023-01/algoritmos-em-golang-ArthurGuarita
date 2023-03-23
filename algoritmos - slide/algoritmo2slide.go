package main

import "fmt"

func main() {
	var base, altura, area float64
	fmt.Print("Qual a base do retângulo? ")
	_, err := fmt.Scanln(&base)

	if err != nil {
		// Trata o erro de entrada inválida
		fmt.Println("Erro: Entrada inválida!")
		return // Adicionado um retorno para evitar continuar com a execução em caso de erro
	}
	fmt.Print("Qual a altura do retângulo? ")
	_, err = fmt.Scanln(&altura)

	if err != nil {
		// Trata o erro de entrada inválida
		fmt.Println("Erro: Entrada inválida!")
		return // Adicionado um retorno para evitar continuar com a execução em caso de erro
	}

	area = base * altura
	fmt.Print("A área do retângulo é: ", area)
}
