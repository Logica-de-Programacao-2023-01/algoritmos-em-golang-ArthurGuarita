package main

import "fmt"

func main() {
	var i int
	i = 1
	for i <= 100 {
		var div3 int
		var div5 int
		div3 = i % 3
		div5 = i % 5
		if div3 != 0 && div5 != 0 {
			fmt.Println(i)
		}
		if div3 == 0 && div5 == 0 {
			fmt.Println(i, "= FizzBuzz.")
		} else {
			if div3 == 0 {
				fmt.Println(i, "= Fizz.")
			}
			if div5 == 0 {
				fmt.Println(i, "= Buzz.")
			}
		}
		i++
	}
}
