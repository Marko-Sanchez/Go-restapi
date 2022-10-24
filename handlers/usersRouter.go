package handlers

import (
    "net/http"
    "strings"

    "gopkg.in/mgo.v2/bson"
)

// Handles the users route
// @param: r {*http.Request} holds clients request.
//         w {http.ResponseWriter} structure lets us reply to client.
func UsersRouter(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimSuffix(r.URL.Path, "/")

    // query all users, or create a new user
    if path == "/users" {
        switch r.Method {
        case http.MethodGet:
            usersGetAll(w, r)
            return
        case http.MethodPost:
            usersPostOne(w, r)
            return
        default:
            postError(w, http.StatusMethodNotAllowed)
            return
        }
    }

    // grab ID passed by client
    path = strings.TrimPrefix(path, "/users/")
    if !bson.IsObjectIdHex(path) {
        postError(w, http.StatusNotFound)
        return
    }

    // single user
    id := bson.ObjectIdHex(path)
    switch r.Method {
    case http.MethodGet:
        usersGetOne(w, r, id)
        return
    case http.MethodPut:
        usersPutOne(w, r, id)
        return
    case http.MethodPatch:
        usersPatchOne(w, r, id)
        return
    case http.MethodDelete:
        usersDeleteOne(w, r, id)
        return
    default:
        postError(w, http.StatusMethodNotAllowed)
        return
    }
}
