package middleware

import (
	"github.com/labstack/echo/v4"
	"github/meshachdamilare/trimly/settings/constant"
	"github/meshachdamilare/trimly/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//var accessToken string
		accessToken := c.Request().Header.Get("Authorization")
		if accessToken != "" {
			// Remove the "Bearer " prefix from the token string
			accessToken = strings.Split(accessToken, " ")[1]
		} else {
			// Token not found in headers, try to extract from the cookie
			cookie, err := c.Cookie("access_token")
			if err == nil {
				accessToken = cookie.Value
			}
		}
		if accessToken == "" {
			rd := utils.ErrorResponse(http.StatusUnauthorized, constant.StatusFailed, "no token provided", "Unauthorized error")
			return c.JSON(http.StatusUnauthorized, rd)
		}
		claims, err := ValidateToken(accessToken)
		if err != nil {
			rd := utils.ErrorResponse(http.StatusUnauthorized, constant.StatusFailed, err.Error(), "Unauthorized error")
			return c.JSON(http.StatusUnauthorized, rd)
		}
		userId := claims.ID
		email := claims.Email
		role := claims.Role

		c.Set("userId", userId)
		c.Set("email", email)
		c.Set("role", role)

		return next(c)
	}
}
