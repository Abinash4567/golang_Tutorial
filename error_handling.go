package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 12
	var denominator int = 0
	var quotient, remainder, error = functioncall(numerator, denominator)
	if error == nil {
		fmt.Print(quotient, remainder)
	} else {
		fmt.Println("Error occured!!!")
	}
}

func functioncall(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot be divided by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}
