package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	rpi "github.com/nathan-osman/go-rpigpio"
)

func main() {

	p, err := rpi.OpenPin(18, rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	onTemp := getVar("ON_TEMP", 65)
	offTemp := getVar("OFF_TEMP", 55)

	fmt.Printf("Turning off at %d and on at %d\n", offTemp, onTemp)

	for {
		temp := getTemp()
		//fmt.Println(temp)

		if temp >= onTemp {
			p.Write(rpi.HIGH)
		}
		if temp <= offTemp {
			p.Write(rpi.LOW)
		}

		time.Sleep(5 * time.Second)
	}

	// /sys/class/thermal/thermal_zone0/temp  divide by 1000

}

func getTemp() int64 {
	dat, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		panic(err)
	}

	temp, err := strconv.ParseInt(strings.TrimSpace(string(dat)), 10, 64)
	if err != nil {
		panic(err)
	}

	return temp / 1000
}

func getVar(varname string, def int64) int64 {

	val, ok := os.LookupEnv(varname)
	if !ok {
		return def
	}

	ival, err := strconv.ParseInt(strings.TrimSpace(string(val)), 10, 64)
	if err != nil {
		panic(err)
	}
	return ival
}
