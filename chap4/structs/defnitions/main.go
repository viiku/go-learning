package main

import "fmt"

type Point struct {
	X, Y int
}

func scaleUsingPointers(p *Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func addition(p Point) int {
	val := p.X + p.Y
	return val
}

// func Scale(p Point, factor int) Point {
// 	return Point{p.X * factor, p.Y * factor}
// }

func main() {
	p := Point{1, 2}
	fmt.Printf("%d %d\n", p.X, p.Y)

	// Multiplication
	// val := Scale(Point{1, 2}, 5)
	// fmt.Printf("%d\n", val)
	// adding
	// fmt.Printf("%d\n", addition(val))

	// passing structs in function using pointers
	// this is equivalent to below assign
	// pp := new(Point)
	// *pp = Point{1, 2}
	pp := &Point{1, 2}
	val := scaleUsingPointers(pp, 5)
	fmt.Printf("%d\n", val)
}
