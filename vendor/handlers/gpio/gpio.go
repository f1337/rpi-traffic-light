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
  params := mux.Vars(r)
  pinId := params["id"]
  pinValue, err := gpiox.ReadPinValue(pinId)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  pinString := strconv.Itoa(pinValue)

  pin := &Pin{
    Id:    pinId,
    Value: pinString}
  respondWithJSON(w, http.StatusOK, pin)
}

func Update (w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  pinId := params["id"]

  var pin Pin
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&pin); err != nil {
      respondWithError(w, http.StatusBadRequest, "Unable to decode request JSON.")
      return
  }
  defer r.Body.Close()
  pin.Id = pinId

  pinValue, err := strconv.Atoi(pin.Value)
  if err != nil {
    pinValue = 0
  }
  if err := gpiox.WritePinValue(pinId, pinValue); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, pin)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
