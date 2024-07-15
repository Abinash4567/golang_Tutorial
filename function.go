package main

import "fmt"

type car struct {
	name  string
	model int
	tyre  Tyre
}

type Tyre struct {
	radius int
	width  int
}

func main() {
	x, _ := getPoints()
	y := getYPoint(x)
	fmt.Printf("The x-coordinate and Y-Coordinate are (%d, %d).\n", x, y)

	z, _ := getZPoint(x, y)
	fmt.Printf("The Z-Coordinate is %d.\n", z)
}

func getPoints() (int, int) {
	return 2, 3
}

func getYPoint(x int) int {
	return x * 2
}

func getZPoint(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("invalid coordinates: x=%d, y=%d", x, y)
	} else {
		return x * y, nil
	}
}
