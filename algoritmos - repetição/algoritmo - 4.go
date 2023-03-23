package main

import "fmt"

func main() {
	var i int
	i = 0

	for i <= 30 {
		var div int

		div = i % 3
		if div == 0 {
			fmt.Println(i)
		}
		i++
	}
}
