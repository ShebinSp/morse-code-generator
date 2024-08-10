package main

import (
	"github.com/ShebinSp/morse-code-generator/internal/morse-service/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
