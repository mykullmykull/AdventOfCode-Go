package day

import (
	"fmt"
	"goAoc2022/utils"
)

type monkeys struct {
	monkeyMap map[int]monkey
}

type monkey struct {
	items            []int
	op               operation
	testDivisor      int
	trueMonkeyId     int
	falseMonkeyId    int
	countInspections int
}

type operation struct {
	v1       string
	operator string
	v2       string
}

func (ms monkeys) takeTurn(id int) monkeys {
	thrower := ms.monkeyMap[id]
	for _, item := range thrower.items {
		newWorry := thrower.getNewWorry(item)
		receiverId := thrower.getReceiverId(newWorry)
		receiver := ms.monkeyMap[receiverId]
		receiver.items = append(receiver.items, newWorry)
		ms.monkeyMap[receiverId] = receiver
		thrower.countInspections++
	}
	thrower.items = []int{}
	ms.monkeyMap[id] = thrower
	return ms
}

func (m monkey) getNewWorry(item int) int {
	v1 := item
	if m.op.v1 != "old" {
		v1 = utils.StrToInt(m.op.v1)
	}
	v2 := item
	if m.op.v2 != "old" {
		v2 = utils.StrToInt(m.op.v2)
	}

	switch m.op.operator {
	case "+":
		return (v1 + v2) / 3
	case "*":
		return (v1 * v2) / 3
	default:
		panic(fmt.Sprintf("unrecognised operator %s\n", m.op.operator))
	}
}

func (m monkey) getReceiverId(worry int) int {
	if worry%m.testDivisor == 0 {
		return m.trueMonkeyId
	}
	return m.falseMonkeyId
}

func (ms monkeys) getMonkeyBusiness() int {
	mostActive := 0
	nextMostActive := 0
	for _, m := range ms.monkeyMap {
		if m.countInspections > nextMostActive {
			nextMostActive = m.countInspections
		}
		if m.countInspections > mostActive {
			nextMostActive = mostActive
			mostActive = m.countInspections
		}
	}
	return mostActive * nextMostActive
}

func (ms monkeys) printMonkeys() {
	for id := 0; id < len(ms.monkeyMap); id++ {
		println()
		println("id", id)
		ms.monkeyMap[id].printMonkey()
	}
}

func (ms monkeys) printItems() {
	for id := 0; id < len(ms.monkeyMap); id++ {
		fmt.Print("id ", id, " items ")
		for _, item := range ms.monkeyMap[id].items {
			fmt.Print(item, " ")
		}
		println()
	}
}

func (m monkey) printMonkey() {
	println("items", fmt.Sprintf("%v", m.items))
	println("op", fmt.Sprintf("%v", m.op))
	println("testDivisor", m.testDivisor)
	println("trueMonkeyId", m.trueMonkeyId)
	println("falseMonkeyId", m.falseMonkeyId)
	println("countInspections", m.countInspections)
}
