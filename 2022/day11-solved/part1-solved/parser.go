package day

import (
	"fmt"
	"goAoc2022/utils"
	"strings"
)

func parse(input []string) monkeys {
	monkeyMap := map[int]monkey{}
	m := monkey{}
	id := 0
	for _, line := range input {
		if line == "" {
			monkeyMap[id] = m
			m = monkey{}
			id++
			continue
		}

		if line[:6] == "Monkey" {
			continue
		}

		itemsPrefix := "Starting items: "
		if line[:len(itemsPrefix)] == itemsPrefix {
			items := strings.Split(strings.Split(line, itemsPrefix)[1], ", ")
			for _, item := range items {
				m.items = append(m.items, utils.StrToInt(item))
			}
			continue
		}

		opsPrefix := "Operation: new = "
		if line[:len(opsPrefix)] == opsPrefix {
			split := strings.Split(strings.Split(line, opsPrefix)[1], " ")
			m.op.v1 = split[0]
			m.op.operator = split[1]
			m.op.v2 = split[2]
			continue
		}

		testPrefix := "Test: divisible by "
		if line[:len(testPrefix)] == testPrefix {
			divisibles := strings.Split(strings.Split(line, testPrefix)[1], " ")
			if len(divisibles) > 1 {
				panic(fmt.Sprintf("Found %d test divisibles when there should be only 1. %s", len(divisibles), line))
			}
			m.testDivisor = utils.StrToInt(divisibles[0])
			continue
		}

		truePrefix := "If true: throw to monkey "
		if line[:len(truePrefix)] == truePrefix {
			ids := strings.Split(strings.Split(line, truePrefix)[1], " ")
			if len(ids) > 1 {
				panic(fmt.Sprintf("Found %d true ids when there should be only 1. %s", len(ids), line))
			}
			m.trueMonkeyId = utils.StrToInt(ids[0])
			continue
		}

		falsePrefix := "If false: throw to monkey "
		if line[:len(falsePrefix)] == falsePrefix {
			ids := strings.Split(strings.Split(line, falsePrefix)[1], " ")
			if len(ids) > 1 {
				panic(fmt.Sprintf("Found %d false ids when there should be only 1. %s", len(ids), line))
			}
			m.falseMonkeyId = utils.StrToInt(ids[0])
			continue
		}

		panic(fmt.Sprintf("unrecognised line: %s", line))
	}
	return monkeys{monkeyMap}
}
