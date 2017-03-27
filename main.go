package main

import (
    "fmt"
    "os"
    "net/http"
    "io"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "github.com/devigner/authorisation/auth"
    "github.com/devigner/authorisation/user"
    "github.com/devigner/authorisation/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
    utils.Logger().Info("Uncaugh:", r.RequestURI)
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

    port := os.Getenv("APP_PORT")

    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/", handler)

    auth.Router(r)
    user.Router(r)

    r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        t, err := route.GetPathTemplate()
        if err != nil {
            return err
        }
        utils.Logger().Info(fmt.Sprintf("Route: %s ; %s", t, route.GetName()))
        return nil
    })

    r.NotFoundHandler = http.HandlerFunc(NotFound)

    utils.Logger().Info("App running on port:", port)
    utils.Logger().Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),r))
}
