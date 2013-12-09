package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problems. and only %d solutions\n",
        math.Nextafter(2, 3), 2)
}