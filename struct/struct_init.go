package main

type Simple struct {
	n   int32
	str string
}

type Complex struct {
	arr [3]int
	sl  []string
	m   map[string]int
}

type Composite struct {
	s Simple
	c Complex
}
