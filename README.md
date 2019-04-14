# go-plantower
This is a command-line utility to read data off a Plantower sensor (serial-based). See releases for pre-built binaries.

### Usage
```shell
$ go-plantower -h
Usage of go-plantower:
  -device string
    	name of the serial device. e.g. COM1 on Windows, /dev/ttyAMA0 on Linux (default "/dev/ttyAMA0")
  -l	list devices supported
  -model string
    	model name of the device (default "pms5003")
  -wait int
    	time to wait before getting reading from sensor device (default 2)
```

### Currently supported devices
- PMS5003
