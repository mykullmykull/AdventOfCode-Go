package day

import (
	"fmt"
	"goAoc2021/utils"
	"strings"
)

func split(s string) string {
	digitStr := ""
	for _, r := range s {
		if isDigit(byte(r)) {
			digitStr += string(r)
			if len(digitStr) == 2 {
				splitStr := splitDigit(digitStr)
				return strings.Replace(s, digitStr, splitStr, 1)
			}
			continue
		}
		digitStr = ""
	}
	return s
}

func splitDigit(s string) string {
	n := utils.Atoi(s)
	lt := n / 2
	rt := (n + 1) / 2
	return fmt.Sprintf("[%d,%d]", lt, rt)
}
