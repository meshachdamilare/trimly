package router

import (
	"github.com/labstack/echo/v4"
	"github/meshachdamilare/trimly/api/handler"
	"github/meshachdamilare/trimly/api/middleware"
	"github/meshachdamilare/trimly/repository/storage/postgres"
	"github/meshachdamilare/trimly/service"
)

func SetUserRouter(e *echo.Echo) *echo.Echo {
	pgDb := postgres.GetDB()
	userService := service.NewUserService(pgDb)
	userHandler := handler.NewUserHandler(userService)

	userGrp := e.Group("/users")
	{

		authGrp := userGrp.Group("/auth")
		{
			authGrp.POST("/register", userHandler.Register)
			authGrp.POST("/login", userHandler.Login)
			authGrp.GET("/logout", middleware.AuthMiddleware(userHandler.LogoutUser))
		}

		userGrp.Use(middleware.AuthMiddleware)
		userGrp.GET("/me", userHandler.Me)
		userGrp.GET("/urls", userHandler.GetUserURLs)

	}
	return e
}
