// @SubApi Some API [/service/swagger]
package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// service is http server
type api struct{}

// Init creates and returns new api with handlers
func Init() *api {
	return &api{}
}

//
// Handler1 is handler that returns some text
//
// @Title Return some text
// @Description Some endpoint description
// @Accept  json
// @Success 200 {object}  string
// @Resource swagger/handler1
// @Router /service/swagger/handler1 [get]
func (a *api) Handler1(c echo.Context) error {
	return c.HTML(http.StatusOK, "handler 1")
}

//
// Handler2 is handler that returns some text
//
// @Title Return some text
// @Description Some endpoint description
// @Accept  json
// @Success 200 {object}  string
// @Resource swagger/handler2
// @Router /service/swagger/handler2 [get]
func (a *api) Handler2(c echo.Context) error {
	return c.HTML(http.StatusOK, "handler 2")
}
