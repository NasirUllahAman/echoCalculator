package route

import (
	"github.com/NasirUllahAman/echoCalculator/pkg/controller"
	"github.com/NasirUllahAman/echoCalculator/pkg/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {

	e.POST("/add", controller.Addition, middleware.AuthMiddleWare)
	e.POST("/subtraction", controller.Subtraction, middleware.AuthMiddleWare)
	e.POST("/multiply", controller.Multiply, middleware.AuthMiddleWare)
	e.POST("/division", controller.Division, middleware.AuthMiddleWare)
	e.POST("/mod", controller.Modlus, middleware.AuthMiddleWare)
	e.POST("/sqrt", controller.Sqroot, middleware.AuthMiddleWare)
	e.POST("/square", controller.Square, middleware.AuthMiddleWare)
	e.GET("/history", controller.GetAllRecords, middleware.AdminMiddleWare)
	e.GET("/history/:id", controller.GetRecord, middleware.AdminMiddleWare)
	e.DELETE("/history/:id", controller.DeleteRecord, middleware.AdminMiddleWare)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
