package di

import (
    log "github.com/Sirupsen/logrus"
    "github.com/jmoiron/sqlx"
    "github.com/gorilla/mux"
    "github.com/codegangsta/inject"
    "github.com/devigner/authorisation/utils"
)

var dc = DependencyContainer{}

type DependencyContainer struct {
    Logger *log.Logger `inject`
    Db *sqlx.DB `inject`
    Router *mux.Router `inject`
}

func Container() DependencyContainer{
    return dc
}

func Setup() DependencyContainer{
    error := getInjector().Apply(&dc)
    if error != nil {

    }
    return dc
}

func getInjector() inject.Injector {
    injector := inject.New()
    injector.Map(utils.CreateLogger())
    injector.Map(utils.GetDbConnection())
    injector.Map(mux.NewRouter())
    return injector
}