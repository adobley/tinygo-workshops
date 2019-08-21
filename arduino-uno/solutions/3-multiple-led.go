package main

import (
	"machine"
	"time"
)

const (
	firstRedPin  = 11
	secondRedPin = 10
)

func main() {
	firstRed := machine.Pin(firstRedPin)
	firstRed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	secondRed := machine.Pin(secondRedPin)
	secondRed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		firstRed.High()
		secondRed.Low()
		time.Sleep(time.Millisecond * 500)

		firstRed.Low()
		secondRed.High()
		time.Sleep(time.Millisecond * 500)
	}
}
