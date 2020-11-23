package model

import "gorm.io/gorm"

// Vuelo es la entidad que se va a persistir en la base de datos
type Car struct {
	gorm.Model
	name   string
	colour string
	brand  string
	price  int
}
