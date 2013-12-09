package main

import (
    "fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	z1 := 0.0
	i := 0
	for  math.Abs(z - z1) > .00001 {
		z1 = z
		z = z - (math.Pow(z,2)-x)/(2*z)
		i++
	} 
	return z, i
}

func main() {
	
	fmt.Println(Sqrt(2))
}