package main

// Apenas inteiros, sem float
func mySqrt(x uint) uint {
	y := x
	y = (y + x/y) / 2
	return y
}

// Não funciona com x = 8, resolver depois
