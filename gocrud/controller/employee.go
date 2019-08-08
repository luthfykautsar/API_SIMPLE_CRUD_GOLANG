package controller

import (
	"gocrud/model"
	"net/http"
	"github.com/labstack/echo"
	"strconv"
	"fmt"
)

func GetEmployees(c echo.Context) error {
	result := model.GetEmployee()
	println("get")

	return c.JSON(http.StatusOK,result)
}

func UpdateEmloyees(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	city := c.FormValue("city")

	model.UpdateEmployee(id, name, email, password, city)

	json := model.Employee{Id:id, Name:name, Email:email, Password:password, City:city}

	return c.JSON(http.StatusOK, json)
}

func SaveEmployees(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	city := c.FormValue("city")

	model.SaveEmployee(name, email, password, city)
	return c.JSON(http.StatusCreated, "SAVE SUCCESS")
}

func DeleteEmployees(c echo.Context)error {
	id, _ := strconv.Atoi(c.Param("id"))

	model.DeleteEmployee(id)
	return  c.JSON(http.StatusOK, "DELETE SUCCESS")
}

func LoginEmployees(c echo.Context) error  {
	email := c.FormValue("email")
	password := c.FormValue("password")

	model.Login(email,password)

	return c.String(http.StatusOK, fmt.Sprintf("LOGIN SUCCESS \n WELCOME :"+email))
}