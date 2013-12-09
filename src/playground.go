package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    fmt.Println("Welcome to the playground!")

    fmt.Println("The time is", time.Now())

    fmt.Println("And if you try to open a file:")
	
	_, err := os.Open("/Users/alex/alex/web.go") // For read access.
	if err != nil {
		fmt.Println(err)
	} else {
		
		fmt.Println("Found it!")
	}
		
    fmt.Println("Or access the network:")
    fmt.Println(net.Dial("tcp", "google.com"))
}