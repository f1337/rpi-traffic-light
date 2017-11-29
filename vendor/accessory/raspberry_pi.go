package accessory

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "github.com/imdario/mergo"
  "os"
)

type RaspberryPi struct {
  Bridge        *accessory.Accessory
  Accessories []*accessory.Accessory
}

// create a new RaspberryPi bridge accessory
func NewRaspberryPi (info Info) *RaspberryPi {
  bridgeInfo := accessory.Info{
    Name:         "RaspberryPi",
    Manufacturer: "Raspberry Pi Foundation",
    Model:        "Raspberry Pi",
  }
  mergo.MergeWithOverwrite(&bridgeInfo, info)

  acc := RaspberryPi{}
  acc.Bridge = accessory.New(bridgeInfo, accessory.TypeBridge)

  return &acc
}

// setup a digital switch accessory, and add it to the Accessories array
func (this *RaspberryPi) AddDigitalSwitch (info Info, pin string, inverted bool) {
  // TODO: return an error on dupe pins
  acc := NewDigitalSwitch(info, pin, inverted)
  this.Accessories = append(this.Accessories, acc.Accessory)
}

// start the HAP server
func (this *RaspberryPi) Start (config Config) (err error) {
  serverConfig := hc.Config{
    IP:          config.IP,
    Pin:         config.Pin,
    Port:        config.Port,
    StoragePath: config.StoragePath,
  }

  t, err := hc.NewIPTransport(serverConfig, this.Bridge, this.Accessories...)
  if err != nil {
    return
  }

  hc.OnTermination(func() {
    t.Stop()
    // TODO: close GPIO connection
    os.Exit(0)
  })

  t.Start()
  return
}
