package accessory

import (
  "github.com/brutella/hc/accessory"
  "github.com/brutella/hc/service"
  "github.com/kidoman/embd"
  _ "github.com/kidoman/embd/host/rpi"
  "log"
  "time"
)

type DigitalSwitch struct {
  *accessory.Accessory
  Switch *service.Switch
  Pin string
  Inverted bool
}

func NewDigitalSwitch (info Info, pin string, inverted bool) *DigitalSwitch {
  accInfo := accessory.Info{
    Name:         info.Name,
    Manufacturer: info.Manufacturer,
    Model:        info.Model,
    SerialNumber: info.SerialNumber,
  }

  acc := DigitalSwitch{}
  acc.Accessory = accessory.New(accInfo, accessory.TypeOutlet)
  acc.Pin = pin
  acc.Inverted = inverted
  acc.Switch = service.NewSwitch()
  acc.AddService(acc.Switch.Service)

  acc.Accessory.OnIdentify(acc.OnIdentify)
  acc.Switch.On.OnValueRemoteUpdate(acc.SetOn)

  return &acc
}

func (this *DigitalSwitch) OnIdentify() {
  timeout := 1 * time.Second
  this.SetOn(true)
  time.Sleep(timeout)
  this.SetOn(false)
}

func (this *DigitalSwitch) SetOn (on bool) {
  var value int

  if on == this.Inverted {
    value = embd.Low
  } else {
    value = embd.High
  }

  if err := this.Write(value); err != nil {
    log.Fatal(err)
  }
}

func (this *DigitalSwitch) Write (value int) (err error) {
  if err = embd.InitGPIO(); err != nil {
    return
  }

  if err = embd.SetDirection(this.Pin, embd.Out); err != nil {
    return
  }

  err = embd.DigitalWrite(this.Pin, value)
  return
}
