package main

import (
    "Go-restapi/handlers"
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/users/", handlers.UsersRouter)
    http.HandleFunc("/users", handlers.UsersRouter)
    http.HandleFunc("/", handlers.RootHandler)

    err := http.ListenAndServe("localhost:1111", nil)
    if err!= nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
