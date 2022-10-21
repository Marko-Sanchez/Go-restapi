package handlers

import (
    "net/http"
    "strings"

    "gopkg.in/mgo.v2/bson"
)

// Handles the users route
func UsersRouter(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimSuffix(r.URL.Path, "/")

    // all users
    if path == "/users" {
        switch r.Method {
        case http.MethodGet:
            usersGetAll(w, r)
            return
        case http.MethodPost:
            return
        default:
            postError(w, http.StatusMethodNotAllowed)
        }
    }

    path = strings.TrimPrefix(path, "/users/")
    if !bson.IsObjectIdHex(path) {
        postError(w, http.StatusNotFound)
        return
    }

    // single user
    // id := bson.ObjectIdHex(path)
    switch r.Method {
    case http.MethodGet:
        return
    case http.MethodPost:
        return
    case http.MethodPatch:
        return
    case http.MethodDelete:
        return
    default:
        postError(w, http.StatusMethodNotAllowed)
        return
    }
}
