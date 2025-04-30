package day

import (
	"fmt"
	"goAoc2022/utils"
)

func (s state) toString() string {
	return fmt.Sprintf("Time: %d, Materials: %v, Robots: %v", s.time, s.materials, s.robots)
}

func (s state) branchToString() string {
	return utils.JoinArrayOfInts(s.branch[:], "")
}
