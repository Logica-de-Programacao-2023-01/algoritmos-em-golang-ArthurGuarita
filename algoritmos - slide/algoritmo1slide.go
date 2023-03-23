package main

import "fmt"

func main() {
	var nome string
	fmt.Print("Qual o seu nome? ")
	_, err := fmt.Scanln(&nome)

	if err != nil {
		// Trata o erro de entrada inválida
		fmt.Println("Erro: Entrada inválida!")
		return // Adicionado um retorno para evitar continuar com a execução em caso de erro
	}

	var idade int
	fmt.Print("Quantos anos você têm? ")
	_, err = fmt.Scanln(&idade) // Corrigido para apenas uma atribuição

	if err != nil {
		// Trata o erro de entrada inválida
		fmt.Println("Erro: Entrada inválida!")
		return // Adicionado um retorno para evitar continuar com a execução em caso de erro
	}

	fmt.Print("Qual sua massa corporal? ")
	var peso float64
	_, err = fmt.Scanln(&peso)

	if err != nil {
		// Trata o erro de entrada inválida
		fmt.Println("Erro: Entrada inválida!")
		return // Adicionado um retorno para evitar continuar com a execução em caso de erro
	}

	fmt.Printf("Olá, seu nome é %s, você possui %d anos de idade e %.2f kg de massa corporal", nome, idade, peso)
}
