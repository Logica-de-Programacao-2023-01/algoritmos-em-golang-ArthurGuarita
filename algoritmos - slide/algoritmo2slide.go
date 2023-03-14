package main

import "fmt"

func main() {
	var base float64
	fmt.Print("Qual a base do retângulo? ")
	fmt.Scan(&base)
	var altura float64
	fmt.Print("Qual a altura do retângulo? ")
	fmt.Scan(&altura)
	var area float64
	area = base * altura
	fmt.Print("A área do retângulo é: ", area)
}
