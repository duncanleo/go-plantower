package devices

import (
	"errors"
	"log"
	"time"

	"github.com/tarm/serial"
)

func pms5003(device string, opts map[string]interface{}) (Data, error) {
	var result Data
	var err error
	var waitTime = 2

	c := &serial.Config{Name: device, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return result, err
	}

	if wt, ok := opts["waitTime"]; ok {
		waitTime = wt.(int)
	}

	// Set to passive mode
	_, err = s.Write(pltCreateChecksumByteArray(
		[]byte{
			0x42,
			0x4d,
			0xe1, // Change mode
			0x00,
			0x00, // Passive,
		},
	))
	if err != nil {
		return result, err
	}

	// Wake it up
	_, err = s.Write(pltCreateChecksumByteArray(
		[]byte{
			0x42,
			0x4d,
			0xe4, // Sleep set
			0x00,
			0x01, // Wake
		},
	))
	if err != nil {
		return result, err
	}

	// Wait some time
	time.Sleep(time.Duration(waitTime) * time.Second)

	// Flush any extra data, possibly left from active mode
	err = s.Flush()
	if err != nil {
		return result, err
	}

	// Read command
	_, err = s.Write(pltCreateChecksumByteArray(
		[]byte{
			0x42,
			0x4d,
			0xe2, // Read
			0x00,
			0x00,
		},
	))
	if err != nil {
		return result, err
	}

	var singleReadingBuf []byte

	numRead := 0
	for numRead < 32 {
		buf := make([]byte, 32)
		n, err := s.Read(buf)
		if err != nil {
			return result, err
		}
		numRead += n
		singleReadingBuf = append(singleReadingBuf, buf[:n]...)
	}
	if numRead == 32 {
		// We have a complete frame!
		result.CF.PM1 = int(singleReadingBuf[5])
		result.CF.PM25 = int(singleReadingBuf[7])
		result.CF.PM10 = int(singleReadingBuf[9])
		result.Atmospheric.PM1 = int(singleReadingBuf[11])
		result.Atmospheric.PM25 = int(singleReadingBuf[13])
		result.Atmospheric.PM10 = int(singleReadingBuf[15])
		result.ConcUnit = string(singleReadingBuf[17])
		result.ParticleCount.PointThree = int(singleReadingBuf[19])
		result.ParticleCount.PointFive = int(singleReadingBuf[21])
		result.ParticleCount.One = int(singleReadingBuf[23])
		result.ParticleCount.TwoPointFive = int(singleReadingBuf[25])
		result.ParticleCount.Five = int(singleReadingBuf[27])
		result.ParticleCount.Ten = int(singleReadingBuf[29])

		var (
			checkSumHigh = singleReadingBuf[30]
			checkSumLow  = singleReadingBuf[31]
		)

		var checkSum int
		for j := 0; j <= 29; j++ {
			checkSum += int(singleReadingBuf[j])
		}

		if hlBytesToInt(checkSumHigh, checkSumLow) != checkSum { // Checksum mismatch
			return result, errors.New("Checksum mismatch")
		}

		numRead = 0
		singleReadingBuf = make([]byte, 0)
	} else if numRead > 32 {
		log.Println("Weird, we're not supposed to have a frame > 32 bytes. Ignoring.")
		numRead = 0
		singleReadingBuf = make([]byte, 0)
	}

	// Sleep
	_, err = s.Write(pltCreateChecksumByteArray(
		[]byte{
			0x42,
			0x4d,
			0xe4, // Sleep set
			0x00,
			0x00, // Sleep
		},
	))

	return result, err
}
