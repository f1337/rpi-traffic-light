package main

import (
  rpi "accessory"
  "log"
)

func main () {
  // TODO: read rpiInfo, per-light config, rpiConfig from .json file
  // cf. https://github.com/timoschilling/dashbridge/blob/master/dashbridge.go
  trafficLight := rpi.NewRaspberryPi(rpi.Info{
    Name:         "TrafficLight",
    Manufacturer: "brighid ronan and dad",
  })

  trafficLight.AddDigitalSwitch(rpi.Info{Name: "Red"}, "14", true)
  trafficLight.AddDigitalSwitch(rpi.Info{Name: "Yellow"}, "15", true)
  trafficLight.AddDigitalSwitch(rpi.Info{Name: "Green"}, "18", true)

  if err := trafficLight.Start(rpi.Config{Pin: "86753091", StoragePath: "krakatoa"}); err != nil {
    log.Fatal(err)
  }

  // using hc directly:
  //
  // import (
  //   "github.com/brutella/hc"
  //   "github.com/brutella/hc/accessory"
  //   rpi "accessory"
  //   "log"
  // )
  //
  // trafficLight := rpi.NewRaspberryPi(accessory.Info{
  //   Name:         "TrafficLight",
  //   Manufacturer: "brighid ronan and dad",
  // }).Bridge
  //
  // red := rpi.NewDigitalSwitch(accessory.Info{Name: "Red"}, "14", true).Accessory
  // yellow := rpi.NewDigitalSwitch(accessory.Info{Name: "Yellow"}, "15", true).Accessory
  // green := rpi.NewDigitalSwitch(accessory.Info{Name: "Green"}, "18", true).Accessory
  //
  // ipConfig := hc.Config{Pin: "86753091", StoragePath: "krakatoa"}
  // t, err := hc.NewIPTransport(ipConfig, trafficLight, red, yellow, green)
  // if err != nil {
  //   log.Fatal(err)
  // }
  //
  // hc.OnTermination(func() {
  //   log.Println("stopping...")
  //   t.Stop()
  //   log.Println("stopped, exiting")
  //   os.Exit(0)
  // })
  //
  // t.Start()
}
