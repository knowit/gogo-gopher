package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/authors/{name}", func(writer http.ResponseWriter, request *http.Request) {
        vars := mux.Vars(request)
        name := vars["name"]

        fmt.Fprintf(writer, "Name of the author: %s\n", name)
    })

    http.ListenAndServe(":8080", router)
}