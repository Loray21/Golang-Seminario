package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Loray21/Golang-Seminario/Golang-Seminario/model"
	"github.com/Loray21/Golang-Seminario/Golang-Seminario/model/database"
	"github.com/Loray21/Golang-Seminario/Golang-Seminario/service"
	"github.com/jeremiascaballero/ApiGolangMvc/controller"
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
	Auto := model.Auto{Name: strconv.Itoa(time.Now().Nanosecond()), Origin: "Ezeiza", Destiny: "Chicago"}

	// agrego el vuelo
	service.Add(&Auto)

	// pongo en la variable vuelo el contenido del punntero que me devuelve la funcion findByID
	auto = *service.FindByID(vuelo.ID)
	//      ^
	//      |
	//      Esto me devuelve el valor que aloja el puntero que me devuelve la
	//      funcion service.FindByID.

	fmt.Printf("\tID=%v\n", auto.ID)
	fmt.Printf("\tName=%v\n", auto.Name)
	fmt.Printf("\tOrigin=%v\n", auto.Origin)
	fmt.Printf("\tDestiby=%v\n", auto.Destiny)
	fmt.Printf("\tCreatedAt=%v\n", auto.CreatedAt)
	fmt.Printf("\tUpdatedAt=%v\n", Auto.UpdatedAt)
	fmt.Printf("\tDeletedAt=%v\n", auto.DeletedAt)
	// -------- ACA TERMINA EL EJEMPLO DE COMO USAR EL SERVICIO -------

	// Ahora voy a crear el servicio rest (http)
	http := controller.NewHTTPController(&service)
	http.Run()
}
