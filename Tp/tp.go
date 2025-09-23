package main

import (
	"fmt"
	"strings"
)

type Car struct {
	Brand string
	Name  string
	Power int
	Color string
	Trunk []T_item
}
type T_item struct {
	Nom      string
	Quantity int
}

func (c Car) afficher_une_voiture() {
	fmt.Println("===Mon véhicule===")
	fmt.Printf("Brand:%s\n", c.Brand)
	fmt.Printf("Name:%s\n", c.Name)
	fmt.Printf("Power:%d\n", c.Power)
	fmt.Printf("Color:%s\n", c.Color)
	if len(c.Trunk) < 1 {
		fmt.Println("Le coffre du véhicule est vide!!!")
		return
	}
	for _, val := range c.Trunk {
		fmt.Printf("%s , %d\n ", val.Nom, val.Quantity)
	}
}
func (c *Car) changer_couleur(color string) {
	if strings.EqualFold(c.Color, color) {
		fmt.Println("Le véhicule est déjà de cette couleur!!!")
	} else {
		c.Color = strings.ToLower(color)
		fmt.Printf("La couleur du véhicule à changé : %s\n", c.Color)
	}

}
func main() {
	item := T_item{"Jerrican", 2}
	car := Car{"Peugeot", "Peugeot206WRC", 280, "gris", []T_item{item}}
	car.afficher_une_voiture()
	car.changer_couleur("gris")
	car.changer_couleur("Noir")
}
