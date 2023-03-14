package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	fmt.Print("Digite a sua massa corporal em kg: ")
	fmt.Scan(&peso)
	fmt.Println("Digite sua altura em metros: ")
	fmt.Scan(&altura)
	var imc float64
	imc = peso / (altura * altura)
	fmt.Println("O seu IMC é", imc)
}
