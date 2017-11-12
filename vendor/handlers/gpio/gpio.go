package gpio

import (
  gpiox "adapters/embdx"
  "io"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
)

func Show (w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  pin := params["id"]
  pinValue, err := gpiox.ReadPinValue(pin)
  if err != nil {
    io.WriteString(w, "Error: " + err.Error())
  }
  pinString := strconv.Itoa(pinValue)
  io.WriteString(w, "GPIO #" + pin + " value: " + pinString)
}

func Update (w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  pin := params["id"]
  pinValue, err := strconv.Atoi(params["value"])
  if err != nil {
    pinValue = 0
  }
  if err := gpiox.WritePinValue(pin, pinValue); err != nil {
    io.WriteString(w, "Error: " + err.Error())
  }
  pinString := strconv.Itoa(pinValue)
  io.WriteString(w, "GPIO #" + pin + " set value: " + pinString)
}
