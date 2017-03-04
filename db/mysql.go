package db

import (
    "github.com/gocraft/dbr"
    "github.com/kelseyhightower/envconfig"
    "github.com/Sirupsen/logrus"
    _ "github.com/go-sql-driver/mysql"
)

type DATABASE struct {
    HOST     string
    USER     string
    PASSWORD string
}

func GetSession() *dbr.Session {
    var conf DATABASE

    err := envconfig.Process("BANDEJA_DATABASE", &conf)

    if err != nil {
        logrus.Error(err)
    }

    db, err := dbr.Open("mysql",
        conf.USER+":"+conf.PASSWORD+"@tcp("+conf.HOST+":3306)/bandeja_production",
        nil)

    if err != nil {
        logrus.Error(err)
    } else {
        session := db.NewSession(nil)
        return session
    }

    return nil
}