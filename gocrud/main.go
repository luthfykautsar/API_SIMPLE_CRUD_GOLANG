package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"gocrud/controller"
)


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:[]string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return  c.JSON(http.StatusCreated,"WELCOME")
	})

	e.GET("/employee/login", controller.LoginEmployees)
	e.GET("/employee", controller.GetEmployees)
	e.POST("/employee", controller.SaveEmployees)
	e.DELETE("/employee/:id", controller.DeleteEmployees)
	e.PUT("/employee/:id", controller.UpdateEmloyees)

	e.Start(":8000")

}

