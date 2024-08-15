package day

func main(input []string) int {
	sumOfSignalStrength := 0
	times := []int{20, 60, 100, 140, 180, 220}
	c := cpu{
		cycle:            0,
		x:                1,
		instructionIndex: 0,
	}
	for _, time := range times {
		c = c.advanceToTime(time, input)
		signalStrength := c.x * time
		// println(time, ":", signalStrength)
		sumOfSignalStrength += signalStrength
	}
	return sumOfSignalStrength
}
