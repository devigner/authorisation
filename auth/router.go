package auth

import (
    "net/http"
    "io"
    "fmt"
    "github.com/devigner/authorisation/di"
)

// Router
// Setup auth package to connect to the main router
func Router() {
    a := di.Container().Router.PathPrefix("/auth").Subrouter()
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
    di.Container().Logger.Info("LoginHandler")
    io.WriteString(w, fmt.Sprintf("User login: %s", r.RequestURI))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    di.Container().Logger.Info("LogoutHandler")
    io.WriteString(w, fmt.Sprintf("User logout: %s", r.RequestURI))
}
