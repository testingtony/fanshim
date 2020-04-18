package main

import (
	"fmt"
	"time"

	rpi "github.com/nathan-osman/go-rpigpio"
)

func main() {

	p, err := rpi.OpenPin(18, rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	for {
		fmt.Println("High")
		// set the pin to high (on)
		p.Write(rpi.HIGH)
		time.Sleep(5 * time.Second)
		fmt.Println("Low")
		// set the pin to low (off)
		p.Write(rpi.LOW)
		time.Sleep(5 * time.Second)

	}
	// /sys/class/thermal/thermal_zone0/temp  divide by 1000

}
