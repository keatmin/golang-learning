package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
	optional()
	for_while()
	if_statements()
	fmt.Println(
		if_with_a_short_statement(2, 3, 10),
		if_with_a_short_statement(3, 4, 10),
	)
	go_switch()
	switch_exec_order()
	defer_stuff()
	defer_stacking()
}

func optional() {
	// init and post is optional
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func for_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func if_statements() {
	// Need no parentheses but need braces
	x := 20
	if x >= 20 {
		fmt.Println("x > 20")
	}
}

func if_with_a_short_statement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		// Variables declared in if will be available in else
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func go_switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch_exec_order() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// switch can be used with no condition, removing time.Saturday below but case has to be a boolean
	switch time.Saturday {
	//if time.Saturday = today + 0
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func defer_stuff() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func defer_stacking() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// defer is added into a stack and follows a LIFO
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
