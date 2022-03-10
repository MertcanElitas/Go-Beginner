package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct {
	a float64
	h float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

type reactangle struct {
	a, b float64
}

func (r reactangle) area() float64 {
	return (r.a * r.b)
}

func printArea(s ...shape) {
	for _, shape := range s {
		fmt.Println("Alan", shape.area())
	}
}

func main() {
	r := reactangle{3, 4}
	s := square{5}
	t := triangle{4, 5}

	printArea(r, s, t)
}
