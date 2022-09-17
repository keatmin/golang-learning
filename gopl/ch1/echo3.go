package main

import "fmt"
import "os"
import "strings"

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// Prints everything in slice
	fmt.Println(os.Args[1:])
}
