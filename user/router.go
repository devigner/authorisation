package user

import (
    "net/http"
    "io"
    "fmt"
    //"github.com/devigner/authorisation/utils"
    "github.com/devigner/authorisation/di"
)

func Router() {
    a := di.Container().Router.PathPrefix("/user").Subrouter()
    a.Path("/").HandlerFunc(ListUsers)
    a.Path("/logout").HandlerFunc(LogoutHandler)
}



func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    di.Container().Logger.Info("LogoutHandler")
    io.WriteString(w, fmt.Sprintf("User logout: %s", r.RequestURI))
}
