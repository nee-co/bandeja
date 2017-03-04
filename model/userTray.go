package model

import (
    "github.com/gocraft/dbr"
)

type (
    UserTray struct {
        UserId  int `db:"user_id"  json:"-"`
        SpaceId int `db:"space_id" json:"space_id"`
        TrayId  int `db:"tray_id"  json:"tray_id"`
    }

    UserTrays struct {
        UserTrays []UserTray `json:"trays"`
    }
)

func NewUserTray(userId int, trayId int, spaceId int) *UserTray {
    return &UserTray{
        UserId:  userId,
        TrayId:  trayId,
        SpaceId: spaceId,
    }
}

func GetUserTrays(s dbr.Session, userId int) (*UserTrays, error) {
    var results []UserTray
    userTrays := new(UserTrays)
    _, err := s.Select("*").
                From("users_trays").
                Where("user_id = ?", userId).
                Load(&results)

    if err == nil {
        userTrays.UserTrays = results
    }

    return userTrays, err
}