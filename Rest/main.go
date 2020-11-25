package main

import (
	"github.com/gin-gonic/gin"
)

//para pedirle algo especifico por id digamos	r.GET("/users/:name", getUserHandler)

func main() {
	r := gin.Default()

	//cuando te llegue un get con tal cosa contestatalo con esto
	r.GET("/users", getUserHandler)

	r.GET("/users/:name", getUserHandlerID)

	//create

	r.POST("/users", addUserhandler)

	//ejecuta el servidor en el puerto 8080 por es el default y escuha asociando la ruta con funcion por parametro

	r.Run()
}

//le paso un punterio a gin.Context para que el router lo entienda
func getUserHandler(c *gin.Context) {

	//2 parametros request status , interface
	c.JSON(200, gin.H{
		"status": "ok",
	})

}

func getUserHandlerID(c *gin.Context) {

	//por id

	name := c.Param("name")
	//otro tipo de parametro
	lastName := c.Query("lastname")
	c.JSON(200, gin.H{
		"name": name + "" + lastName,
	})

}

//201 mensaje de creado
func addUserhandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"name ":    "tomi",
		"lastname": "loray",
	})
}
