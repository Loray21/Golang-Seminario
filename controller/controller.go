package controller

import (
	"net/http"
	"strconv"

	"github.com/Loray21/GolangSeminario/Golang-Seminario/service"
	"github.com/gin-gonic/gin"
)

// HTTPController ...
type HTTPController truct {
	router  *gin.Engine
	ervice *service.Car


// NewHTTPController ...
func NewHTTPControler(s *service.Vuelo) HTTPController {
	r := gin.Default()
	makeEndpoints(r, s) // regitro los endpoints
	eturn HTTPController{r, s}


// Run ejecuta el controller
func (c *HTTPCotroller) Run() {
	.router.Run()


func makeEndpoints(r *gin.Engine, s *servce.Vuelo) {
	.GET("/vuelos/:id", makeFindHandler(s))


func makeFindHandler(s *service.Vuelo) gin.HandlerFun {
	// Fijate que aca devuelvo un funcion y no un valor
	return func(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("id"))

	v := (*s).FindByID(uint(i))

		c.JSON(http.tatusOK, gin.H{
			"uelo": v,
		)
	
}
