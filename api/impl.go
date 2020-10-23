package api

import (
	"github.com/labstack/echo/v4"
)

type APIEcho struct {
	engine *echo.Echo
}

func NewAPIHandler() *APIEcho {
	e := echo.New()
	return &APIEcho{engine: e}

}

func (a *APIEcho) RouteInit() {

	a.engine.GET("/", index)
}

func (a *APIEcho) Run(port string) {

	a.engine.Logger.Fatal(a.engine.Start(":" + port))

}

func (a APIEcho) Info() string {

	return "Using Echo Framework"
}
