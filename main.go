package main

import (
    "github.com/gorilla/mux"
    "fmt"
    "os"
    "net/http"
    "io"

    "github.com/devigner/authorisation/auth"
    "github.com/devigner/authorisation/utils"
    "github.com/devigner/authorisation/user"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log := utils.GetLogger()
    log.Info("Request:", r.RequestURI)
    io.WriteString(w,fmt.Sprintf("You reach page %s",r.RequestURI))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(404)
    io.WriteString(w,fmt.Sprintf("404 %s",r.RequestURI))
}

func main() {
    port, ok := os.LookupEnv("APP_PORT")
    if !ok {
        port = "3000"
    }

    r := mux.NewRouter()
    r.HandleFunc("/", handler)
    auth.Router(r)
    user.Router(r)

    r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        t, err := route.GetPathTemplate()
        if err != nil {
            return err
        }

        utils.GetLogger().Info(fmt.Sprintf("Route: %s ; %s", t, route.GetName()))
        return nil
    })

    r.NotFoundHandler = http.HandlerFunc(NotFound)

    utils.GetLogger().Info("App running on port:", port)
    utils.GetLogger().Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),r))
}
