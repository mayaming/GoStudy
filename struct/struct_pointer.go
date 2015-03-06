package main

import (
	"fmt"
)

type HasArea interface {
	area() int32
}

type Rectangle struct {
	width, height int32
}

func (r Rectangle) double() {
	r.width *= 2
	r.height *= 2
}

func (r *Rectangle) triple() {
	r.width *= 3
	r.height *= 3
}

func (r *Rectangle) area() int32 {
	return r.width * r.height
}

func pointer_example() {
	fmt.Println("pointer_example...")
	rec := Rectangle{10, 20}
	rec.double()
	//10 20
	fmt.Println(rec.width, " ", rec.height)

	recp := &rec
	recp.double()
	//10 20
	fmt.Println(rec.width, " ", rec.height)

	rec.triple()
	//30 60
	fmt.Println(rec.width, " ", rec.height)

	//1800
	fmt.Println(rec.area())

	/*This can't work
	var ar HasArea = rec
	//1800
	fmt.Println(ar.area())
	*/

}
