package main

import (
	"fmt"
)

type Returnable interface {
	returnWith()
}

type ReturnString struct {
}

func (rs ReturnString) returnWith() string {
	return "hello, world"
}

type Integer uint32

func (i Integer) returnWith() Integer {
	return i
}

type ReturnNothing struct {
}

func (rn ReturnNothing) returnWith() {
	fmt.Println("Return nothing...")
}

func interface_return() {
	fmt.Println("interface_return...")
	/*Won't work because return type does not match
	rs := ReturnString{}
	n := Integer(10)

	retable := []Returnable{rs, n}
	for _, ret := range retable {
		fmt.Println(ret.returnWith())
	}
	*/

	rn := ReturnNothing{}
	rn.returnWith()
}
