package accessory

import (
  "github.com/brutella/hc/accessory"
  "github.com/brutella/hc/service"
  "github.com/kidoman/embd"
  _ "github.com/kidoman/embd/host/rpi"
  "log"
)

type DigitalSwitch struct {
  *accessory.Accessory
  Switch *service.Switch
  Pin string
  Inverted bool
}

func (this *DigitalSwitch) OnValueRemoteUpdate (on bool) {
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

func NewDigitalSwitch (info accessory.Info, pin string, inverted bool) *DigitalSwitch {
  acc := DigitalSwitch{}
  acc.Accessory = accessory.New(info, accessory.TypeOutlet)
  acc.Pin = pin
  acc.Inverted = inverted
  acc.Switch = service.NewSwitch()
  acc.AddService(acc.Switch.Service)
  acc.Switch.On.OnValueRemoteUpdate(acc.OnValueRemoteUpdate)
  return &acc
}
