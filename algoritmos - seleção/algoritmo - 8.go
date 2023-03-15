package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número entre 1 e 7: ")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Println("Domingo")
	}
	if num == 2 {
		fmt.Println("Segunda")
	}
	if num == 3 {
		fmt.Println("Terça")
	}
	if num == 4 {
		fmt.Println("Quarta")
	}
	if num == 5 {
		fmt.Println("Quinta")
	}
	if num == 6 {
		fmt.Println("Sexta")
	}
	if num == 7 {
		fmt.Println("Sábado")
	}
}
