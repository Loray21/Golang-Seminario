package service

import (
	"github.com/Loray21/Golang-Seminario/Golang-Seminario/model"
	"github.com/Loray21/Golang-Seminario/Golang-Seminario/model/database"
)

// CAR define el comportamiento del servicio de vuelos
type Car interface {
	Add(*model.Car) *model.Car
	FindByID(ID uint) *model.Car
	Remove(ID uint)
}

type service struct {
	db *database.Database
}

// NewInstance devuelve una instancia del servicio de vuelos
func NewInstance(db *database.Database) Car {
	return &service{db}
}

func (s *service) Add(v *model.Car) *model.Car {
	s.db.Add(v)
	return v
}

func (s *service) FindByID(ID uint) *model.Car {
	return s.db.FindByID(ID)
}

func (s *service) Remove(ID uint) {
}
