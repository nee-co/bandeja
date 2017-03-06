package model

import (
    "github.com/gocraft/dbr"
    "github.com/Sirupsen/logrus"
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

func GetUserTray(s dbr.Session, userId int, spaceId int) (*UserTray, error) {
    var userTray UserTray
    _, err := s.Select("*").
                From("users_trays").
                Where("user_id = ? AND space_id = ?", userId, spaceId).
                Load(&userTray)

    return &userTray, err
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

func (t *UserTray) Create(s dbr.Session) error {
    _, err := s.InsertInto("users_trays").
                Columns("user_id", "tray_id", "space_id").
                Record(t).
                Exec()

    return err
}

func (t *UserTray) Update(s dbr.Session, trayId int) error {
    _, err := s.Update("users_trays").
                Set("tray_id", trayId).
                Where("user_id = ? AND space_id = ?", t.UserId, t.SpaceId).
                Exec()

    return err
}

func UsableSpaceId(s dbr.Session, userId int) int {
    var results []UserTray
    count, err := s.Select("*").
                    From("users_trays").
                    Where("user_id = ?", userId).
                    Load(&results)

    if err != nil {
        logrus.Error(err)
        return 0
    }

    return count + 1
}