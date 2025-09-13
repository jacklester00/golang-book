package main

import "fmt"

type MultiShape struct {
	shapes []Shape
}

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.area()
	}
	return area
}

func (c Circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (r Rectangle) area() float64 {
	return (r.x2 - r.x1) * (r.y2 - r.y1)
}

func (m MultiShape) area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 1}
	fmt.Println(c.area())

	r := Rectangle{0, 0, 1, 1}
	fmt.Println(r.area())

	fmt.Println(totalArea(c, r))

	ms := MultiShape{shapes: []Shape{c, r}}
	fmt.Println(ms.area())
}
