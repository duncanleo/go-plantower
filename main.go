package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/duncanleo/go-plantower/devices"
)

func main() {
	device := flag.String("device", "/dev/ttyAMA0", "name of the serial device. e.g. COM1 on Windows, /dev/ttyAMA0 on Linux")
	model := flag.String("model", "pms5003", "model name of the device")
	waitTime := flag.Int("wait", 2, "time to wait before getting reading from sensor device")
	listMode := flag.Bool("l", false, "list devices supported")

	flag.Parse()

	if *listMode {
		for k := range devices.DeviceFuncs {
			fmt.Println(k)
		}
		return
	}

	data, err := devices.DeviceFuncs[*model](*device, map[string]interface{}{
		"waitTime": *waitTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	json, _ := json.Marshal(data)
	fmt.Println(string(json))
}
