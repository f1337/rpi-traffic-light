package gpio

import (
  "io"
  "net/http"
  "github.com/gorilla/mux"
)

func Show (w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  io.WriteString(w, "Hello gpio #" + id + "!")
  // gobot vs embd?
  // embd.InitGPIO()
  // defer embd.CloseGPIO()
  // pin, err := embd.NewDigitalPin(id)
  // handle err?
  // pinValue := pin.Read()
}

func Update (w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]
  io.WriteString(w, "Hello gpio #" + id + "!")
  // gobot vs embd?
  // embd.InitGPIO()
  // defer embd.CloseGPIO()
  // pin, err := embd.NewDigitalPin(id)
  // handle err?
  // pin.SetDirection(embd.Out)
  // pin.Write(emdb.High)
}
