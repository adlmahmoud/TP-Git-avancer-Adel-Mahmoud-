package main

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
