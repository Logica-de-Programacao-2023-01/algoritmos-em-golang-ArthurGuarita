package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	var dobro, triplo, quadruplo int
	dobro = num * 2
	triplo = num * 3
	quadruplo = num * 4

	fmt.Println("O dobro do número digitado é: ", dobro)
	fmt.Println("O triplo do número digitado é: ", triplo)
	fmt.Println("O quadruplo do número digitado é: ", quadruplo)
}
