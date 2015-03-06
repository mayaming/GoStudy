package main

import (
	"fmt"
	"reflect"
)

type NegInteger struct {
	n int32
}

func abs(n int32) int32 {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

func (ni NegInteger) abs() int32 {
	if ni.n >= 0 {
		return ni.n
	} else {
		return -ni.n
	}
}

func main() {
	ni1 := NegInteger{-10}
	ni2 := NegInteger{-20}
	fmt.Println(abs(-10))

	//The same output for the ni1.abs and ni2.abs because they
	//point to the same function
	fmt.Println(ni1.abs)
	fmt.Println(ni2.abs)
	fmt.Println(NegInteger.abs)

	fmt.Println(ni1.abs())
	fmt.Println(ni2.abs())

	fp1 := abs
	fmt.Println(fp1)
	fmt.Println(reflect.TypeOf(fp1))
	fp2 := NegInteger.abs
	fmt.Println(fp2)
	fmt.Println(reflect.TypeOf(fp2))
}
