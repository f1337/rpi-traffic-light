package accessory

import (
  "github.com/brutella/hc/accessory"
  "github.com/imdario/mergo"
)

type RaspberryPi struct {
  Bridge       *accessory.Accessory
  Accessories []accessory.Accessory
}

func NewRaspberryPi (info accessory.Info, accessories ...*accessory.Accessory) *RaspberryPi {
  acc := RaspberryPi{}
  defaultInfo := accessory.Info{
    Name:         "RaspberryPi",
    Manufacturer: "Raspberry Pi Foundation",
    Model:        "Raspberry Pi",
  }
  mergo.Merge(&info, defaultInfo)
  acc.Bridge = accessory.New(info, accessory.TypeBridge)
  // acc.Accessories = accessories

  return &acc
}

// func (this *RaspberryPi) Start (config Config) (err error) {
//   ipConfig := hc.Config{Pin: "86753091", StoragePath: "krakatoa"}
//   t, err := hc.NewIPTransport(ipConfig, this.Accessory, this.Switches)
//   if err != nil {
//     return
//   }
//
//   hc.OnTermination(func() {
//     log.Println("stopping...")
//     t.Stop()
//     log.Println("stopped, exiting")
//     os.Exit(0)
//   })
//
//   t.Start()
// }
