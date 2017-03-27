package user

import (
    "net/http"
    "io"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/devigner/authorisation/utils"
)

func Router(router *mux.Router) {
    a := router.PathPrefix("/user").Subrouter()
    a.Path("/").HandlerFunc(ListUsers)
    a.Path("/logout").HandlerFunc(LogoutHandler)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    utils.Logger().Info("LogoutHandler")
    io.WriteString(w, fmt.Sprintf("User logout: %s", r.RequestURI))
}
