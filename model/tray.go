package model

import (
    "github.com/gocraft/dbr"
)

type (
    Tray struct {
        Id       int    `db:"id"       json:"id"`
        Title    string `db:"title"    json:"title"`
        Endpoint string `db:"endpoint" json:"endpoint"`
    }

    Trays struct {
        Trays    []Tray `json:"trays"`
    }
)

func (t *Tray)GetTray(s dbr.Session,id int) error {
    _, err := s.Select("*").
                From("trays").
                Where("id = ?", id).
                Load(t)

    return err
}

func GetTrays(s dbr.Session) (*Trays, error) {
    var results []Tray
    trays := new(Trays)
    _, err := s.Select("*").
                From("trays").
                Load(&results)

    if err == nil {
        trays.Trays = results
    }

    return trays, err
}