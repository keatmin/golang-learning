package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// It is a good style to use the factored import statement

//type comes last
var c, python, java bool

// Can define with initializer

var i, j int = 1, 2

// Constants are declared like variables, but cannot be delcared using :=

const Hi = 31.2

func main() {
	fmt.Println("Welcome")
	fmt.Println("The time is", time.Now())
	fmt.Println("The rand number is", rand.Intn(10))
	exportedNames()
	fmt.Println(add(2, 3))
	x, y := swap("world", "hello")
	fmt.Println(x, y)
	fmt.Println(split(5))

	fmt.Println(i, j, c, python, java)

}

func exportedNames() {
	fmt.Println("Name is exported if it begins with a capital letter math.pi will not run", math.Pi)
}

func add(x int, y int) int {
	return x + y
}

func add_shorterned(x, y int) int {
	// This works too
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Doing a naked return, named return

func split(sum int) (x, y int) {
	// Only should be used in short function
	x = sum * 4
	y = sum - 4
	return
}

func var_assignment() {
	// Inside a function, := assignment can be used in place of a var
	// Outside a function, every statement begins with a keyword(var, func etc) and := not available
}

// Variables declared without initial input are given zero

var integer int  //0
var str string   // ""
var boolean bool // False
