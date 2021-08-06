package main
import (
	"fmt"
	"math"
)
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}
func (r Rectangle) area() float64 {
	return r.width*r.height
}
func (c Circle) area() float64 {
	return c.radius*c.radius*math.Pi
}
func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println(r1.area(),r2.area(),c1.area(),c2.area())
}