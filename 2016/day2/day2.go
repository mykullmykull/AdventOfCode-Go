package day2

import (
	"fmt"
)

type button struct {
	name string
	up   string
	dn   string
	lt   string
	rt   string
}

var b1 = &button{"1", "", "3", "", ""}
var b2 = &button{"2", "", "6", "", "3"}
var b3 = &button{"3", "1", "7", "2", "4"}
var b4 = &button{"4", "", "8", "3", ""}
var b5 = &button{"5", "", "", "", "6"}
var b6 = &button{"6", "2", "a", "5", "7"}
var b7 = &button{"7", "3", "b", "6", "7"}
var b8 = &button{"8", "4", "c", "7", "9"}
var b9 = &button{"9", "", "", "8", ""}
var ba = &button{"A", "6", "", "", "b"}
var bb = &button{"B", "7", "d", "a", "c"}
var bc = &button{"C", "8", "", "b", ""}
var bd = &button{"D", "b", "", "", ""}
var buttons = map[string]*button{
	"1": b1,
	"2": b2,
	"3": b3,
	"4": b4,
	"5": b5,
	"6": b6,
	"7": b7,
	"8": b8,
	"9": b9,
	"a": ba,
	"b": bb,
	"c": bc,
	"d": bd,
}

func (b *button) move(direction byte) *button {
	nextName := ""
	switch direction {
	case 'U':
		nextName = b.up
	case 'D':
		nextName = b.dn
	case 'L':
		nextName = b.lt
	case 'R':
		nextName = b.rt
	default:
		panic(fmt.Sprintf("invalid direction: %c", direction))
	}
	if nextName == "" {
		return b
	}

	return buttons[nextName]
}

func runB(input []string) string {
	code := ""
	button := b5
	for _, row := range input {
		for i := 0; i < len(row); i++ {
			direction := row[i]
			button = button.move(direction)
		}
		code += button.name
	}
	return code
}

func runA(input []string) string {
	code := ""
	num := 5
	for _, row := range input {
		for i := 0; i < len(row); i++ {
			char := row[i]
			switch char {
			case 'U':
				if num > 3 {
					num = num - 3
				}
			case 'D':
				if num < 7 {
					num = num + 3
				}
			case 'L':
				if num != 1 && num != 4 && num != 7 {
					num = num - 1
				}
			case 'R':
				if num != 3 && num != 6 && num != 9 {
					num = num + 1
				}
			}
		}
		code += fmt.Sprintf("%d", num)
	}
	return code
}
