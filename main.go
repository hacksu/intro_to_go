package main

import "fmt"

type product struct {
	name string
	cost int
}

func main() {
	p := product{
		name: "Brick",
		cost: 10,
	}

	fmt.Println(p.name, "costs $", p.cost)
}
