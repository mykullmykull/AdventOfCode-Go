package day

import "strconv"

func main(input []string) int64 {
	gStr, eStr := "", ""
	for x := 0; x < len(input[0]); x++ {
		count0s, count1s := 0, 0
		for _, line := range input {
			if line[x] == '0' {
				count0s++
			} else {
				count1s++
			}
		}
		if count0s > count1s {
			gStr += "0"
			eStr += "1"
		} else {
			gStr += "1"
			eStr += "0"
		}
	}
	g, gOk := strconv.ParseInt(gStr, 2, 64)
	e, eOk := strconv.ParseInt(eStr, 2, 64)

	if gOk != nil || eOk != nil {
		panic("Failed to convert binary to decimal")
	}

	return g * e
}
