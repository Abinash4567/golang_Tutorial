package main

import "fmt"

type gasEngine struct {
	mpg     int
	gallons int
}

type electricEngine struct {
	mpkwh int
	kwh   int
}

type vehicle interface {
	milesLeft() int
}

func (e gasEngine) milesLeft() int {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() int {
	return e.kwh * e.mpkwh
}

func canMakeIt(e vehicle, miles int) {
	if miles <= e.milesLeft() {
		fmt.Println("Reached Home")
	} else {
		fmt.Println("Cannot Reach Home")
	}
}

func main() {
	var nano gasEngine = gasEngine{mpg: 12, gallons: 2}
	var tesla electricEngine = electricEngine{mpkwh: 12, kwh: 23}

	fmt.Println(nano.milesLeft())
	fmt.Println(tesla.milesLeft())

	canMakeIt(nano, 34)
}
