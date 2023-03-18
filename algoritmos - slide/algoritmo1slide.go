package main

import "fmt"

func main() {
	var nome string
	fmt.Print("Qual o seu nome? ")
	fmt.Scan(&nome)

	var idade int
	fmt.Print("Quantos anos você têm? ")
	fmt.Scan(&idade)

	fmt.Print("Qual sua massa corporal? ")
	var peso float64
	fmt.Scan(&peso)

	fmt.Println("Olá, seu nome é", nome, ", você têm", idade, "anos de idade e sua massa corporal é", peso, "kg.")
}
