package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite o número: ")
	fmt.Scan(&num)

	var antecessor int
	antecessor = num - 1

	var sucessor int
	sucessor = num + 1

	fmt.Println("O antecessor do número digitado é ", antecessor, "e o sucessor é", sucessor)
}
