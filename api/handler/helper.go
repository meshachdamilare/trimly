package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func CreateCookie(name string, value string, time int, c echo.Context, httpOnly bool) *http.Cookie {

	host := c.Request().Host

	// Remove the port number from the host if present
	if strings.Contains(host, ":") {
		host = host[:strings.Index(host, ":")]
	}
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   host,
		MaxAge:   time,
		Secure:   false,
		HttpOnly: httpOnly,
	}
}
