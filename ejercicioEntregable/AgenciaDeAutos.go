package main

import "fmt"

type Car struct {
	name   string
	colour string
	brand  string
	price  int
}

var Agencia []Car
var Agencias []int

//agregar autos

func addCar(a Car) {
	Agencia = append(Agencia, a)
}

func Read() {
	for i, Agencia := range Agencia {
		fmt.Println(i, Agencia)
	}
}

func RemoveIndex(s []Car, index int) []Car {
	return append(s[:index], s[index+1:]...)
}

func update(s []Car, name string, colour string, brand string, price int, index int) {
	s[index].name = name
	s[index].colour = colour
	s[index].brand = brand
	s[index].price = price

}

func main() {

	p := Car{name: "hilux", colour: "gray", brand: "toyota", price: 1200}
	addCar(p)
	p1 := Car{name: "hilux", colour: "gray", brand: "toyota", price: 1200}
	addCar(p1)

	Read()
	Agencia = RemoveIndex(Agencia, 1)
	update(Agencia, "rodrigo", "marron", "nike", 200, 0)
	Read()

}
