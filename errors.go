package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d\n", int(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	limit := 0.000000000000001
	for i := 0; i < 10; i++ {
		y := (z*z - x) / (2 * z)
		if (y < 0 && y > -limit) || (y >= 0 && y < limit) {
			break
		}
		z -= y
		//fmt.Println(i, z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
