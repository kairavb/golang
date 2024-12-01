package main

import (
	"errors"
	"fmt"
)

func main() {
	var val string = "yo yo yo"
	printMe(val)

	var numr int = 11
	var denr int = 0
	var result, remainder, err = intDivision(numr, denr)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("division result: %v\n", result)
	} else {
		fmt.Printf("division result: %v\nremanider result: %v\n", result, remainder)
	}
}

func printMe(txt string) {
	fmt.Println(txt)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
