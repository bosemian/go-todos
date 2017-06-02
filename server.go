package main

import (
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	e.POST("/tasks", func(c echo.Context) error { return c.JSON(200, "POST Tasks") })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	e.Logger.Fatal(e.Start(":1323"))

}
