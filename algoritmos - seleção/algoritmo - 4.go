package main

import "fmt"

func main() {
	var altura float64
	fmt.Print("Digite a sua altura em m: ")
	fmt.Scan(&altura)

	var sexo string
	fmt.Print("Digite o seu sexo: ")
	fmt.Scan(&sexo)

	var peso float64
	fmt.Print("Digite a sua massa corporal em kg: ")
	fmt.Scan(&peso)

	var imc float64
	imc = peso / (altura * altura)
	
	if imc < 18.5 {
		fmt.Println("Abaixo do peso")
	}
	if (imc >= 18.5) && (imc <= 24.9) {
		fmt.Println("Peso normal")
	}
	if imc >= 25 {
		fmt.Println("Acima do peso")
	}
}