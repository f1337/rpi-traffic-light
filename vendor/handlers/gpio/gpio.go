package gpio

import (
  gpiox "adapters/embdx"
  "encoding/json"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
)

type Pin struct {
  Id    string `json:"id"`
  Value string `json:"value"`
}

func Show (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  pinId := params["id"]
  pinValue, err := gpiox.ReadPinValue(pinId)
  if err != nil {
    WriteJSON(w, err)
    return
  }
  pinString := strconv.Itoa(pinValue)

  pin := &Pin{
    Id:    pinId,
    Value: pinString}
  WriteJSON(w, pin)
}

func Update (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  pinId := params["id"]
  pinValue, err := strconv.Atoi(params["value"])
  if err != nil {
    pinValue = 0
  }
  if err := gpiox.WritePinValue(pinId, pinValue); err != nil {
    WriteJSON(w, err)
    return
  }
  pinString := strconv.Itoa(pinValue)

  pin := &Pin{
    Id:    pinId,
    Value: pinString}
  WriteJSON(w, pin)
}

func WriteJSON (w http.ResponseWriter, v interface{}) {
  j, _ := json.Marshal(v)
  w.Write(j)
}
