package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite outro número: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("O número", num1, "é maior que ", num2)

	} else {
		fmt.Println("O número", num2, "é maior que ", num1)
	}

}
