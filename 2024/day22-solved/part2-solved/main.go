package day

import (
	"fmt"
	"strconv"
)

type monkey struct {
	start   int
	prices  []int
	changes []int
}

type seq struct {
	change1 int
	change2 int
	change3 int
	chagne4 int
}

func main(input []int, count int) int {
	monkeys := parseMonkeys(input, count)
	testedSeqs := map[seq]int{}
	for x, m := range monkeys {
		seqsFound := map[seq]bool{}
		for y := 3; y < len(m.changes); y++ {
			fmt.Printf("monkey: %d/%d, step: %d/2000\n", x+1, len(monkeys), y)
			s, p := m.getSeqPrice(y)
			if _, ok := seqsFound[s]; ok {
				continue
			}
			testedSeqs[s] += p
			seqsFound[s] = true
		}
	}
	mostBananas := -1
	bestSeq := seq{-9, -9, -9, -9}
	for s, b := range testedSeqs {
		if b > mostBananas {
			mostBananas = b
			bestSeq = s
		}
	}
	fmt.Printf("best sequence: %v\n", bestSeq)
	return mostBananas
}

func parseMonkeys(input []int, count int) []monkey {
	monkeys := []monkey{}
	for _, secret := range input {
		prices, changes := getPricesAndChanges(secret, count)
		if len(prices) != count || len(changes) != count {
			panic(fmt.Sprintf("count %d and prices %d and changes %d should have the same length", count, len(prices), len(changes)))
		}
		monkeys = append(monkeys, monkey{secret, prices, changes})
	}
	return monkeys
}

func getPrice(secret int) int {
	str := fmt.Sprintf("%d", secret)
	lastChar := str[len(str)-1:]
	price, err := strconv.Atoi(lastChar)
	if err != nil {
		panic("can't convert last char to int: " + lastChar)
	}
	return price
}

func getPricesAndChanges(secret int, count int) ([]int, []int) {
	prices := []int{}
	changes := []int{}
	p0 := getPrice(secret)
	for i := 0; i < count; i++ {
		secret = evolve(secret, 1)
		p := getPrice(secret)
		c := p - p0
		prices = append(prices, p)
		changes = append(changes, c)
		p0 = p
	}
	return prices, changes
}

func evolve(start int, count int) int {
	for count > 0 {
		start = prune(mix(start, start*64))
		start = prune(mix(start, start/32))
		start = prune(mix(start, start*2048))
		count--
	}
	return start
}

func prune(n int) int {
	return n % 16777216
}

func mix(a int, b int) int {
	return a ^ b
}

func (m monkey) getSeqPrice(x int) (seq, int) {
	s := seq{m.changes[x-3], m.changes[x-2], m.changes[x-1], m.changes[x]}
	p := m.prices[x]
	return s, p
}

func buyBananasFromMonkey(m monkey, s seq) int {
	for x := 3; x < len(m.changes); x++ {
		lastSeq, p := m.getSeqPrice(x)
		if lastSeq == s {
			return p
		}
	}
	return 0
}

func buyBananas(monkeys []monkey, s seq) int {
	bananas := 0
	for _, m := range monkeys {
		bananas += buyBananasFromMonkey(m, s)
	}
	return bananas
}
