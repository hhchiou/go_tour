package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	z := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		z[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			//z[i][j] = uint8((j + i) / 2)
			//z[i][j] = uint8(j * i)
			z[i][j] = uint8(j ^ i)
		}
	}
	return z
}

func main() {
	pic.Show(Pic)
}
