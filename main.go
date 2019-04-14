package main

import (
	"log"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyAMA0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	var singleReadingBuf []byte

	numRead := 0
	for {
		buf := make([]byte, 32)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		numRead += n
		singleReadingBuf = append(singleReadingBuf, buf[:n]...)
		if numRead == 32 {
			// We have a complete frame!
			var (
				cfPM001           = int(singleReadingBuf[5])
				cfPM025           = int(singleReadingBuf[7])
				cfPM100           = int(singleReadingBuf[9])
				atmosPM1          = int(singleReadingBuf[11])
				atmosPM25         = int(singleReadingBuf[13])
				atmosPM10         = int(singleReadingBuf[15])
				concUnit          = string(singleReadingBuf[17])
				numParticles003um = int(singleReadingBuf[19])
				numParticles005um = int(singleReadingBuf[21])
				numParticles010um = int(singleReadingBuf[23])
				numParticles025um = int(singleReadingBuf[25])
				numParticles050um = int(singleReadingBuf[27])
				numParticles100um = int(singleReadingBuf[29])
				checkSumHigh      = singleReadingBuf[30]
				checkSumLow       = singleReadingBuf[31]
			)

			log.Printf(
				"[DATA] CF=[PM1=%d PM2.5=%d PM10=%d] ATMOS=[PM1=%d PM2.5=%d PM10=%d] UNIT=%s NUM_PARTICLES=[0.3um=%d 0.5um=%d 1.0um=%d 2.5um=%d 5.0um=%d 10um=%d] CHECK=[High=%q Low=%q]\n",
				cfPM001,
				cfPM025,
				cfPM100,
				atmosPM1,
				atmosPM25,
				atmosPM10,
				concUnit,
				numParticles003um,
				numParticles005um,
				numParticles010um,
				numParticles025um,
				numParticles050um,
				numParticles100um,
				checkSumHigh,
				checkSumLow,
			)

			numRead = 0
			singleReadingBuf = make([]byte, 0)
		} else if numRead > 32 {
			log.Println("Weird, we're not supposed to have a frame > 32 bytes. Ignoring.")
			numRead = 0
			singleReadingBuf = make([]byte, 0)
		}
	}
}
