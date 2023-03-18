package main

import "fmt"

func main() {
	var peso, altura, imc float64

	fmt.Print("Digite a sua massa corporal em kg: ")
	fmt.Scan(&peso)

	fmt.Print("Digite sua altura em metros: ")
	fmt.Scan(&altura)

	imc = peso / (altura * altura)
	fmt.Print("O seu IMC é ", (fmt.Sprintf("%.2f", imc)))
}
