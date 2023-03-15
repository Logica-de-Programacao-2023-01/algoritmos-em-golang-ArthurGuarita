package main

import "fmt"
import "sort"

func main() {
	var num1, num2, num3 float64
	fmt.Print("Digite três números: ")
	fmt.Scan(&num1, &num2, &num3)

	numeros := []float64{num1, num2, num3}
	sort.Float64s(numeros)

	fmt.Println("Os números em ordem crescente são:", numeros)
}
