package main

import (
  "github.com/f1337/hc"
  "github.com/brutella/hc/accessory"
  "log"
  "os"
  gpiox "adapters/embdx"
)

type DigitalOutput struct {
  accessory *accessory.Lightbulb
  pin string
}

func NewDigitalOutput (name string, pin string) *DigitalOutput {
  info := accessory.Info {
    Name:         name,
    Manufacturer: "brighid ronan and dad",
  }

  acc := accessory.NewLightbulb(info)

  light := &DigitalOutput{acc, pin}

  acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
    if on == true {
      turnLightOn(light)
    } else {
      turnLightOff(light)
    }
  })

  return light
}

func turnLightOn (light *DigitalOutput) {
  log.Println("Turn " + light.accessory.Info.Name.GetValue() + " On")
  if err := gpiox.WritePinValue(light.pin, 0); err != nil {
    log.Fatal(err)
  }
}

func turnLightOff (light *DigitalOutput) {
  log.Println("Turn " + light.accessory.Info.Name.GetValue() + " Off")
  if err := gpiox.WritePinValue(light.pin, 1); err != nil {
    log.Fatal(err)
  }
}

func main () {
  red := NewDigitalOutput("Red", "14")
  yellow := NewDigitalOutput("Yellow", "15")
  green := NewDigitalOutput("Green", "18")

  t, err := hc.NewIPTransport(hc.Config{Name: "TrafficLight", Pin: "32191123"}, red.accessory.Accessory, yellow.accessory.Accessory, green.accessory.Accessory)
  if err != nil {
    log.Fatal(err)
  }

  // TODO: does not KILL on ctrl-c
  hc.OnTermination(func() {
    log.Println("on term")
    t.Stop()
    log.Println("stopped, exiting")
    os.Exit(0)
  })

  t.Start()
}
