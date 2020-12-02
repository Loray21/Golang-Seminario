package database

import (
	"github.com/Loray21/Golang-Seminario/model"
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

// FindAll devuelve todo
func (database *Database) FindAll() []model.Car {
	var autos []model.Car
	database.db.Find(&autos)
	return autos
}

// CreateSchema create el esquema a partir de las entidades.
func (database *Database) CreateSchema() {
	database.db.AutoMigrate(&model.Car{})
}

func (database *Database) Remove(i uint) {
	var v model.Car
	database.db.Unscoped().Delete(&v, i)
}

// Add inserta un auto
func (database *Database) Add(v *model.Car) {
	database.db.Create(v)
}

// FindByID ...
func (database *Database) FindByID(ID uint) *model.Car {
	var v model.Car
	database.db.Find(&v, ID)
	return &v
}
