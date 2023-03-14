package main

import "fmt"

func main() {
	var nome string
	fmt.Print("What's your name? ")
	fmt.Scan(&nome)
	var idade int
	fmt.Print("What's your age? ")
	fmt.Scan(&idade)
	fmt.Print("Whats your weight? ")
	var peso float64
	fmt.Scan(&peso)
	fmt.Println("Your name is", nome)
	fmt.Println("Your age is", idade)
	fmt.Println("Your weight is", peso)
}
