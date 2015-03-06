package main

/*
import (
	"fmt"
)
*/

type Increaser interface {
	inc(int32) int32
}

type Decreaser interface {
	dec(int32) int32
}

type OneIncDec struct {
	n int32
}

func (sid OneIncDec) inc(i int32) int32 {
	sid.n += 1
	return i + 1
}

func (sid OneIncDec) dec(i int32) int32 {
	sid.n -= 1
	return i - 1
}

type TwoIncDec struct {
	n int32
}

func (sid TwoIncDec) inc(i int32) int32 {
	sid.n += 2
	return i + 2
}

func (sid *TwoIncDec) dec(i int32) int32 {
	sid.n -= 2
	return i - 2
}

type Compositer1 struct {
	*OneIncDec
	twoIncDec *TwoIncDec
}

/*
func main() {
	oneID := OneIncDec{10}
	twoID := TwoIncDec{20}
	comp1 := Compositer1{&oneID, &twoID}
	//6
	fmt.Println(comp1.inc(5))
	//3
	fmt.Println(comp1.twoIncDec.dec(5))
	//11
	fmt.Println(oneID.n)
	//18
	fmt.Println(twoID.n)
}
*/
