package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Print("Digite o valor de a: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de b: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de c: ")
	fmt.Scan(&c)
	if a < b && a < c {
		fmt.Println("O valor de a(", a, ") é o menor.")
	}
	if b < a && b < c {
		fmt.Println("O valor de b(", b, ") é o menor.")
	}
	if c < a && c < b {
		fmt.Println("O valor de c(", c, ") é o menor.")
	}
}
