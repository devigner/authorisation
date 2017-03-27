package user


import(
    "net/http"
    "io"
    "fmt"

    "github.com/devigner/authorisation/utils"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
    utils.Logger().Info("ListUsers")
    io.WriteString(w, fmt.Sprintf("ListUsers: %s", r.RequestURI))
}