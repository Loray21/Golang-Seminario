package model

import "gorm.io/gorm"

/* VCar es la entidad que se va a persistir en la base de datos*/

type Car struct {
	gorm.Model
	Name   string
	Colour string
	Brand  string
	Price  int
}
