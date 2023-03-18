package main

import "fmt"

func main() {
	var num1, num2, num3, num4 float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)
	fmt.Print("Digite o quarto número: ")
	fmt.Scan(&num4)

	var media float64
	media = (num1 + num2 + num3 + num4) / 4
	fmt.Print("A média aritmética dos números digitados é: ", media)
}
