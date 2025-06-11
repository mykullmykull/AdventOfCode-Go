package day

const (
	literal     = 4
	bitsLength  = 0
	countLength = 1
)

type packet struct {
	parent   int
	version  int
	typeId   int
	count    int
	value    int
	children []int
}

func (p *packet) parseLiteral(binary string) string {
	stillParsing := true
	bin := ""
	for stillParsing {
		if len(binary) < 5 { // not enough bits for a literal value
			panic("Not enough bits for a literal value")
		}
		group := binary[:5]
		binary = binary[5:]
		if group[0] == '0' {
			stillParsing = false
		}
		bin += group[1:5]
	}
	p.value = convertBinToInt(bin)
	return binary
}

func (p *packet) parseBitsLength(binary string) {
}
