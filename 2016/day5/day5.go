package day5

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func runB(input string) string {
	code := "________"
	i := -1
	for {
		i++
		str := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(str))
		hex := hex.EncodeToString(hash[:])
		if hex[0:5] == "00000" {
			positionByte := hex[5]
			positionStr := string(positionByte)
			position, err := strconv.Atoi(positionStr)
			if err != nil || position >= len(code) || code[position] != '_' {
				continue
			}

			code = code[:position] + string(hex[6]) + code[position+1:]
		}
		if isFinished(code) {
			break
		}
	}
	return code
}

func isFinished(code string) bool {
	for _, c := range code {
		if c == '_' {
			return false
		}
	}
	return true
}

func runA(input string) string {
	code := ""
	i := 0
	for {
		str := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(str))
		hex := hex.EncodeToString(hash[:])
		if hex[0:5] == "00000" {
			code += string(hex[5])
		}
		if len(code) == 8 {
			break
		}
		i++
	}
	return code
}
