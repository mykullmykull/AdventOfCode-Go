package day

type transmission struct {
	binary  string
	parent  int
	packets []packet
}

func (t *transmission) parseNextPacket() {
	if len(t.binary) < 6 || // not enough bits for version and typeId{
		t.isFinishedCount() {
		t.binary = ""
		t.getGrandparent()
		return
	}
	version := convertBinToInt(t.binary[:3])
	typeId := convertBinToInt(t.binary[3:6])
	t.binary = t.binary[6:]

	p := packet{
		parent:  t.parent,
		version: version,
	}
	t.packets = append(t.packets, p)

	if typeId == literal {
		p.typeId = literal
		// Remaining binary depends on how many groups are contained in the literal value
		// 5 bits per group, last group starts with 0
		t.binary = p.parseLiteral(t.binary)
		return
	}

	if t.binary == "" {
		t.getGrandparent()
		return
	}
	lengthTypeId := convertBinToInt(t.binary[:1])
	t.binary = t.binary[1:]
	if lengthTypeId == bitsLength {
		p.typeId = bitsLength
		if len(t.binary) < 15 {
			t.getGrandparent()
			return
		}
		p.parseBitsLength(t.binary[0:15])
		// Remaining bits will always be what's left after the 15-bit-long packet
		t.binary = t.binary[15:]
		return
	}

	if lengthTypeId == countLength {
		p.typeId = countLength
		if len(t.binary) < 15 {
			t.getGrandparent()
			return
		}
		p.count = convertBinToInt(t.binary[0:11])
		// Remaining bits will always be what's left after the 11-bit-long children count
		t.binary = t.binary[11:]
		t.parent = version
	}
}

func (t *transmission) isFinishedCount() bool {
	if len(t.packets) == 0 || t.parent < 0 || t.parent >= len(t.packets) {
		return false
	}
	p := t.packets[t.parent]
	if p.typeId == countLength && len(p.children) == p.count {
		return true
	}
	return false
}

func (t *transmission) getGrandparent() {
	p := t.packets[t.parent]
	t.parent = p.parent
}
