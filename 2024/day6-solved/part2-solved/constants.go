package day

const (
	u    = "0"
	ur   = "1"
	urd  = "2"
	urdl = "3"
	url  = "4"
	ud   = "5"
	udl  = "6"
	ul   = "7"
	r    = "8"
	rd   = "9"
	rdl  = "A"
	rl   = "B"
	d    = "C"
	dl   = "D"
	l    = "E"
	up   = 0
	rt   = 1
	dn   = 2
	lt   = 3
)

var arrows = map[int]string{
	up: "^",
	rt: ">",
	dn: "v",
	lt: "<",
}

var directions = map[string][]int{
	".": {},
	"0": {up},
	"1": {up, rt},
	"2": {up, rt, dn},
	"3": {up, rt, dn, lt},
	"4": {up, rt, lt},
	"5": {up, dn},
	"6": {up, dn, lt},
	"7": {up, lt},
	"8": {rt},
	"9": {rt, dn},
	"A": {rt, dn, lt},
	"B": {rt, lt},
	"C": {dn},
	"D": {dn, lt},
	"E": {lt},
}
