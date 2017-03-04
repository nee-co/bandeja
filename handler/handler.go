package handler

import (
    "bandeja/db"
    "bandeja/model"
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "github.com/Sirupsen/logrus"
)

var session = db.GetSession()

func Tray(c echo.Context) error {
    if c.Request().Header.Get("x-consumer-custom-id") == "" {
        return c.JSON(http.StatusUnauthorized, nil)
    }

    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        logrus.Error(err)
        return err
    }

    var tray model.Tray

    if err := tray.GetTray(*session, id); err != nil {
        c.Error(err)
        return err
    } else if tray.Id != 0 {
        return c.JSON(http.StatusOK, tray)
    } else {
        return c.JSON(http.StatusNotFound, nil)
    }
}

func Trays(c echo.Context) error {
    if c.Request().Header.Get("x-consumer-custom-id") == "" {
        return c.JSON(http.StatusUnauthorized, nil)
    }

    trays, err := model.GetTrays(*session)

    if err != nil {
        c.Error(err)
        return err
    } else {
        return c.JSON(http.StatusOK, trays)
    }
}