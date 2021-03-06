package main

import (
    "github.com/nee-co/bandeja/route"
    "github.com/Sirupsen/logrus"
    "github.com/kelseyhightower/envconfig"
)

type Conf struct {
    PORT string
}

func init() {
    logrus.SetLevel(logrus.DebugLevel)
    logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
    var conf Conf

    err := envconfig.Process("BANDEJA", &conf)

    if err != nil {
        logrus.Error(err)
    }

    router := route.Init()
    router.Logger.Fatal(router.Start(":" + conf.PORT))
}