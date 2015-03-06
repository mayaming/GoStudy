package main

import (
	"fmt"
)

type Gettable interface {
	get() int32
}

type ChangableP interface {
	Gettable
	changeP()
}

type ChangableV interface {
	Gettable
	changeV()
}

type IntInside struct {
	n int32
}

func (ii IntInside) get() int32 {
	return ii.n
}

func (ii *IntInside) changeP() {
	ii.n++
}

func (ii IntInside) changeV() {
	ii.n++
}

func interface_receiver() {
	fmt.Println("interface_receiver...")
	ii := IntInside{10}

	//Does not work because changeP method of IntInside class needs
	//a pointer receiver
	//cp := ChangableP(ii)
	//cp.changeP()

	//10
	fmt.Println(ii.n)

	ii.changeV()
	//10
	fmt.Println(ii.n)

	ii.changeP()
	//11
	fmt.Println(ii.n)

	cv := ChangableV(ii)
	cv.changeV()
	//11
	fmt.Println(cv.get())

	ii.changeP()
	//12
	fmt.Println(ii.get())
	//11
	fmt.Println(cv.get())

	cp := ChangableP(&ii)
	cp.changeP()
	//13
	fmt.Println(cp.get())
	//13
	fmt.Println(ii.get())
}
