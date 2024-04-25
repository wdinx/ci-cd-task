package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// Connect With MySql
	LoadEnv()
	InitDB(InitConfigMySQL())

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "succesfully connected",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
