package day7

import (
	"fmt"
)

var cardOrderWithJokers = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
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

func runB(input []string) int {
	hands := parseHands(input)
	hands = orderHandsByRankWithJokers(hands)
	return payBids(hands)
}

func orderHandsByRankWithJokers(hands []hand) []hand {
	ordered := []hand{hands[0]}
	for i := 1; i < len(hands); i++ {
		thisHand := hands[i]
		inserted := false
		for j, thatHand := range ordered {
			if thatBeatsThisWithJokers(thatHand, thisHand) {
				ordered = append(ordered[:j+1], ordered[j:]...)
				ordered[j] = thisHand
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

func thatBeatsThisWithJokers(thatHand hand, thisHand hand) bool {
	thatRank := rankWithJokers(thatHand.cards)
	thisRank := rankWithJokers(thisHand.cards)
	if thatRank != thisRank {
		return thatRank > thisRank
	}

	for i := 0; i < len(thatHand.cards); i++ {
		thatValue := cardOrderWithJokers[thatHand.cards[i]]
		thisValue := cardOrderWithJokers[thisHand.cards[i]]
		if thatValue != thisValue {
			return thatValue > thisValue
		}
	}
	panic(fmt.Sprintf("Couldn't determine a winner between that %v and this %v", thatHand.cards, thisHand.cards))
}

func rankWithJokers(cards []string) float64 {
	cardCounts := map[string]float64{}
	rank := 0.0
	for _, c := range cards {
		cardCounts[c] += 1
		// Jokers don't count toward the base rank
		if c != "J" && cardCounts[c] > rank {
			rank = cardCounts[c]
		}
	}

	jokers := cardCounts["J"]
	// 4 or 5 of a kind
	if rank+jokers > 3 {
		return rank + jokers
	}

	triples := 0
	pairs := 0
	for _, v := range cardCounts {
		if v == 3 {
			triples++
		}
		if v == 2 {
			pairs++
		}
	}

	// Full House
	if (triples == 1 && pairs == 1) ||
		// (triples == 1 && jokers > 0) this would be caught previously since it could make the hand 4 or 5 of a kind
		(pairs == 2 && jokers == 1) {
		return 3.5
	}

	if rank+jokers == 3 {
		return 3
	}

	// Two Pair
	if pairs == 2 || jokers == 2 {
		return 2.5
	}

	if jokers == 1 {
		return 2
	}

	return rank
}
