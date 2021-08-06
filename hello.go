package main
import "fmt"
func add1(a *int) int {
	*a = *a+1
	return *a
}
type person struct {
	name string
	age int
}
func older(p1, p2 person)(person, int) {
	if p1.age>p2.age {
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}
func main() {
	var tom person
	tom.name, tom.age = "Tom", 18
	bob := person{age:25,name:"Bob"}
	//paul := person{"Paul",43}
	tb_older,tb_diff := older(tom,bob)
	fmt.Println(tb_older,tb_diff)
}