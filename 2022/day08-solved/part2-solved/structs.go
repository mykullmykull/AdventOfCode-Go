package day

import (
	"goAoc2022/utils"
)

type pointer struct {
	visibleTrees map[int]bool
	maxHeight    int
	r            int
	c            int
}

func (p pointer) resetMax() pointer {
	p.maxHeight = -1
	return p
}

func (p pointer) getTreeId() int {
	return p.r*1000 + p.c
}

func (p pointer) getHeight(input []string) int {
	return utils.StrToInt(string(input[p.r][p.c]))
}

func (p pointer) addToVisibleTrees(input []string) pointer {
	treeHeight := p.getHeight(input)
	if treeHeight <= p.maxHeight {
		return p
	}
	p.maxHeight = treeHeight
	p.visibleTrees[p.getTreeId()] = true
	return p
}

func (p pointer) resetAll(base pointer) pointer {
	p.visibleTrees = map[int]bool{}
	p.maxHeight = base.maxHeight
	p.r = base.r
	p.c = base.c
	return p
}

func (p pointer) getScenicScore(input []string) int {
	up := p.resetAll(p)
	up.r--
	for up.r >= 0 {
		up.visibleTrees[up.getTreeId()] = true
		if up.getHeight(input) >= p.maxHeight {
			break
		}
		up.r--
	}

	rt := p.resetAll(p)
	rt.c++
	for rt.c < len(input[0]) {
		rt.visibleTrees[rt.getTreeId()] = true
		if rt.getHeight(input) >= p.maxHeight {
			break
		}
		rt.c++
	}

	dn := p.resetAll(p)
	dn.r++
	for dn.r < len(input) {
		dn.visibleTrees[dn.getTreeId()] = true
		if dn.getHeight(input) >= p.maxHeight {
			break
		}
		dn.r++
	}

	lt := p.resetAll(p)
	lt.c--
	for lt.c >= 0 {
		lt.visibleTrees[lt.getTreeId()] = true
		if lt.getHeight(input) >= p.maxHeight {
			break
		}
		lt.c--
	}

	return len(up.visibleTrees) *
		len(rt.visibleTrees) *
		len(dn.visibleTrees) *
		len(lt.visibleTrees)
}

func (p pointer) addShorterTrees(input []string) pointer {
	treeHeight := p.getHeight(input)
	if treeHeight >= p.maxHeight {
		return p
	}
	p.visibleTrees[p.getTreeId()] = true
	return p
}
