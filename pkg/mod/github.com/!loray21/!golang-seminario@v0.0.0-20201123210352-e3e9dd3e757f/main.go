package main

import (
	"fmt"

	"github.com/Loray21/Golang-Seminario/Golang-Seminario/controller"
	"github.com/Loray21/Golang-Seminario/Golang-Seminario/model/database"
	"github.com/Loray21/Golang-Seminario/Golang-Seminario/service"
)

func main() {
	// Creo una instancia de la base de datos
	db := database.NewDatabase("database.db") // aca esta hardcoded el nombre de la base de datos
	// Creo el esquema
	db.CreateSchema()

	// creo el servicio y le paso la base de datos
	service := service.NewInstance(db)

	// ------- ESTO ES A MODO DE EJEMPLO DE COMO USAR EL SERVICIO -----------
	// genero un vuelo a insertar cuyo nombre es un numero
	Car := Model.Car{name: "hilux", colour: "white", brand: "toyota", price: 200}

	// agrego el vuelo
	service.Add(&Car)

	// pongo en la variable vuelo el contenido del punntero que me devuelve la funcion findByID
	Car = *service.FindByID(Car.ID)
	//      ^
	//      |
	//      Esto me devuelve el valor que aloja el puntero que me devuelve la
	//      funcion service.FindByID.
	fmt.Printf("\tID=%v\n", Car.ID)
	fmt.Printf("\tname=%v\n", Car.name)
	fmt.Printf("\tcolour=%v\n", Car.colour)
	fmt.Printf("\tbrand=%v\n", Car.brand)
	fmt.Printf("\tprice=%v\n", Car.price)
	fmt.Printf("\tCreatedAt=%v\n", Car.CreatedAt)
	fmt.Printf("\tUpdatedAt=%v\n", Car.UpdatedAt)
	fmt.Printf("\tDeletedAt=%v\n", Car.DeletedAt)
	// -------- ACA TERMINA EL EJEMPLO DE COMO USAR EL SERVICIO -------

	// Ahora voy a crear el servicio rest (http)
	http := controller.NewHTTPController(&service)
	http.Run()
}
