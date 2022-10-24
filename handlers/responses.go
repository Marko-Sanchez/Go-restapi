package handlers

import (
    "encoding/json"
    "net/http"
)

type jsonResponse map[string]interface{}

func postError(w http.ResponseWriter, code int) {
    http.Error(w, http.StatusText(code), code)
}

// makes a http reponse to client
func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse) {
    if content != nil {
        // convert content to a json byte sequence
        js, err := json.Marshal(content)
        if err != nil {
            postError(w, http.StatusInternalServerError)
            return
        }

        // return response to user
        w.Header().Set("Content-Type" , "aplication/json")
        w.WriteHeader(code)
        w.Write(js)
        return
    }

    // respond with error code
    w.WriteHeader(code)
    w.Write([]byte(http.StatusText(code)))
}
