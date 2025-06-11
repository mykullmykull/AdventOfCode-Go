package day

var hexToBin = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func convertHexToBin(hex string) string {
	bin := ""
	for _, char := range hex {
		if val, exists := hexToBin[char]; exists {
			bin += val
		} else {
			panic("Invalid hex character: " + string(char))
		}
	}
	return bin
}

func convertBinToInt(bin string) int {
	result := 0
	for _, char := range bin {
		result <<= 1 // Shift left by 1 (equivalent to multiplying by 2)
		if char == '1' {
			result += 1 // Add 1 if the character is '1'
		}
	}
	return result
}
