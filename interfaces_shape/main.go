package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLenght float64
}

func main() {
	triangle := triangle{base: 10, height: 5}
	square := square{sideLenght: 10}

	printArea(triangle)
	printArea(square)
}

func (triangle triangle) getArea() float64 {
	return 0.5 * triangle.base * triangle.height
}

func (square square) getArea() float64 {
	return square.sideLenght * square.sideLenght
}

func printArea(shape shape) {
	fmt.Println(shape.getArea())
}
