package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	var Imparpar int
	Imparpar = num % 2
	if Imparpar == 0 {
		fmt.Println("O número digitado é par.")
	} else {
		fmt.Println("O número digitado é ímpar.")
	}
}
