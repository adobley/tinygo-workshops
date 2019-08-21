package main

import (
	"machine"
	"time"
)

const (
	ledPin    = 11
	buttonPin = 2
)

func main() {
	led := machine.Pin(ledPin)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button := machine.Pin(buttonPin)
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if button.Get() {
			led.High()
		} else {
			led.Low()
		}
		time.Sleep(time.Millisecond * 10)
	}
}
