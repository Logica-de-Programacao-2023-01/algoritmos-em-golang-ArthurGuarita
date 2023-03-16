package main

import "fmt"

func main() {
	var num uint
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	var i uint
	i = 1
	for i <= num {
		var div uint
		div = num % i
		if div == 0 {
			fmt.Println(i)
		}
		i++
	}
}
