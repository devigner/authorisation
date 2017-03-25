package auth

import (
    "github.com/gorilla/mux"
    "net/http"
    "io"
    "fmt"
    "github.com/devigner/authorisation/utils"
)

// Router
// Setup auth package to connect to the main router
func Router(router *mux.Router) {
    a := router.PathPrefix("/auth").Subrouter()
    a.Path("/login").HandlerFunc(LoginHandler).
        Methods("GET").
        Name("auth-login")
    a.Path("/logout").HandlerFunc(LogoutHandler).
        Methods("GET").
        Name("auth-logout")
}

// LoginHandler
// Handle the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    utils.GetLogger().Info("LoginHandler")
    io.WriteString(w, fmt.Sprintf("User login: %s", r.RequestURI))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    utils.GetLogger().Info("LogoutHandler")
    io.WriteString(w, fmt.Sprintf("User logout: %s", r.RequestURI))
}