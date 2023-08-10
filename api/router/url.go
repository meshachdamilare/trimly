package router

import (
	"github.com/labstack/echo/v4"
	"github/meshachdamilare/trimly/api/handler"
	"github/meshachdamilare/trimly/api/middleware"
	"github/meshachdamilare/trimly/repository/storage/postgres"
	"github/meshachdamilare/trimly/service"
)

func SetUrlRouter(e *echo.Echo) *echo.Echo {
	pgDb := postgres.GetDB()
	urlService := service.NewUrlService(pgDb)
	urlHandler := handler.NewUrlHandler(urlService)

	e.GET("/:code", urlHandler.Redirect)

	urlGrp := e.Group("/url")
	urlGrp.Use(middleware.AuthMiddleware)
	{
		urlGrp.POST("/trim", urlHandler.TrimUrl)
		urlGrp.GET("/trim", urlHandler.GetAllUrls)
	}
	return e
}
