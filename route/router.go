package route

import (
    "github.com/nee-co/bandeja/handler"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
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