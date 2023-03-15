package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	if ((num % 3) == 0) && ((num % 5) == 0) {
		fmt.Println("O número", num, "é múltiplo de 3 e 5.")
	} else {
		fmt.Println("O número 'não' é múltiplo de 3 e 5")
	}
}
