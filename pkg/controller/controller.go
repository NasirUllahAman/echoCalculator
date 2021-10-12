package controller

import (
	"fmt"
	"math"
	"net/http"

	"github.com/NasirUllahAman/echoCalculator/pkg/database"
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
	total := sum.Num1 + sum.Num2
	db := database.Conc()
	defer db.Close()

	sql := "INSERT INTO calcu(number1, number2, operation, result) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(sum.Num1, sum.Num2, "+", total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
	}

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
	fmt.Println(subtract.Num1)
	db := database.Conc()
	defer db.Close()
	total := subtract.Num1 - subtract.Num2
	sql := "INSERT INTO calcu(number1, number2, operation, result) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(subtract.Num1, subtract.Num2, "-", total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
	}
	res := Response{
		total,
	}

	fmt.Println(res)
	return c.JSON(http.StatusOK, res)

}
func Multiply(c echo.Context) error {
	mul := new(Calculator)
	if err := c.Bind(mul); err != nil {
		return err
	}
	fmt.Println(mul.Num1)
	db := database.Conc()
	defer db.Close()
	total := mul.Num1 * mul.Num2
	sql := "INSERT INTO calcu(number1, number2, operation, result) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(mul.Num1, mul.Num2, "*", total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
	}
	res := Response{
		total,
	}
	fmt.Println(res)

	return c.JSON(http.StatusOK, res)

}
func Division(c echo.Context) error {
	div := new(Calculator)
	if err := c.Bind(div); err != nil {
		return err
	}
	fmt.Println(div.Num1)
	db := database.Conc()
	defer db.Close()
	total := div.Num1 / div.Num2
	sql := "INSERT INTO calcu(number1, number2, operation, result) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(div.Num1, div.Num2, "/", total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
	}
	res := Response{
		total,
	}
	fmt.Println(res)

	return c.JSON(http.StatusOK, res)

}

func Modlus(c echo.Context) error {
	mod := new(Calculator)
	if err := c.Bind(mod); err != nil {
		return err
	}
	fmt.Println(mod.Num1)
	db := database.Conc()
	defer db.Close()
	total := float64(int64(mod.Num1) % int64(mod.Num2))
	sql := "INSERT INTO calcu(number1, number2, operation, result) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(mod.Num1, mod.Num2, "%", total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
	}

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
	fmt.Println(sq.Num1)
	db := database.Conc()
	defer db.Close()
	total := sq.Num1 * sq.Num1
	sql := "INSERT INTO calcu(number1, number2, operation, result) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(sq.Num1, sq.Num1, "*", total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
	}

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
	fmt.Println(sqt.Num1)
	db := database.Conc()
	defer db.Close()
	total := math.Sqrt(sqt.Num1)

	sql := "INSERT INTO calcu(number1, number2, operation, result) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(sqt.Num1, sqt.Num1, "*", total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
	}

	res := Response{
		total,
	}

	return c.JSON(http.StatusOK, res)

}
