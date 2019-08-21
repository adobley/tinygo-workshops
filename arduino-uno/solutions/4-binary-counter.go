package main

import (
	"machine"
	"time"
)

const (
	firstRedPin  = 11
	secondRedPin = 10
	thirdRedPin  = 9
)

func main() {
	firstRed := machine.Pin(firstRedPin)
	firstRed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	secondRed := machine.Pin(secondRedPin)
	secondRed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	thirdRed := machine.Pin(thirdRedPin)
	thirdRed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		firstRed.Low()
		secondRed.Low()
		thirdRed.Low()
		time.Sleep(time.Millisecond * 500)

		firstRed.High()
		secondRed.Low()
		thirdRed.Low()
		time.Sleep(time.Millisecond * 500)

		firstRed.Low()
		secondRed.High()
		thirdRed.Low()
		time.Sleep(time.Millisecond * 500)

		firstRed.High()
		secondRed.High()
		thirdRed.Low()
		time.Sleep(time.Millisecond * 500)

		firstRed.Low()
		secondRed.Low()
		thirdRed.High()
		time.Sleep(time.Millisecond * 500)

		firstRed.High()
		secondRed.Low()
		thirdRed.High()
		time.Sleep(time.Millisecond * 500)

		firstRed.Low()
		secondRed.High()
		thirdRed.High()
		time.Sleep(time.Millisecond * 500)

		firstRed.High()
		secondRed.High()
		thirdRed.High()
		time.Sleep(time.Millisecond * 500)
	}
}
