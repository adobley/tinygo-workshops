package main

import (
	"machine"
	"time"
)

func main() {
	red := machine.Pin(11)
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		red.Low()
		time.Sleep(time.Millisecond * 500)

		red.High()
		time.Sleep(time.Millisecond * 500)
	}
}
