package day16

type direction string

const (
	up direction = "up"
	lt direction = "lt"
	dn direction = "dn"
	rt direction = "rt"
)

type point struct {
	row int
	col int
}

type beam struct {
	location  point
	direction direction
}

type reflector string

const (
	empty     reflector = "."
	vSplitter reflector = "|"
	hSplitter reflector = "-"
	aMirror   reflector = "/"
	bMirror   reflector = "\\"
)

type gridPoint struct {
	uBeam     bool
	lBeam     bool
	dBeam     bool
	rBeam     bool
	reflector reflector
}

func (b beam) advance(gridMap map[point]gridPoint) []beam {
	switch b.direction {
	case up:
		b.location.row -= 1
	case lt:
		b.location.col -= 1
	case dn:
		b.location.row += 1
	case rt:
		b.location.col += 1
	}

	newBeams := []beam{}
	switch gridMap[b.location].reflector {
	case empty:
		newBeams = []beam{b}
	case vSplitter:
		newBeams = b.splitV()
	case hSplitter:
		newBeams = b.splitH()
	case aMirror:
		newBeams = b.reflectA()
	case bMirror:
		newBeams = b.reflectB()
	}

	validBeams := []beam{}
	for _, nb := range newBeams {
		if nb.isInside(gridMap) && nb.isNew(gridMap[nb.location]) {
			validBeams = append(validBeams, nb)
		}
	}

	return validBeams
}

func (b beam) splitV() []beam {
	if b.direction == up || b.direction == dn {
		return []beam{b}
	}
	return []beam{
		{b.location, up},
		{b.location, dn},
	}
}

func (b beam) splitH() []beam {
	if b.direction == lt || b.direction == rt {
		return []beam{b}
	}
	return []beam{
		{b.location, lt},
		{b.location, rt},
	}
}

func (b beam) reflectA() []beam {
	switch b.direction {
	case up:
		return []beam{{b.location, rt}}
	case lt:
		return []beam{{b.location, dn}}
	case dn:
		return []beam{{b.location, lt}}
	case rt:
		return []beam{{b.location, up}}
	}
	return []beam{}
}

func (b beam) reflectB() []beam {
	switch b.direction {
	case up:
		return []beam{{b.location, lt}}
	case lt:
		return []beam{{b.location, up}}
	case dn:
		return []beam{{b.location, rt}}
	case rt:
		return []beam{{b.location, dn}}
	}
	return []beam{}
}

func (b beam) isInside(gridMap map[point]gridPoint) bool {
	_, ok := gridMap[b.location]
	return ok
}

func (b beam) isNew(gp gridPoint) bool {
	switch b.direction {
	case up:
		return !gp.uBeam
	case lt:
		return !gp.lBeam
	case dn:
		return !gp.dBeam
	case rt:
		return !gp.rBeam
	}
	return false
}

func (gp gridPoint) addBeams(beams []beam) gridPoint {
	for _, b := range beams {
		switch b.direction {
		case up:
			gp.uBeam = true
		case lt:
			gp.lBeam = true
		case dn:
			gp.dBeam = true
		case rt:
			gp.rBeam = true
		}
	}
	return gp
}
