package main

import (
    "fmt"
    "math/rand"
)

func main() {
	rand.Seed(23987623784)
    fmt.Println("My favorite number is", rand.Intn(10))
}