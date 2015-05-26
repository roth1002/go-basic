package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android_Deyu struct {
	Person Person
	Model  string
}

type Android_Tina struct {
	Person
	Model string
}

type Shape interface {
	area() float64
}

func main() {
	//Circla
	var a Circle     // create circle struct instance with x,y,z = 0
	b := new(Circle) //new
	c := Circle{x: 0, y: 0, r: 5}
	d := Circle{0, 0, 5}

	fmt.Println(a.x, a.y, a.r)
	fmt.Println(b.x, b.y, b.r)
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(d.x, d.y, d.r)

	fmt.Println(circleArea(c))
	fmt.Println(circleAreaByRef(&c))
	fmt.Println(c.area())

	//Rectangle
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	//Embedded Types
	//Origin
	e := new(Android_Deyu)
	e.Person.Name = "Leo"
	e.Person.Talk()
	//Embedded
	f := new(Android_Tina)
	f.Name = "Tim"
	f.Talk()

	//use interface
	fmt.Println(totalArea(&c, &r))
	//g := MultiShape
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaByRef(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}
