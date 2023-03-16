package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	var i int
	i = 1
	for i <= 10 {
		var mul int
		mul = i * num
		fmt.Println(num, "x", i, "=", mul)
		i++
	}
}
