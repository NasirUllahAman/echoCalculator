package main

import (
	"github.com/NasirUllahAman/echoCalculator/pkg/route"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	route.Route(e)
	e.Logger.Fatal(e.Start(":1331"))
}
