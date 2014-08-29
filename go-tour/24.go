package main

import (
	"fmt"
	"math"
)

const delta = 1e-9

func Sqrt(x float64) float64 {
	z, tmp := 1.0, 0.0
	for i := 0; i < 100; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		if tmp == z {
			return z
		}
		tmp = z
	}
	return z
}

func Sqrt2(x float64) float64 {
	z, tmp := 1.0, 0.0
	for math.Abs(z - tmp) > delta {
		tmp, z = z, z - (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt2(2))
}
