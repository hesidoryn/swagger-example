// @APIVersion 1.0.0
// @APITitle API for some service
// @APIDescription My API usually works as expected.
// @Contact api@gmail.com
// @BasePath /
package router

import (
	"github.com/hesidoryn/swagger-example/handler"
	"github.com/labstack/echo"
)

// Load is used to get router
func Load() *echo.Echo {
	e := echo.New()

	h := handler.Init()

	service := e.Group("service")
	service.Static("/swagger-ui", "/home/sidoringi/go/src/github.com/hesidoryn/swagger-example/static/swagger-ui")
	service.GET("/swagger/:module", h.GetSwagger)

	// handler 1
	service.GET("/handler1", h.Handler1)

	// handler 2
	service.GET("/handler2", h.Handler2)

	return e
}
