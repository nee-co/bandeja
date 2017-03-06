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

func UserTrays(c echo.Context) error {
    userId, err := strconv.Atoi(c.Request().Header.Get("x-consumer-custom-id"))

    if err != nil {
        return c.JSON(http.StatusUnauthorized, nil)
    }

    userTrays, err := model.GetUserTrays(*session, userId)

    if err != nil {
        c.Error(err)
        return err
    } else {
        return c.JSON(http.StatusOK, userTrays)
    }
}

func CreateOrUpdate(c echo.Context) error {
    userId, _ := strconv.Atoi(c.Request().Header.Get("x-consumer-custom-id"))

    if userId == 0 {
        return c.JSON(http.StatusUnauthorized, nil)
    }

    params,  err := c.FormParams()

    if err != nil {
        logrus.Error(err)
        return err
    }

    trayId,  _ := strconv.Atoi(params.Get("tray_id"))
    spaceId, _ := strconv.Atoi(params.Get("space_id"))

    if trayId == 0 || spaceId == 0 {
        return c.JSON(http.StatusUnprocessableEntity, nil)
    }

    userTray, err := model.GetUserTray(*session, userId, spaceId)

    if err != nil {
        logrus.Error(err)
    } else if !model.IsUsableTray(*session, trayId) {
        logrus.Error("Not found tray.")
    } else if userTray.SpaceId != 0 {
        err := userTray.Update(*session, trayId)

        if err != nil {
            logrus.Error(err)
        } else {
            return c.JSON(http.StatusNoContent, nil)
        }
    } else if spaceId == model.UsableSpaceId(*session, userId) {
        err := model.NewUserTray(userId, trayId, spaceId).
                     Create(*session)

        if err != nil {
            logrus.Error(err)
        } else {
            return c.JSON(http.StatusNoContent, nil)
        }
    } else {
        return c.JSON(http.StatusUnprocessableEntity, nil)
    }

    return c.JSON(http.StatusInternalServerError, nil)
}