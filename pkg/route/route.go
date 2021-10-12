package route

import (
	"github.com/NasirUllahAman/echoCalculator/pkg/controller"
	"github.com/NasirUllahAman/echoCalculator/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {

	e.POST("/add", controller.Addition, middleware.AuthMiddleWare)
	e.POST("/Subtraction", controller.Subtraction, middleware.AuthMiddleWare)
	e.POST("/Multiply", controller.Multiply, middleware.AuthMiddleWare)
	e.POST("/Division", controller.Division, middleware.AuthMiddleWare)
	e.POST("/mod", controller.Modlus, middleware.AuthMiddleWare)
	e.POST("/sqrt", controller.Sqroot, middleware.AuthMiddleWare)
	e.POST("/Squ", controller.Square, middleware.AuthMiddleWare)
}
