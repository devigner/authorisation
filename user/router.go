package user

import (
    "github.com/gorilla/mux"
    "net/http"
    "io"
    "fmt"
    "github.com/devigner/authorisation/utils"
)

func Router(router *mux.Router) {
    a := router.PathPrefix("/user").Subrouter()
    a.Path("/").HandlerFunc(ListUsers)
    a.Path("/logout").HandlerFunc(LogoutHandler)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
    utils.GetLogger().Info("ListUsers")
    io.WriteString(w, fmt.Sprintf("Users: %s", r.RequestURI))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    utils.GetLogger().Info("LogoutHandler")
    io.WriteString(w, fmt.Sprintf("User logout: %s", r.RequestURI))
}
