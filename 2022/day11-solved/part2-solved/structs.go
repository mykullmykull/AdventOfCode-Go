package day

import (
	"fmt"
	"goAoc2022/utils"
)

type monkeys struct {
	monkeyMap map[int]monkey
}

type monkey struct {
	items            []item
	op               operation
	divisor          int
	modMonkeyId      int
	notMonkeyId      int
	countInspections int
}

type item struct {
	start int
	ops   []operation
}

type operation struct {
	v1       string
	operator string
	v2       string
}

func (ms monkeys) takeTurn(id int) monkeys {
	thrower := ms.monkeyMap[id]
	for _, i := range thrower.items {
		item, receiverId := thrower.getReceiverId(i)
		receiver := ms.monkeyMap[receiverId]
		receiver.items = append(receiver.items, item)
		ms.monkeyMap[receiverId] = receiver
		thrower.countInspections++
	}
	thrower.items = []item{}
	ms.monkeyMap[id] = thrower
	return ms
}

func (o operation) getNewWorry(worry int) int {
	v1 := worry
	if o.v1 != "old" {
		v1 = utils.StrToInt(o.v1)
	}
	v2 := worry
	if o.v2 != "old" {
		v2 = utils.StrToInt(o.v2)
	}

	switch o.operator {
	case "+":
		return v1 + v2
	case "*":
		return v1 * v2
	default:
		panic(fmt.Sprintf("unrecognised operator %s\n", o.operator))
	}
}

func (m monkey) getReceiverId(i item) (item, int) {
	worry := i.start
	i.ops = append(i.ops, m.op)
	for _, op := range i.ops {
		worry = op.getNewWorry(worry) % m.divisor
	}

	if worry%m.divisor == 0 {
		return i, m.modMonkeyId
	}
	return i, m.notMonkeyId
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

func (ms monkeys) printMonkeys(allFields bool) {
	if allFields {
		for id := 0; id < len(ms.monkeyMap); id++ {
			println()
			println("id", id)
			ms.monkeyMap[id].printMonkey()
		}
		return
	}

	for id := 0; id < len(ms.monkeyMap); id++ {
		println("monkey", id, "inspected items", ms.monkeyMap[id].countInspections, "times")
	}
}

func (m monkey) printMonkey() {
	println("items", fmt.Sprintf("%v", m.items))
	println("op", fmt.Sprintf("%v", m.op))
	println("testDivisor", m.divisor)
	println("trueMonkeyId", m.modMonkeyId)
	println("falseMonkeyId", m.notMonkeyId)
	println("countInspections", m.countInspections)
}
