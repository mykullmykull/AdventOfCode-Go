package day

const (
	up = iota
	dn
	lt
	rt
	wall
	open
	robot
	boxLt
	boxRt
)

func toConst(d byte) int {
	byteMap := map[byte]int{
		'^': up,
		'v': dn,
		'<': lt,
		'>': rt,
		'#': wall,
		'.': open,
		'@': robot,
		'[': boxLt,
		']': boxRt,
	}
	return byteMap[d]
}

func toString(d int) string {
	intMap := map[int]string{
		up:    "^",
		dn:    "v",
		lt:    "<",
		rt:    ">",
		wall:  "#",
		open:  ".",
		robot: "@",
		boxLt: "[",
		boxRt: "]",
	}
	return intMap[d]
}
