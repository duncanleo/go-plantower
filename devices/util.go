package devices

// intToHlBytes Convert integer to High and Low bytes
func intToHlBytes(n int) (byte, byte) {
	return byte(n>>8) & 0xff, byte(n & 0xff)
}

// hlBytesToInt Convert High and Low bytes to integer
func hlBytesToInt(h byte, l byte) int {
	return int(uint64(l) | uint64(h)<<8)
}

// pltCreateChecksumByteArray (Plantower) Create a new byte array from an existing one and append its checksum values.
func pltCreateChecksumByteArray(ba []byte) []byte {
	var total int
	for index := 0; index < len(ba); index++ {
		total += int(ba[index])
	}
	h, l := intToHlBytes(total)
	return append(ba, h, l)
}
