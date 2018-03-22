package handler

import (
	"github.com/hesidoryn/swagger-example/swagger"
	"github.com/labstack/echo"
)

// GetSwagger loads swagger's modules
func (a *api) GetSwagger(c echo.Context) error {
	return swagger.LoadModule(c)
}
