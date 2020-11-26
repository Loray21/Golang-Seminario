package database

import (
	"github.com/Loray21/Golang-Seminario/Golang-Seminario/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database ...
type Database struct {
	db *gorm.DB
}

// NewDatabase ...
func NewDatabase(dsn string) *Database { //dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Database{db}
}

// FindAll devuelve todos los vuelos registrados
func (database *Database) FindAll() []model.Car {
	var vuelos []model.Car
	database.db.Find(&vuelos)
	return vuelos
}

// CreateSchema create el esquema a partir de las entidades.
func (database *Database) CreateSchema() {
	database.db.AutoMigrate(&model.Car{})
}

// Add inserta un vuelo
func (database *Database) Add(v *model.Car) {
	database.db.Create(v)
}

// FindByID ...
func (database *Database) FindByID(ID uint) *model.Car {
	var v model.Car
	database.db.Find(&v, ID)
	return &v
}
