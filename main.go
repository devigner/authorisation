package main

import (
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "fmt"
    "os"
    "net/http"
    "io"

    "github.com/devigner/authorisation/auth"
    "github.com/devigner/authorisation/user"
    "github.com/devigner/authorisation/di"
)

func handler(w http.ResponseWriter, r *http.Request) {
    di.Container().Logger.Info("Uncaugh:", r.RequestURI)
    io.WriteString(w,fmt.Sprintf("You reach page %s",r.RequestURI))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(404)
    io.WriteString(w,fmt.Sprintf("404 %s",r.RequestURI))
}

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Errorf("Error env",err.Error())
    }

    dc := di.Setup()
    port := os.Getenv("APP_PORT")

    r := dc.Router.StrictSlash(false)
    r.HandleFunc("/", handler)

    auth.Router()
    user.Router()

    r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        t, err := route.GetPathTemplate()
        if err != nil {
            return err
        }
        dc.Logger.Info(fmt.Sprintf("Route: %s ; %s", t, route.GetName()))
        return nil
    })

    r.NotFoundHandler = http.HandlerFunc(NotFound)

    dc.Logger.Info("App running on port:", port)
    dc.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),r))
}
