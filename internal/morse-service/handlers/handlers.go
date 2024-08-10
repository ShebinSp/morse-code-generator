package handlers

import (
	"net/http"

	"github.com/ShebinSp/morse-code-generator/internal/morse-service/services"
	"github.com/labstack/echo"
)

func GenerateMorseCode(c echo.Context) error {

	text := c.QueryParam("text")

	morseCode := services.ConvertToMorseCode(text)

	return c.JSON(http.StatusOK, map[string]string{
		"morse_code": morseCode,
	})
}

func GenerateMorseToLetter(c echo.Context) error {
	morse := c.QueryParam("morse")

	morseToLetter := services.MorseToLetter(morse)

	return c.JSON(http.StatusOK, map[string]string{
		"text": morseToLetter,
	})
}
