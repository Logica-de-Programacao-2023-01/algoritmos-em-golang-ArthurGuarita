package main

import "fmt"

func main() {
	var base float64
	fmt.Print("Qual a base da caixa? ")
	fmt.Scan(&base)
	var altura float64
	fmt.Print("Qual a altura da caixa? ")
	fmt.Scan(&altura)
	var profundidade float64
	fmt.Print("Qual a profundidade da caixa? ")
	fmt.Scan(&profundidade)
	var volume float64
	volume = base * altura * profundidade
	fmt.Print("O volume da caixa é ", volume)
}
