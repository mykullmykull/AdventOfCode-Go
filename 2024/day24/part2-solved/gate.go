package day

import (
	"fmt"
	"strings"
)

type gate struct {
	a   string
	opp string
	b   string
}

func (g gate) isFullMatch(other gate) bool {
	if g.opp != other.opp {
		return false
	}
	ab := g.a + g.b
	return strings.Contains(ab, other.a) && strings.Contains(ab, other.b)
}

func (g gate) swapToMatch(other gate) (string, string, error) {
	if g.opp != other.opp || g.isFullMatch(other) {
		return "", "", fmt.Errorf("no swap found")
	}
	if g.a == other.a {
		return g.b, other.b, nil
	}
	if g.a == other.b {
		return g.b, other.a, nil
	}
	if g.b == other.a {
		return g.a, other.b, nil
	}
	if g.b == other.b {
		return g.a, other.a, nil
	}
	return "", "", fmt.Errorf("no match: %s %s", g.toString(), other.toString())
}

func (g gate) toString() string {
	return strings.Trim(g.a+" "+g.opp+" "+g.b, " ")
}

func (g *gate) swap(a, b string) {
	if g.a == a {
		g.a = b
	}
	if g.a == b {
		g.a = a
	}
	if g.b == a {
		g.b = b
	}
	if g.b == b {
		g.b = a
	}
}
