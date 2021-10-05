package controller

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Calculator struct {
	Num1 float64 `json:"number1"`
	Num2 float64 `json:"number2"`
}
type Response struct {
	Result float64 `json:"result"`
}

func Addition(c echo.Context) error {

	sum := new(Calculator)

	if err := c.Bind(sum); err != nil {
		return err
	}
	fmt.Println(sum.Num1)
	total := sum.Num1 + sum.Num2
	res := Response{
		total,
	}
	fmt.Println(res)

	return c.JSON(http.StatusOK, res)

}
func Subtraction(c echo.Context) error {
	subtract := new(Calculator)
	if err := c.Bind(subtract); err != nil {
		return err
	}
	total := subtract.Num1 - subtract.Num2
	res := Response{
		total,
	}

	return c.JSON(http.StatusOK, res)

}
func Multiply(c echo.Context) error {
	mul := new(Calculator)
	if err := c.Bind(mul); err != nil {
		return err
	}
	total := mul.Num1 * mul.Num2

	res := Response{
		total,
	}

	return c.JSON(http.StatusOK, res)

}
func Division(c echo.Context) error {
	div := new(Calculator)
	if err := c.Bind(div); err != nil {
		return err
	}
	total := div.Num1 / div.Num2

	res := Response{
		total,
	}

	return c.JSON(http.StatusOK, res)

}

func Modlus(c echo.Context) error {
	mod := new(Calculator)
	if err := c.Bind(mod); err != nil {
		return err
	}
	total := float64(int64(mod.Num1) % int64(mod.Num2))

	res := Response{
		total,
	}

	return c.JSON(http.StatusOK, res)

}
func Square(c echo.Context) error {
	sq := new(Calculator)
	if err := c.Bind(sq); err != nil {
		return err
	}
	total := sq.Num1 * sq.Num1

	res := Response{
		total,
	}

	return c.JSON(http.StatusOK, res)
}
func Sqroot(c echo.Context) error {
	sqt := new(Calculator)
	if err := c.Bind(sqt); err != nil {
		return err
	}
	total := math.Sqrt(sqt.Num1)
	res := Response{
		total,
	}

	return c.JSON(http.StatusOK, res)

}
