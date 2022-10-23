package handlers

import (
    "Go-restapi/user"
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"

    "gopkg.in/mgo.v2/bson"
)

// convert request into a user
func bodyToUser(r *http.Request, u *user.User) error {
    if r.Body == nil {
        return errors.New("request body is empty")
    }
    if u == nil {
        return errors.New("a user is required")
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return err
    }
    return json.Unmarshal(body, u)
}

// Get all user data in database
func usersGetAll(w http.ResponseWriter, r *http.Request) {
    users, err := user.All()
    if err != nil {
        postError(w, http.StatusInternalServerError)
        return
    }
    postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}

// create a new user
func usersPostOne(w http.ResponseWriter, r *http.Request) {
    u := new(user.User)
    err := bodyToUser(r, u)
    if err != nil {
        postError(w, http.StatusBadRequest)
        return
    }

    // create unique id for user
    u.ID = bson.NewObjectId()
    err = u.Save()
    if err != nil {
        if err == user.ErrRecordInvalid {
            postError(w, http.StatusBadRequest)
        } else {
            postError(w, http.StatusInternalServerError)
        }

        return
    }

    w.Header().Set("Location", "/users/" + u.ID.Hex())
    w.WriteHeader(http.StatusCreated)
}
