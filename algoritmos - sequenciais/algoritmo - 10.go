package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Digite a sua massa corporal em kg: ")
	fmt.Scan(&peso)

	var conversao float64
	conversao = peso * 2.205
	fmt.Println("A conversão da massa corporal de kg para libras é de ", conversao)
}
