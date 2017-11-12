package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "handlers/gpio"
)

func main () {
  r := mux.NewRouter()
  r.HandleFunc("/gpio/{id}", gpio.Show).Methods("GET")
  // r.HandleFunc("/gpio/{id}", gpio.Update).Methods("PUT")
  r.HandleFunc("/gpio/{id}/{value}", gpio.Update).Methods("GET")
  http.ListenAndServe(":8000", r)
}
