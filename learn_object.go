package main
import (
	"fmt"
	//"math"
)
/*
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
*/
type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n",h.name,h.phone)
}
func main() {
	/*
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println(r1.area(),r2.area(),c1.area(),c2.area())
	*/
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}