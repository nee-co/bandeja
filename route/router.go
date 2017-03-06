package route

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "bandeja/handler"
)

func Init() *echo.Echo {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/trays",     handler.Trays)
    e.GET("/trays/:id", handler.Tray)
    e.GET("/trays/my",  handler.UserTrays)
    e.PUT("/trays/my",  handler.CreateOrUpdate)
    return e
}