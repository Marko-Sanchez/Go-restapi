package handlers

import (
    "net/http"
)

// Handles root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Asset Not Found\n"))
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to Marko's Rest API\n"))
}


