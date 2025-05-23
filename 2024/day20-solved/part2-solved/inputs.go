package day

// ------------------------------ inputs ------------------------------------//
var testInput = []string{
	"###############",
	"#...#...#.....#",
	"#.#.#.#.#.###.#",
	"#S#...#.#.#...#",
	"#######.#.#.###",
	"#######.#.#...#",
	"#######.#.###.#",
	"###..E#...#...#",
	"###.#######.###",
	"#...###...#...#",
	"#.#####.#.###.#",
	"#.#...#.#.#...#",
	"#.#.#.#.#.#.###",
	"#...#...#...###",
	"###############",
}
var expected = 44

var realInput = []string{
	"#############################################################################################################################################",
	"#.........#...#...###...###...#.............###...#...#...#.....#...#...#...###...#.....#.....#.....###.....#...#...#...#...#...........#...#",
	"#.#######.#.#.#.#.###.#.###.#.#.###########.###.#.#.#.#.#.#.###.#.#.#.#.#.#.###.#.#.###.#.###.#.###.###.###.#.#.#.#.#.#.#.#.#.#########.#.#.#",
	"#...#...#...#.#.#.#...#...#.#.#.........#...#...#.#.#.#.#...#...#.#.#.#...#...#.#...#...#...#.#.#...#...#...#.#.#.#.#.#...#...#.......#...#.#",
	"###.#.#.#####.#.#.#.#####.#.#.#########.#.###.###.#.#.#.#####.###.#.#.#######.#.#####.#####.#.#.#.###.###.###.#.#.#.#.#########.#####.#####.#",
	"###...#.....#.#.#.#.....#.#.#.#...###...#...#.#...#.#.#.....#.#...#.#.....#...#.....#.......#.#.#...#...#.....#...#.#.....#.....#...#.....#.#",
	"###########.#.#.#.#####.#.#.#.#.#.###.#####.#.#.###.#.#####.#.#.###.#####.#.#######.#########.#.###.###.###########.#####.#.#####.#.#####.#.#",
	"#...#...#...#.#.#.#...#.#.#.#.#.#...#.#.....#.#.###.#.#...#.#.#.#...###...#...#...#.........#.#.#...#...#...#.......#...#.#.......#...###...#",
	"#.#.#.#.#.###.#.#.#.#.#.#.#.#.#.###.#.#.#####.#.###.#.#.#.#.#.#.#.#####.#####.#.#.#########.#.#.#.###.###.#.#.#######.#.#.###########.#######",
	"#.#...#.#...#.#.#.#.#...#.#.#.#...#.#.#.###...#.#...#.#.#...#...#.#...#...#...#.#.#.......#.#...#...#.#...#...###.....#...#...#.....#.......#",
	"#.#####.###.#.#.#.#.#####.#.#.###.#.#.#.###.###.#.###.#.#########.#.#.###.#.###.#.#.#####.#.#######.#.#.#########.#########.#.#.###.#######.#",
	"#.#...#.....#.#.#.#.....#.#.#.#...#.#.#...#...#.#...#.#.........#.#.#...#.#.#...#.#...#...#.......#...#.........#.....#.....#...###.#...#...#",
	"#.#.#.#######.#.#.#####.#.#.#.#.###.#.###.###.#.###.#.#########.#.#.###.#.#.#.###.###.#.#########.#############.#####.#.###########.#.#.#.###",
	"#...#.......#...#.#.....#.#.#.#...#.#.#...#...#...#.#.###...#...#...#...#.#.#...#.#...#.#.........#.......#...#.#.....#.#.....#...#...#...###",
	"###########.#####.#.#####.#.#.###.#.#.#.###.#####.#.#.###.#.#.#######.###.#.###.#.#.###.#.#########.#####.#.#.#.#.#####.#.###.#.#.###########",
	"###...#...#.....#.#...#...#.#...#.#...#...#...#...#.#...#.#.#...#.....#...#.#...#.#.#...#.#.......#.#...#...#...#.#...#.#.#...#.#.....###...#",
	"###.#.#.#.#####.#.###.#.###.###.#.#######.###.#.###.###.#.#.###.#.#####.###.#.###.#.#.###.#.#####.#.#.#.#########.#.#.#.#.#.###.#####.###.#.#",
	"#...#...#.......#.....#...#...#.#...#...#...#.#...#...#.#.#...#.#.....#...#.#.#...#.#.###...#.....#...#...###...#...#.#...#.....#...#.....#.#",
	"#.#######################.###.#.###.#.#.###.#.###.###.#.#.###.#.#####.###.#.#.#.###.#.#######.###########.###.#.#####.###########.#.#######.#",
	"#.#.............#.......#.#...#...#...#.#...#.#...###.#.#...#.#.#.....###.#...#...#.#...#...#.....#.....#...#.#.###...#...........#.#.......#",
	"#.#.###########.#.#####.#.#.#####.#####.#.###.#.#####.#.###.#.#.#.#######.#######.#.###.#.#.#####.#.###.###.#.#.###.###.###########.#.#######",
	"#...#...........#.#...#.#...#...#.#...#.#.....#.......#.#...#...#...###...#.....#.#...#.#.#...#...#.###.....#.#.#...#...#.........#...#.....#",
	"#####.###########.#.#.#.#####.#.#.#.#.#.###############.#.#########.###.###.###.#.###.#.#.###.#.###.#########.#.#.###.###.#######.#####.###.#",
	"#...#...........#.#.#...#...#.#.#...#...#.........#.....#...#...#...#...#...###.#.#...#...#...#...#.#...#.....#.#...#...#.#.......#...#.#...#",
	"#.#.###########.#.#.#####.#.#.#.#########.#######.#.#######.#.#.#.###.###.#####.#.#.#######.#####.#.#.#.#.#####.###.###.#.#.#######.#.#.#.###",
	"#.#...........#...#.......#...#.#...#...#.......#.#.....###...#.#.#...#...#...#...#.....#...#...#.#...#.#.#...#...#...#...#.........#...#.###",
	"#.###########.#################.#.#.#.#.#######.#.#####.#######.#.#.###.###.#.#########.#.###.#.#.#####.#.#.#.###.###.###################.###",
	"#...........#.#...#...#...#...#...#...#.#.......#.......#...#...#...#...#...#...###...#.#.###.#...#.....#...#.#...#...#...#...#.....#...#...#",
	"###########.#.#.#.#.#.#.#.#.#.#########.#.###############.#.#.#######.###.#####.###.#.#.#.###.#####.#########.#.###.###.#.#.#.#.###.#.#.###.#",
	"#...........#...#...#...#...#...........#.....#...###...#.#.#...#.....###.#.....#...#...#...#.....#...#...###.#.###.#...#...#...###...#...#.#",
	"#.###########################################.#.#.###.#.#.#.###.#.#######.#.#####.#########.#####.###.#.#.###.#.###.#.###################.#.#",
	"#.........#.......#...#.................#.....#.#.#...#.#.#.#...#.....#...#...#...#.........#...#.#...#.#.#...#.....#.........#...#.....#...#",
	"#########.#.#####.#.#.#.###############.#.#####.#.#.###.#.#.#.#######.#.#####.#.###.#########.#.#.#.###.#.#.#################.#.#.#.###.#####",
	"#.........#...###...#.#.#.....#...#...#.#.#...#.#.#...#.#.#.#...#.....#...#...#...#.#...#...#.#.#.#.###.#.#.#.....#...........#.#.#...#.....#",
	"#.###########.#######.#.#.###.#.#.#.#.#.#.#.#.#.#.###.#.#.#.###.#.#######.#.#####.#.#.#.#.#.#.#.#.#.###.#.#.#.###.#.###########.#.###.#####.#",
	"#.............#...###...#...#...#...#...#.#.#...#.#...#.#.#.#...#.#.....#.#.#...#.#.#.#...#.#.#.#.#.#...#...#...#.#.............#.....#...#.#",
	"###############.#.#########.#############.#.#####.#.###.#.#.#.###.#.###.#.#.#.#.#.#.#.#####.#.#.#.#.#.#########.#.#####################.#.#.#",
	"#...............#.....#...#.....#.....###.#...###.#...#.#.#...###.#...#.#.#.#.#.#.#...#.....#.#.#.#...#.........#.....#...#...#...#.....#...#",
	"#.###################.#.#.#####.#.###.###.###.###.###.#.#.#######.###.#.#.#.#.#.#.#####.#####.#.#.#####.#############.#.#.#.#.#.#.#.#########",
	"#...#...#...#.......#...#...###...#...#...#...#...#...#...#.......#...#.#.#.#.#.#...#...#...#.#.#.#...#.............#.#.#.#.#...#.#.#...#...#",
	"###.#.#.#.#.#.#####.#######.#######.###.###.###.###.#######.#######.###.#.#.#.#.###.#.###.#.#.#.#.#.#.#############.#.#.#.#.#####.#.#.#.#.#.#",
	"###...#...#...#.....#.....#.........#...#...###...#.....###...#...#...#.#.#...#.#...#...#.#.#.#...#.#...#.......#...#...#.#.....#...#.#...#.#",
	"###############.#####.###.###########.###.#######.#####.#####.#.#.###.#.#.#####.#.#####.#.#.#.#####.###.#.#####.#.#######.#####.#####.#####.#",
	"#...............#...#.###...........#.....#.....#.#.....#.....#.#.#...#.#.....#.#...#...#.#.#.....#...#...#...#...#.....#.....#.#.....#.....#",
	"#.###############.#.#.#############.#######.###.#.#.#####.#####.#.#.###.#####.#.###.#.###.#.#####.###.#####.#.#####.###.#####.#.#.#####.#####",
	"#.........#.....#.#.#.#.............#.......#...#.#.....#.....#.#.#.###...#...#...#.#.#...#.#...#.#...#...#.#.......#...#...#...#.#.....#...#",
	"#########.#.###.#.#.#.#.#############.#######.###.#####.#####.#.#.#.#####.#.#####.#.#.#.###.#.#.#.#.###.#.#.#########.###.#.#####.#.#####.#.#",
	"#.......#...###...#...#...............#...#...###.#...#...#...#.#.#.....#.#.....#.#.#.#.###.#.#...#.....#...#...#.....#...#.......#.......#.#",
	"#.#####.###############################.#.#.#####.#.#.###.#.###.#.#####.#.#####.#.#.#.#.###.#.###############.#.#.#####.###################.#",
	"#.....#.#...#...#...#...#.......#.......#.#.....#...#.#...#.....#.#...#.#.#...#.#.#.#...#...#.#...............#.#.......#...#.....#.......#.#",
	"#####.#.#.#.#.#.#.#.#.#.#.#####.#.#######.#####.#####.#.#########.#.#.#.#.#.#.#.#.#.#####.###.#.###############.#########.#.#.###.#.#####.#.#",
	"#.....#...#.#.#.#.#.#.#.#.....#.#.......#.......#...#.#.........#.#.#...#.#.#.#.#...#...#...#.#...#...#.......#.....#.....#.#.#...#.....#...#",
	"#.#########.#.#.#.#.#.#.#####.#.#######.#########.#.#.#########.#.#.#####.#.#.#.#####.#.###.#.###.#.#.#.#####.#####.#.#####.#.#.#######.#####",
	"#.........#.#.#...#...#.......#.........#...#...#.#.#.#.....#...#...#...#...#.#...#...#.....#...#...#...#...#...###.#...###...#.#.....#...###",
	"#########.#.#.###########################.#.#.#.#.#.#.#.###.#.#######.#.#####.###.#.###########.#########.#.###.###.###.#######.#.###.###.###",
	"#.........#...#...#...#...#...............#...#...#.#...###.#.........#...#...#...#.............#...#.....#...#...#...#.....#...#...#...#...#",
	"#.#############.#.#.#.#.#.#.#######################.#######.#############.#.###.#################.#.#.#######.###.###.#####.#.#####.###.###.#",
	"#...#...#.....#.#.#.#...#...#.......................###...#.........#.....#...#.#.........#.....#.#...#...###.....###.......#.#...#...#.#...#",
	"###.#.#.#.###.#.#.#.#########.#########################.#.#########.#.#######.#.#.#######.#.###.#.#####.#.###################.#.#.###.#.#.###",
	"###...#...###...#...#...#...#...............###.........#.........#.#.......#.#.#...#...#.#...#.#.#.....#...................#.#.#...#.#.#...#",
	"#####################.#.#.#.###############.###.#################.#.#######.#.#.###.#.#.#.###.#.#.#.#######################.#.#.###.#.#.###.#",
	"#.........#.......#...#...#...............#.#...#.......#...#.....#.........#.#.#...#.#.#...#.#.#...#...#...#...#.....#.....#...###...#...#.#",
	"#.#######.#.#####.#.#####################.#.#.###.#####.#.#.#.###############.#.#.###.#.###.#.#.#####.#.#.#.#.#.#.###.#.#################.#.#",
	"#.......#.#.....#.#.....................#...#.#...#...#...#.#.........#.....#.#.#...#.#...#...#.###...#...#...#...###.#.......#...#...#...#.#",
	"#######.#.#####.#.#####################.#####.#.###.#.#####.#########.#.###.#.#.###.#.###.#####.###.#################.#######.#.#.#.#.#.###.#",
	"#.......#.....#.#.......................#...#.#.#...#...###.......#...#...#.#.#.###...###.#...#.....#.......#...#...#.........#.#.#.#.#...#.#",
	"#.###########.#.#########################.#.#.#.#.#####.#########.#.#####.#.#.#.#########.#.#.#######.#####.#.#.#.#.###########.#.#.#.###.#.#",
	"#...........#.#.#...#...#.....#...........#...#...#...#...#.....#...#...#.#.#...#...#...#...#.......#.#.....#.#...#.#.....#...#.#...#...#.#.#",
	"###########.#.#.#.#.#.#.#.###.#.###################.#.###.#.###.#####.#.#.#.#####.#.#.#.###########.#.#.#####.#####.#.###.#.#.#.#######.#.#.#",
	"#...###...#.#.#...#...#.#.###...#.......#...#.......#.....#.#...#...#.#.#.#...#...#...#...#.......#...#.......#...#.#.#...#.#.#.#.......#...#",
	"#.#.###.#.#.#.#########.#.#######.#####.#.#.#.#############.#.###.#.#.#.#.###.#.#########.#.#####.#############.#.#.#.#.###.#.#.#.###########",
	"#.#.....#...#...###...#...#.....#...###.#.#.#...............#...#.#.#.#.#.#...#.......#...#.....#.............#.#...#.#.#...#...#...#...#...#",
	"#.#############.###.#.#####.###.###.###.#.#.###################.#.#.#.#.#.#.#########.#.#######.#############.#.#####.#.#.#########.#.#.#.#.#",
	"#...........#...#...#.....#...#.###...#...#.#.....#...#...#.....#.#...#...#.....#.....#.........#.....#...###...###...#...#.......#.#.#.#.#.#",
	"###########.#.###.#######.###.#.#####.#####.#.###.#.#.#.#.#.#####.#############.#.###############.###.#.#.#########.#######.#####.#.#.#.#.#.#",
	"###...#...#.#...#.......#.#...#.#.....#...#...###...#...#.#.#...#.#...#...#.....#.................###...#.....#...#.....#...#...#...#.#...#.#",
	"###.#.#.#.#.###.#######.#.#.###.#.#####.#.###############.#.#.#.#.#.#.#.#.#.#################################.#.#.#####.#.###.#.#####.#####.#",
	"#...#...#.#...#...#...#.#.#...#.#.......#.............#...#.#.#...#.#...#...###...#.............###...###.....#.#.###...#...#.#.......#...#.#",
	"#.#######.###.###.#.#.#.#.###.#.#####################.#.###.#.#####.###########.#.#.###########.###.#.###.#####.#.###.#####.#.#########.#.#.#",
	"#.......#.###...#...#...#.....#.#...........#.........#.....#...#...#...#.....#.#.#.......#...#.....#.....#.....#...#.....#.#.#.........#...#",
	"#######.#.#####.###############.#.#########.#.#################.#.###.#.#.###.#.#.#######.#.#.#############.#######.#####.#.#.#.#############",
	"#.......#.#...#.#...#.........#.#.........#.#.........#...#.....#.#...#.#...#.#.#...#.....#.#.#...#...#.....#.......#...#.#...#...........###",
	"#.#######.#.#.#.#.#.#.#######.#.#########.#.#########.#.#.#.#####.#.###.###.#.#.###.#.#####.#.#.#.#.#.#.#####.#######.#.#.###############.###",
	"#.......#...#.#...#.#.#.......#...#.......#.....#.....#.#.#.....#.#.#...###.#.#...#.#.....#.#.#.#...#...#.....#.......#...#.....#.......#...#",
	"#######.#####.#####.#.#.#########.#.###########.#.#####.#.#####.#.#.#.#####.#.###.#.#####.#.#.#.#########.#####.###########.###.#.#####.###.#",
	"###...#.#...#...###...#.........#...#...###...#...#.....#.#...#.#...#.#.....#.....#...###...#...#.........#...#...#.........#...#.....#...#.#",
	"###.#.#.#.#.###.###############.#####.#.###.#.#####.#####.#.#.#.#####.#.#############.###########.#########.#.###.#.#########.#######.###.#.#",
	"#...#.#...#...#.....#.....#.....#.....#...#.#...###...#...#.#.#.#.....#.......#.......#.....#..E#...#.......#.....#.......###.........###.#.#",
	"#.###.#######.#####.#.###.#.#####.#######.#.###.#####.#.###.#.#.#.###########.#.#######.###.#.#####.#.###################.###############.#.#",
	"#...#.#.....#...#...#...#...#.....#.......#.#...#...#.#.#...#.#.#.#...###.....#.#...#...#...#.###...#.#.....#...#...#...#...#.....###...#.#.#",
	"###.#.#.###.###.#.#####.#####.#####.#######.#.###.#.#.#.#.###.#.#.#.#.###.#####.#.#.#.###.###.###.###.#.###.#.#.#.#.#.#.###.#.###.###.#.#.#.#",
	"###.#.#...#.....#.....#.....#.....#.#.......#...#.#.#.#.#...#.#.#.#.#...#.....#.#.#.#...#.#...###...#.#...#.#.#.#.#.#.#...#.#.#...#...#.#...#",
	"###.#.###.###########.#####.#####.#.#.#########.#.#.#.#.###.#.#.#.#.###.#####.#.#.#.###.#.#.#######.#.###.#.#.#.#.#.#.###.#.#.#.###.###.#####",
	"#...#...#...........#.#.....#.....#...#...#...#.#.#...#.#...#...#.#.#...#.....#.#.#...#.#.#.#######...###.#...#.#.#.#.###.#...#...#...#.#...#",
	"#.#####.###########.#.#.#####.#########.#.#.#.#.#.#####.#.#######.#.#.###.#####.#.###.#.#.#.#############.#####.#.#.#.###.#######.###.#.#.#.#",
	"#.#...#.............#...#.....###.......#...#...#...#...#.....#...#.#...#.....#.#.#...#.#.#.###.....#...#...#...#.#.#...#...#...#.###.#...#.#",
	"#.#.#.###################.#######.#################.#.#######.#.###.###.#####.#.#.#.###.#.#.###.###.#.#.###.#.###.#.###.###.#.#.#.###.#####.#",
	"#.#.#.#.....#...#...#...#.#.....#...............#...#.......#.#...#...#.###...#...#...#.#...###...#...#.....#.....#.#...#...#.#.#.....#.....#",
	"#.#.#.#.###.#.#.#.#.#.#.#.#.###.###############.#.#########.#.###.###.#.###.#########.#.#########.#################.#.###.###.#.#######.#####",
	"#.#.#.#.#...#.#.#.#.#.#...#...#.###...#.........#.........#.#.#...###.#...#...#.......#.#########.................#...###.#...#.#.....#...###",
	"#.#.#.#.#.###.#.#.#.#.#######.#.###.#.#.#################.#.#.#.#####.###.###.#.#######.#########################.#######.#.###.#.###.###.###",
	"#.#.#.#.#.#...#...#...###.....#.....#...#.....#...#.....#.#.#.#.#...#...#...#.#...#...#.###...#...#...#.....#...#.....#...#...#.#...#...#...#",
	"#.#.#.#.#.#.#############.###############.###.#.#.#.###.#.#.#.#.#.#.###.###.#.###.#.#.#.###.#.#.#.#.#.#.###.#.#.#####.#.#####.#.###.###.###.#",
	"#...#...#...#...#.........#...#.....#...#.#...#.#.#...#.#.#.#.#.#.#.###.#...#.#...#.#.#.###.#.#.#.#.#.#...#...#.......#.......#.....###.#...#",
	"#############.#.#.#########.#.#.###.#.#.#.#.###.#.###.#.#.#.#.#.#.#.###.#.###.#.###.#.#.###.#.#.#.#.#.###.#############################.#.###",
	"#...#...#.....#...#.......#.#.#...#.#.#.#.#.#...#.#...#.#.#.#.#.#.#.#...#.#...#.#...#.#.###.#.#.#.#.#...#...........................#...#...#",
	"#.#.#.#.#.#########.#####.#.#.###.#.#.#.#.#.#.###.#.###.#.#.#.#.#.#.#.###.#.###.#.###.#.###.#.#.#.#.###.###########################.#.#####.#",
	"#.#...#.#...........#.....#.#.#...#.#.#...#.#...#.#...#.#.#.#.#.#.#...#...#.###.#.#...#.###.#.#.#.#...#.#...........................#...#...#",
	"#.#####.#############.#####.#.#.###.#.#####.###.#.###.#.#.#.#.#.#.#####.###.###.#.#.###.###.#.#.#.###.#.#.#############################.#.###",
	"#.....#...#...###...#.#.....#.#.###.#...###.....#.#...#.#.#.#.#.#...###...#...#.#.#.#...#S#.#.#.#...#.#.#.........................#.....#...#",
	"#####.###.#.#.###.#.#.#.#####.#.###.###.#########.#.###.#.#.#.#.###.#####.###.#.#.#.#.###.#.#.#.###.#.#.#########################.#.#######.#",
	"#...#...#.#.#.#...#...#...#...#.#...#...#.........#...#.#.#.#.#...#.#.....#...#...#.#.###.#.#...###.#.#...#.....#.................#.....#...#",
	"#.#.###.#.#.#.#.#########.#.###.#.###.###.###########.#.#.#.#.###.#.#.#####.#######.#.###.#.#######.#.###.#.###.#.#####################.#.###",
	"#.#.....#...#...#.....#...#...#.#.#...#...#...#...#...#.#.#.#.#...#.#.#...#.....#...#...#.#.###.....#.#...#...#.#.......###.........#...#...#",
	"#.###############.###.#.#####.#.#.#.###.###.#.#.#.#.###.#.#.#.#.###.#.#.#.#####.#.#####.#.#.###.#####.#.#####.#.#######.###.#######.#.#####.#",
	"#.#...#...#...#...#...#.....#.#.#.#.###.#...#...#.#...#.#.#...#.....#...#.#...#.#...#...#...#...#.....#.#.....#...#...#.....#.......#.......#",
	"#.#.#.#.#.#.#.#.###.#######.#.#.#.#.###.#.#######.###.#.#.###############.#.#.#.###.#.#######.###.#####.#.#######.#.#.#######.###############",
	"#.#.#...#.#.#.#...#...#...#.#...#...#...#...#.....#...#...###.............#.#.#.#...#.....#...#...#.....#.......#.#.#.#.....#.###...#...#...#",
	"#.#.#####.#.#.###.###.#.#.#.#########.#####.#.#####.#########.#############.#.#.#.#######.#.###.###.###########.#.#.#.#.###.#.###.#.#.#.#.#.#",
	"#...#.....#.#.#...###...#.#.....###...#...#.#.#...#.........#.#...#.....#...#.#.#.#...#...#...#...#...#.....#...#...#.#.#...#...#.#.#.#.#.#.#",
	"#####.#####.#.#.#########.#####.###.###.#.#.#.#.#.#########.#.#.#.#.###.#.###.#.#.#.#.#.#####.###.###.#.###.#.#######.#.#.#####.#.#.#.#.#.#.#",
	"#...#.....#.#...#.........#...#...#...#.#.#.#.#.#...#...#...#.#.#...###.#...#.#.#.#.#...#...#.#...#...#.###.#.......#.#.#...#...#.#.#.#...#.#",
	"#.#.#####.#.#####.#########.#.###.###.#.#.#.#.#.###.#.#.#.###.#.#######.###.#.#.#.#.#####.#.#.#.###.###.###.#######.#.#.###.#.###.#.#.#####.#",
	"#.#.......#...###...#.....#.#...#...#.#.#.#.#.#.#...#.#...###...#...#...#...#...#.#.....#.#.#.#...#.#...#...#.......#...###.#.....#...#.....#",
	"#.###########.#####.#.###.#.###.###.#.#.#.#.#.#.#.###.###########.#.#.###.#######.#####.#.#.#.###.#.#.###.###.#############.###########.#####",
	"#.#...#.....#.....#...###.#...#...#.#.#.#...#...#.#...#...........#...#...#...###.....#.#.#.#.#...#...###...#...#.........#.#.....#...#.#...#",
	"#.#.#.#.###.#####.#######.###.###.#.#.#.#########.#.###.###############.###.#.#######.#.#.#.#.#.###########.###.#.#######.#.#.###.#.#.#.#.#.#",
	"#...#.#...#.#...#...#.....#...#...#.#.#.#.........#.#...#.....#.......#.....#.......#...#.#...#.###...#...#...#...#...#...#.#...#.#.#.#...#.#",
	"#####.###.#.#.#.###.#.#####.###.###.#.#.#.#########.#.###.###.#.#####.#############.#####.#####.###.#.#.#.###.#####.#.#.###.###.#.#.#.#####.#",
	"#.....#...#.#.#.#...#.#...#.###.....#...#.#.......#.#.#...#...#.#.....#.....#.......###...#...#.....#...#...#...#...#.#...#...#.#.#.#.#...#.#",
	"#.#####.###.#.#.#.###.#.#.#.#############.#.#####.#.#.#.###.###.#.#####.###.#.#########.###.#.#############.###.#.###.###.###.#.#.#.#.#.#.#.#",
	"#.....#...#.#.#.#...#...#.#...........#...#.#...#...#.#.#...#...#...#...###...###...#...#...#.#.............#...#.###.#...#...#.#.#.#.#.#.#.#",
	"#####.###.#.#.#.###.#####.###########.#.###.#.#.#####.#.#.###.#####.#.###########.#.#.###.###.#.#############.###.###.#.###.###.#.#.#.#.#.#.#",
	"#.....#...#...#...#.#.....#.......#...#...#.#.#.......#.#...#.#.....#.###.....#...#...#...#...#.............#...#...#.#...#...#.#...#.#.#...#",
	"#.#####.#########.#.#.#####.#####.#.#####.#.#.#########.###.#.#.#####.###.###.#.#######.###.###############.###.###.#.###.###.#.#####.#.#####",
	"#...#...#####...#...#.....#.....#...#...#...#...#.....#...#...#.#...#...#.#...#.....#...###...#...#.....#...###...#.#.#...#...#.###...#...###",
	"###.#.#######.#.#########.#####.#####.#.#######.#.###.###.#####.#.#.###.#.#.#######.#.#######.#.#.#.###.#.#######.#.#.#.###.###.###.#####.###",
	"#...#.....#...#...........#...#.......#.#.......#.#...#...#.....#.#...#...#.....#...#.....#...#.#.#...#.#...#.....#.#.#.#...#...#...#...#...#",
	"#.#######.#.###############.#.#########.#.#######.#.###.###.#####.###.#########.#.#######.#.###.#.###.#.###.#.#####.#.#.#.###.###.###.#.###.#",
	"#.........#.................#...........#.........#.....###.......###...........#.........#.....#.....#.....#.......#...#.....###.....#.....#",
	"#############################################################################################################################################",
}
