package day

// ------------------------------ inputs ------------------------------------//
var testInput = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}
var expected = 31

var realInput = []string{
	"abccccccaaccaaccccaaaaacccccaaaaccccccccccccccccccccccccccccccccaaaaaaaaaaaaaaaaaaaccccccccccccccccaaaccccccccccccaacccccccccccccccccccccccccccccccccccccccccaaaa",
	"abaaaaccaaaaaccccaaaaaccccccaaaaccccccccccccccccccccaaacccccccccccaaaaaaaaaaaaaaaaaaccccccccccccccccaaaaccccccaaacaaccccccccccccccccccccccccccccccccccccccccaaaaa",
	"abaaacccaaaaaaaacaaaaaacccccaaaaccccccccccccccccccccaaaaacccccccccaaaaaaaaaaaaaaaaacccccccccaaccccaaaaaacccccccaaaaaccccaaccccccccccccccacccccccccccccccccccaaaaa",
	"abaaacccccaaaaaccccaaaaccccccaaacccccccccccccccccccaaaaaccccccccccaaaaaacacaaaaaacccccccccccaaccccaaaaacccccccccaaaaaaccaaaaaaccccccccccaaaccccacccccccccccaaaaaa",
	"abaacccccaaaaaccccaacccccccccccccccaaaaacccccccccccaaaaacccccccccaaaaaaaaccaaaaaaacccccccaaaaaaaaccaaaaacccccccaaaaaaaccaaaaacccccccccccaaacccaaaccccccccccccccaa",
	"abaaacccaaacaacccccccccccccccccccccaaaaaccccccccccccaaaaacccccccaaaaaaaaaccaaccaaacccccccaaaaaaaaccaacccccccccaaaaaaccaaaaaaccccccccccccaaaacaaaaccccccccccccccaa",
	"abaaacccccccaaccccccccccccccccccccaaaaaaccccccccccccaaccccccaacccaaaccaaaaccccccaacccccccccaaaacccccccccccccccaacaaaccaaaaaaaccccccccccccajjjjjjjcccccccccccccccc",
	"abcaacccccccccccccccccccccccccccccaaaaaaccccccccccccccccccccaaaaccccccaaaaccccccccccccccaacaaaaaccccccccccccccccccaaccccaaaaaacccccccccccjjjjjjjjjcccccaaaccccccc",
	"abccccccccccccccccccccccccccccccccaaaaaaccaaccccccccccccccaaaaaacccccccaaacccccccccccaacaaaaaaaaccccccccccccccccccccccccaaccaaccccccccaiijjjjojjjjcccccaaacaccccc",
	"abcccccccccccccccccccccccaaacccccccaaacacaaacccccccccccccccaaaaccccaaccccccccccccccccaaaaaaacccaccccccccccccccccccccccccaacccccccccccaiiijjooooojjkccaaaaaaaacccc",
	"abccccccccccccccccccccccaaaaccccccccccaaaaaccccccccccccccccaaaaacccaaaaaccccccccccccccaaaaaacccccccccccccccccccccccccccccccccccccciiiiiiiioooooookkkcaaaaaaaacccc",
	"abccccccccccccccccccccccaaaaccccccccccaaaaaaaacccccccccccccaacaaccaaaaacccccccaaacccaaaaaaaaccccccccccccccccccccccccccccccccccchiiiiiiiiooooouoookkkccaaaaaaccccc",
	"abcccccccccaaccccccccccccaaaccccccccccccaaaaacccccccccccccccccccccaaaaaccccccaaaacccaaaaacaacccccccccccccaacaacccccccccccccccchhhiiiinnnooouuuuoookkkccaaaaaccccc",
	"abcccccccccaaacccccccccccccccccccccccccaaaaacccccccccccccccccccccccaaaaacccccaaaaccccccaaccccccccccccccccaaaaacccccccccccccccchhhnnnnnnnnouuuuuuppkkkkaaaaaaccccc",
	"abccccccaaaaaaaacccaaccccccccccccccccccaacaaccaacaaccccccccccccccccaacccccccccaaaccccccaacccccccccccccccaaaaacccccccccccccccchhhnnnnnnnnntuuxuuupppkkkkkacccccccc",
	"abccccccaaaaaaaacacaaaacccccccccccccccccccaaccaaaaacccccccccccccccccccccccccccccccccccccccccccccccccccccaaaaaacccccccaacccccchhhnnnnttttttuxxxuuppppkkkkkcccccccc",
	"abcccccccaaaaaaccaaaaaaccccccccccaaccccccccccaaaaaccccccccccccccccccccccccaaacccccccccccccccccccccccccccccaaaaccaaccaaacccaaahhhnnntttttttuxxxxuupppppllllccccccc",
	"abcccccccaaaaaacccaaaacccccccccaaaaaaccccccccaaaaaacccccccccccccccccccccccaaacccccccccccccccccccccccccccccacccccaaaaaaacaaaaahhhppntttxxxxxxxxuuuuvpppplllccccccc",
	"abcccccccaaaaaacccaaaacccccccccaaaaaacccccaaaaaaaaaccccccccccccccccccccaaaaaaaacccccccccccccccccccccaaaccccccaacaaaaaaccaaaaahhhpppttxxxxxxxxyuuvvvvvppplllcccccc",
	"abcccccccaaccaacccaacaccaaaaccccaaaaacccccaaaaaaaaaccccccccccccccccccccaaaaaaaacccccccccccccccccccccaaacaaaaaaaccaaaaaaaaaaaaahhppptttxxxxxxyyyyyyvvvppplllcccccc",
	"SbccccccccccccccccccccccaaaacccaaaaacccccccaaaaaaaaacaaaccccccccaacccccccaaaaaccccccccaaaaacccccccccaaaaaaaaaaaaaaaaaaaaaaaaacgggpppttxxxxEzzyyyyyvvvqqqlllcccccc",
	"abccccccccccccccccccccccaaaacccaaaaacccccccaaaaaaaaccaaaccccccccaaacaaccaaaaaaccccccccaaaaacccccccaaaaaaaaaaaaaaaaaaaaaaaaaaacgggpppsssxxxyyyyyyvvvvvqqlllccccccc",
	"abcccaaaccccccccccccccccaaaccccccccccccccccaaaaaaaaaaaaaccccccccaaaaaaccaaaaaacccccccaaaaaacccaaaccaaaaaccaaaaaaaaaaaacccccccccgggppssswwyyyyyyvvvvqqqqlllccccccc",
	"abcaaaaaccccccccccccccccccccccccccccccccccaaaaaaaaaaaaacccccccaaaaaaacccaccaaacccccccaaaaaacccaaacccaaaaaaaaaaaccccaaacccaaaaacgggppsswwwyyyyyyvvqqqqqlllcccccccc",
	"abcaaaaaaccccccccccccccccccccccccccccccccccaaccaaaaaaaaaaaccccaaaaaaacccccccccccccccccaaaaacccaaacaaaacaaaaaaaaccccaaacccaaaaacggpppsswwwywwyyyvvqqqmmmlccccccccc",
	"abcaaaaaacccccccaacaaccccccccccccccccccccccccccaaaaaaaaaaaccccccaaaaacccccccccccccccccaaaccaaaaaaaaaaacccccccaacccccccccaaaaaacggpppsswwwwwwwwyvvqqqmmmcccccccccc",
	"abcaaaaaccccccccaaaaaccccccccccccccccccccccccccccaaaaaaaacccccccaacaaacccccccccccccccccccccaaaaaaaaaccccccccccccccccccccaaaaaagggoossswwwwrrwwwvvqqmmmccccccccccc",
	"abcaaaaacccccccaaaaaccccccccccccccccccccccccccccaaaaaaacccccccccaaccccccccccccccccccccccccccaaaaaaacccccccccccaaaccccccccaaaaagggooosssssrrrrwwwvqqmmmcccaacccccc",
	"abcccccccccccccaaaaaaccccccccccccccccccccaacccccccccaaaccccccccccccccccccccccccccccccccccccccaaaaaaccccccccccccaaaaccccccaaaccgggooosssssrrrrrwwrrqmmmcccaacccccc",
	"abcccccccccccccccaaaacccccccccccccccccccaaaacccccccacaaacccccccccccccccccccccccccccccccccccccaaaaaaacccccccccaaaaaacccccccccccgffoooooosoonrrrrrrrrmmmccaaaaacccc",
	"abcccccccccccccccaccccccccccccccccccccccaaaacccccccaaaaacccccccccccccccccccccccccccccccccccccaaacaaacccccccccaaaaacccccccccccccfffoooooooonnnrrrrrmmmddcaaaaacccc",
	"abccccccccccccccccccccccccccccccccccccccaaaaccccccccaaaaacccccccccccccccccccccccccaaaccccccccaacccccccccccccccaaaaaccccccccccccffffoooooonnnnnnrnnmmmdddaaaaacccc",
	"abcccccccccccccccccccccccccccccccccccccccccccccccccaaaaaacccccccccccccccccaaaaaccaaaacccccccccccccccccccccccccaacccccccccccccccfffffffffeeeennnnnnmmdddaaaacccccc",
	"abcccccccaaaccccccccaccccccccccccccccccccccccccccccaaaaccccccccccccaaaccccaaaaaccaaaaccccccccccccccccccccccccccccccccccccccccccccfffffffeeeeennnnnmddddaaaaaccccc",
	"abcccaaccaaacccccaaaacccccaacccccccccccccccccccccccccaaacccccccccccaaacccaaaaaacccaaaccccccccccccccccccccccccccccccccccccccccccccccffffeeeeeeeedddddddcccaacccccc",
	"abcccaaaaaaacccccaaaaaaccaaacccccccccccccccccccccccccccacccccccccccaaaaccaaaaaaccccccccccccccccccccccccccaacccccccccaaaccccccccccccccaaaaaaeeeeedddddcccccccccccc",
	"abcccaaaaaacccccccaaaacccaaacaaaccccaaaacccccccccaaacaaaccccccaacccaaaacaaaaaaacccccccccccccccccccccccccaaaccccccccaaaacccccccccccccccccccaaaaeeddddccccccccccccc",
	"abccccaaaaaaaacccaaaaaaaaaaaaaaaccccaaaacccccccccaaaaaaacccccaaacccaaaaaaaaaacccccccccccccccccccccccaaacaaaccccccccaaaaccccccccccccccccccccaaaccccccccccccccaaaca",
	"abcccaaaaaaaaacccaacaaaaaaaaaaacccccaaaaccccccccccaaaaaacaaacaaacaaaaaaaaaacccccccccccccccaaacccccccaaaaaaaaaaccccccaaaccccccccccccccccccccaacccccccccccccccaaaaa",
	"abcccaaaaaaaacccccccccccaaaaaacccccccaacccccccccccaaaaaaaaaaaaaaaaaaaaaaaaccccccccccccccccaaaacccccccaaaaaaaaacccccccccccccccccccccccccccccaaacccccccccccccccaaaa",
	"abccaaaaaaacccccccccccccaaaaaaacccccccccccccccccaaaaaaaaaaaaaaaaaaaaaaaaaaacccccccccccccccaaaacccccccaaaaaaaacccccccccccccccccccccccccccccccccccccccccccccccaaaaa",
}
