package main

import "fmt"

func main() {
	var diaria int
	var dias int
	fmt.Print("Digite o valor da diária: ")
	fmt.Scan(&diaria)
	fmt.Print("Digite por quantos dias você trabalhou: ")
	fmt.Scan(&dias)
	var salario int
	salario = diaria * dias
	fmt.Println("O salário é de R$", salario)
}
