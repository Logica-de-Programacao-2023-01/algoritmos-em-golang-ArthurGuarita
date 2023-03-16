package main

import "fmt"

func main() {
	var i int
	i = 0
	for i <= 19 {
		var div int
		div = i % 2
		if div != 0 {
			fmt.Println(i)
		}
		i++
	}
}
