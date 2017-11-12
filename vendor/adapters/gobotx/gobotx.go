package embd

import (
  // "gobot.io/x/gobot/drivers/gpio"
  "gobot.io/x/gobot/platforms/raspi"
)

// https://github.com/hybridgroup/gobot/blob/master/platforms/raspi/raspi_adaptor.go#L164
func ReadPinValue (pinNum string) (pinValue int, err error) {
  r := raspi.NewAdaptor()

  pinValue, err = r.DigitalRead(pinNum)
  return
}

// https://github.com/hybridgroup/gobot/blob/master/platforms/raspi/raspi_adaptor.go#L173
func WritePinValue (pinNum string, pinValue int) (err error) {
  r := raspi.NewAdaptor()

  pinByte := byte(0)
  if pinValue > 0 {
    pinByte = byte(1)
  }

  err = r.DigitalWrite(pinNum, pinByte)
  return
}
