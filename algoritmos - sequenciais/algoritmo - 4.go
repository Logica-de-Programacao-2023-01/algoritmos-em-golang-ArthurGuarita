package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var num3 float64
	fmt.Print("Digite o primeiro número com peso 2: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número com peso 3: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número com peso 5: ")
	fmt.Scan(&num3)
	var media float64
	media = (num1*2 + num2*3 + num3*5) / 10
	fmt.Println("A média ponderada desses números é: ", media)
}
