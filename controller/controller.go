package controller

import (
	"net/http"

	"strconv"

	"github.com/Loray21/Golang-Seminario/Golang-Seminario/service"
	"github.com/gin-gonic/gin"
)

// HTTPController ...
type HTTPController struct {
	router  *gin.Engine
	service *service.Car
}

// NewHTTPController ...
func NewHTTPController(s *service.Car) HTTPController {
	r := gin.Default()
	makeEndpoints(r, s) // registro los endpoints
	return HTTPController{r, s}
}

// Run ejecuta el controller
func (c *HTTPController) Run() {
	c.router.Run()
}

func makeEndpoints(r *gin.Engine, s *service.Car) {
	r.GET("/autos/:id", makeFindHandler(s))
}

func makeFindHandler(s *service.Car) gin.HandlerFunc {
	// Fijate que aca devuelvo una funcion y no un valor
	return func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("id"))

		v := (*s).FindByID(uint(i))

		c.JSON(http.StatusOK, gin.H{
			"auto": v,
		})
	}
}
