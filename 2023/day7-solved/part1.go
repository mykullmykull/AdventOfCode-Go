package day7

import (
	"fmt"
	"goAoc2023/utils"
	"strings"
)

type hand struct {
	cards []string
	bid   int
}

var cardOrder = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func runA(input []string) int {
	hands := parseHands(input)
	hands = orderHandsByRank(hands)
	return payBids(hands)
}

func parseHands(input []string) []hand {
	hands := make([]hand, len(input))
	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], " ")
		cards := strings.Split(split[0], "")
		bid := split[1]
		hands[i] = hand{cards: cards, bid: utils.Atoi(bid)}
	}
	return hands
}

func orderHandsByRank(hands []hand) []hand {
	ordered := []hand{hands[0]}
	for i := 1; i < len(hands); i++ {
		thisHand := hands[i]
		inserted := false
		for j, thatHand := range ordered {
			// fmt.Printf("i: %d, this: %v, j: %d, that: %v\n", i, thisHand, j, thatHand)
			if thatBeatsThis(thatHand, thisHand) {
				// fmt.Println("  that beat this")
				ordered = append(ordered[:j+1], ordered[j:]...)
				ordered[j] = thisHand
				// fmt.Println(ordered)
				inserted = true
				break
			}
		}
		if !inserted {
			ordered = append(ordered, thisHand)
		}
	}
	return ordered
}

func thatBeatsThis(thatHand hand, thisHand hand) bool {
	thatRank := rank(thatHand.cards)
	thisRank := rank(thisHand.cards)
	if thatRank != thisRank {
		return thatRank > thisRank
	}

	for i := 0; i < len(thatHand.cards); i++ {
		thatValue := cardOrder[thatHand.cards[i]]
		thisValue := cardOrder[thisHand.cards[i]]
		if thatValue != thisValue {
			return thatValue > thisValue
		}
	}
	panic(fmt.Sprintf("Couldn't determine a winner between that %v and this %v", thatHand.cards, thisHand.cards))
}

func rank(cards []string) float64 {
	cardCounts := map[string]float64{}
	rank := 0.0
	for _, c := range cards {
		cardCounts[c] += 1
		if cardCounts[c] > rank {
			rank = cardCounts[c]
		}
	}

	threeOfAKinds := 0
	twoOfAKinds := 0
	for _, v := range cardCounts {
		if v == 3 {
			threeOfAKinds++
		}
		if v == 2 {
			twoOfAKinds++
		}
	}

	// Full House
	if threeOfAKinds == 1 && twoOfAKinds == 1 {
		rank = 3.5
	}

	// Two Pair
	if twoOfAKinds == 2 {
		rank = 2.5
	}

	return rank
}

func payBids(hands []hand) int {
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	return total
}
