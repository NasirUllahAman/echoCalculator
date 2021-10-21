package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("I am coming")
		if len(c.Request().Header["Authorization"]) > 0 {
			if c.Request().Header["Authorization"][0] == "secretkey" {
				c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
				return next(c)
			}
		}
		return c.JSON(http.StatusForbidden, "You are not authorized!")
	}
}

func AdminMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("I am coming here")

	return func(c echo.Context) error {
		if len(c.Request().Header["Authorization"]) > 0 {
			if c.Request().Header["Authorization"][0] == "secret123" {
				c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
				return next(c)
			}
		}
		return c.JSON(http.StatusForbidden, "You are not authorized!")
	}
}
