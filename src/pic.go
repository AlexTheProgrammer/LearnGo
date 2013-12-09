package main

import (
    "code.google.com/p/go-tour/pic"
    )

func Pic(dx, dy int) [][]uint8 {
    a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
        for w := range a[i] {
			a[i][w] = uint8((i * w)/2)
		}
	}
    return a
}

func main() {
    pic.Show(Pic)
}