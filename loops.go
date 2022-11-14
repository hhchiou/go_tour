package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	limit := 0.000000000000001
	for i := 0; i < 10; i++ {
		y := (z*z - x) / (2 * z)
		if (y < 0 && y > -limit) || (y >= 0 && y < limit) {
			break
		}
		z -= y
		fmt.Println(i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
