package routes

import (
	"github.com/ShebinSp/morse-code-generator/internal/morse-service/handlers"
	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/morse", handlers.GenerateMorseCode)
	e.POST("/letter", handlers.GenerateMorseToLetter)
}
