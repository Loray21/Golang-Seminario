package model

// VCar es la entidad que se va a persistir en la base de datos
type Car struct {
	gorm.model
	name   string
	colour string
	brand  string
	price  int
}
