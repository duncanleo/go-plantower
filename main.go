package main

import (
	"encoding/json"
	"log"

	"github.com/duncanleo/go-plantower/devices"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyAMA0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	data, err := devices.DeviceFuncs["pms5003"](s, map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
	json, _ := json.Marshal(data)
	log.Println(string(json))
}
