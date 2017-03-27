package user


import(
    "net/http"
    "io"
    "fmt"

    //"github.com/devigner/authorisation/utils"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
    //utils.GetConnection()
    //utils.GetLogger().Info("ListUsers")
    io.WriteString(w, fmt.Sprintf("Users: %s", r.RequestURI))
}