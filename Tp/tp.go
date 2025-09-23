package main

import "fmt"

type Car struct {
	Brand string
	Name  string
	Power int
	Color string
	Trunk Item
}
type Item struct {
	Nom      string
	Quantity int
}

func main() {
	car := Car{
		Brand: "Peugeot",
		Name:  "Peugeot206WRC",
		Power: 280,
		Color: "Gris",
		Trunk: Item{
			Nom:      "Jerrican",
			Quantity: 2,
		},
	}
	fmt.Println(car)
}
