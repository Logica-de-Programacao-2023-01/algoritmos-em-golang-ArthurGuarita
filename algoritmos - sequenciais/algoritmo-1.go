package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)
	var soma = num1 + num2 + num3
	fmt.Print("A soma dos três números é: ", soma)
}
