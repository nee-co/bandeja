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

    e.GET("/trays", handler.Trays)
    e.GET("/trays/:id", handler.Tray)

    return e
}