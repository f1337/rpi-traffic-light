package embdx

import (
  "github.com/kidoman/embd"
  _ "github.com/kidoman/embd/host/rpi"
)

func ReadPinValue (pinNum string) (pinValue int, err error) {
  if err = embd.InitGPIO(); err != nil {
    return
  }

  pinValue, err = embd.DigitalRead(pinNum)
  return
}

func WritePinValue (pinNum string, pinValue int) (err error) {
  if err = embd.InitGPIO(); err != nil {
    return
  }

  if err = embd.SetDirection(pinNum, embd.Out); err != nil {
    return
  }

  if pinValue > 0 {
    pinValue = embd.High
  } else {
    pinValue = embd.Low
  }

  err = embd.DigitalWrite(pinNum, pinValue)
  return
}
