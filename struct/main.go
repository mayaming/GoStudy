package main

import (
	"fmt"
)

func main() {
	sim := Simple{5, "simple string"}
	comp := Complex{[3]int{1, 2, 3}, []string{"hello", "world"}, map[string]int{"foo": 1, "bar": 2}}
	composite := Composite{s: sim, c: comp}
	fmt.Println(composite)
	fmt.Println(Composite{Simple{10, "complex string"}, Complex{[3]int{4, 5, 6}, []string{"show", "me", "the", "code"}, map[string]int{"three": 3, "four": 4}}})

	oneID := OneIncDec{10}
	twoID := TwoIncDec{20}
	comp1 := Compositer1{&oneID, &twoID}
	//6
	fmt.Println(comp1.inc(5))
	//3
	fmt.Println(comp1.twoIncDec.dec(5))
	//10
	fmt.Println(oneID.n)
	//18
	fmt.Println(twoID.n)

	pointer_example()
}
