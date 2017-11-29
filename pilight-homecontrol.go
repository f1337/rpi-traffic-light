package main

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  rpi "accessory"
  "log"
  "os"
)

func main () {
  // TODO: hide the hc interface
  // trafficLight := rpi.NewRaspberryPi(accessory.Info{
  //   Name:         "TrafficLight",
  //   Manufacturer: "brighid ronan and dad",
  // }, red, yellow, green)
  // if err := trafficLight.Start(hc.Config{Pin: "86753091", StoragePath: "krakatoa"}), err != nil {
  //   log.Fatal(err)
  // }

  // TODO: read rpiConfig, per-light config, ipConfig from local .yml or .json
  // cf. https://github.com/timoschilling/dashbridge/blob/master/dashbridge.go
  red := rpi.NewDigitalSwitch(accessory.Info{Name: "Red"}, "14", true).Accessory
  yellow := rpi.NewDigitalSwitch(accessory.Info{Name: "Yellow"}, "15", true).Accessory
  green := rpi.NewDigitalSwitch(accessory.Info{Name: "Green"}, "18", true).Accessory

  trafficLight := rpi.NewRaspberryPi(accessory.Info{
    Name:         "TrafficLight",
    Manufacturer: "brighid ronan and dad",
  }).Bridge

  ipConfig := hc.Config{Pin: "86753091", StoragePath: "krakatoa"}
  t, err := hc.NewIPTransport(ipConfig, trafficLight, red, yellow, green)
  if err != nil {
    log.Fatal(err)
  }

  hc.OnTermination(func() {
    log.Println("stopping...")
    t.Stop()
    log.Println("stopped, exiting")
    os.Exit(0)
  })

  t.Start()
}
