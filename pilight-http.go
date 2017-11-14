package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "handlers/gpio"
)

func main () {
  r := mux.NewRouter()

  // GPIO API
  r.HandleFunc("/gpio/{id}", gpio.Show).Methods("GET")
  // r.HandleFunc("/gpio/{id}", gpio.Update).Methods("PUT")
  r.HandleFunc("/gpio/{id}/{value}", gpio.Update).Methods("GET")

  // serve static files
  fs := http.FileServer(http.Dir("public"))
  r.PathPrefix("/").Handler(fs)
  http.Handle("/", r)

  http.ListenAndServe(":8000", r)
}
