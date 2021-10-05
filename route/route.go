package route

import (
	"github.com/NasirUllahAman/GOL/controller"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {

	e.POST("/add", controller.Addition)
	e.POST("/Subtraction", controller.Subtraction)
	e.POST("/Multiply", controller.Multiply)
	e.POST("/Division", controller.Division)
	e.POST("/mod", controller.Modlus)
	e.POST("/sqrt", controller.Sqroot)
	e.POST("/Squ", controller.Square)
}
