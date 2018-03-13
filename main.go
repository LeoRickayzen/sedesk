package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("home get"))
}

func homePostHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("home post"))
}

func main() {
    r := mux.NewRouter()
    
    //get request
    r.HandleFunc("/home", homeGetHandler).Methods("GET")

    //post request
    r.HandleFunc("/home", homePostHandler).Methods("POST")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}