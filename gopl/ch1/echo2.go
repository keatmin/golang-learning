package main

import "fmt"
import "os"

func main() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}
