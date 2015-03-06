package main

import (
	"fmt"
	"net/http"
)

type Animal interface {
	Eat() string
}

type Cat struct {
}

func (c Cat) Eat() string {
	return "I'm a cat"
}

/* Functions are looked up by name other than signature
 * so there's no overloading stuff
func (c Cat) Eat(string target) string {
	return "Eat " + target
}
*/

type Dog struct {
}

func (d Dog) Eat() string {
	return "I'm a dog"
}

/* The following func does not work because we can't define new methods
 * on non-local types
func (h http.Request) Eat() string {
	return "I'm also an animal?"
}
*/

/* But we could achieve it in another way */
type AnimalRequest http.Request

func (r AnimalRequest) Eat() string {
	return "I'm an animal request"
}

/* subclass Dog with anonymous field; named field will not have the effect
 * of inheritance
 */
type PetDog struct {
	Dog
	name string
}

func (d PetDog) Eat() string {
	return "I'm a PetDog"
}

type HasLeg interface {
	NumOfLeg() int8
}

type WithLeg struct {
	n int8
}

func (w WithLeg) NumOfLeg() int8 {
	return w.n
}

type PetDogMore struct {
	Dog
	WithLeg
}

/* This will cause failure because both Dog and Cat
 * have an Eat() method which causes ambiguity
type PetDogCat struct {
	Dog
	Cat
}
*/

func interface_example() {
	fmt.Println("interface_example...")
	animals := []Animal{new(Cat), new(Dog), new(AnimalRequest), new(PetDog), new(PetDogMore)}
	for _, ani := range animals {
		fmt.Println(ani.Eat())
	}
}
