package main

import "fmt"

func main() {
	var idade uint
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	if idade < 9 {
		fmt.Print("Mirim")
	}
	if idade >= 10 && idade <= 13 {
		fmt.Print("Infantil")
	}
	if idade >= 14 && idade <= 17 {
		fmt.Print("Juvenil")
	}
	if idade >= 18 {
		fmt.Print("Adulto")
	}
}
