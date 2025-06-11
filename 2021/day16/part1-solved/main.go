package day

func main(hex string) int {
	println("Hexadecimal input:", hex)
	t := transmission{
		binary:  convertHexToBin(hex),
		parent:  0,
		packets: []packet{},
	}
	for t.binary != "" {
		t.parseNextPacket()
	}
	versionSum := 0
	for _, p := range t.packets {
		versionSum += p.version
	}
	println("Version sum:", versionSum)
	println()
	return versionSum
}
