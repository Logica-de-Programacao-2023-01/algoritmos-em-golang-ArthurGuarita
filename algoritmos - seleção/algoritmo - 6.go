package main

import "fmt"

func main() {
	var a int
	fmt.Print("Digite o valor de a: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Digite o valor de b: ")
	fmt.Scan(&b)

	if a > 0 && b > 0 {
		var multi int
		multi = a * b
		fmt.Println("A multiplicação entre eles é", multi)
	} else {
		var soma int
		soma = a + b
		fmt.Println("A soma entre eles é", soma)
	}
}