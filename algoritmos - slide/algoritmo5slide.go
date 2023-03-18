package main

import "fmt"

// Programa feito às 15:33 do dia 14/03/2023 — Caso o valor do dólar esteja diferente, altere.

func main() {
	var money float64
	fmt.Print("Digite o valor do dinheiro a ser convertido: ")
	fmt.Scan(&money)

	var tipomoney int
	fmt.Print("Digite 1 para converter dólar para real, ou digite 2 para converter real para dólar: ")
	fmt.Scan(&tipomoney)

	if tipomoney == 1 {
		var valorReal float64
		valorReal = money * 5.24
		fmt.Print("O valor convertido é: ", valorReal)
	}
	if tipomoney == 2 {
		var valorDolar float64
		valorDolar = money / 5.24
		fmt.Print("O valor convertido é: ", valorDolar)
	}
}
