package day13

var testInput = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
	"",
}

var expected = 405
var expected2 = 400

var realInput = []string{
	".##.##.##....",
	"#.######.##..",
	"##......###..",
	"...####......",
	"...####...###",
	"..#....#..#..",
	".#.####.#.###",
	"..#....#..###",
	".##.##.##....",
	"..######..###",
	"#.##..##.####",
	"#.#.##.#.....",
	"##..##..##.##",
	".#..##..#....",
	"#.##..##.##..",
	"",
	".##.#.##.#.",
	"#..##..####",
	"#..##..##.#",
	".##.#.##.#.",
	"....#.##...",
	".#.###..##.",
	"...#.##..##",
	".#....#.###",
	".#....#.###",
	"",
	"##.#.####.#.####.",
	"#..##.##.##..#..#",
	"..###....###..###",
	".###.####.###..##",
	"..#........#..##.",
	"#..###..###..#.##",
	"#..###..###..#.##",
	"..#........#...#.",
	".###.####.###..##",
	"",
	".....####..",
	"##...####.#",
	"##..#....#.",
	".....####..",
	"####.####.#",
	"##..#.##.#.",
	"##.###..###",
	"...#......#",
	"..#.######.",
	"##.##.##.##",
	"..#..#..#..",
	"##.#.####.#",
	"####..##..#",
	"",
	"##..##.#.##.#.#",
	"..##..#....#...",
	"#.#...#.#.###.#",
	"####.##..#.#..#",
	"#..#....#..####",
	"##.#.#..#....##",
	"#.#.####.#..###",
	"#.#.####.#..###",
	"##.#.#..#.#..##",
	"##.#.#..#.#..##",
	"#.#.####.#..###",
	"#.#.####.#..###",
	"##.#.#..#....##",
	"",
	"#..##...#.####.##",
	"#..##...#.####.##",
	"....###.#..#.###.",
	"#####..####..#.##",
	"#..#..##...###..#",
	".##.##..#..###...",
	".##.#.#.#...#..#.",
	"##.#...#.#.....#.",
	"........######.#.",
	"....#..#....#..##",
	"#######..#.##..#.",
	"#..#...#######..#",
	"....###..###.##.#",
	"",
	"#.#..#.##..##",
	"##.##.##....#",
	"###..##......",
	"###..###.##.#",
	".######.#..#.",
	"##.##.##....#",
	"#......##..##",
	"",
	"#.#..#.###.....",
	"..####..#..#..#",
	"#......##..####",
	"#..##..#.#.#..#",
	"...##....#.#...",
	".######...#....",
	"#..##..##.#....",
	"",
	"..##...",
	"..#.#..",
	"..#.#..",
	"..##...",
	".#..###",
	"##....#",
	"....#..",
	"#..#...",
	".##....",
	"...#...",
	".####..",
	"",
	"##..##..###",
	"#.##.######",
	"#.##.######",
	"##..##..###",
	"#######..##",
	".#####..#.#",
	".####....#.",
	"",
	"#....#...##",
	".####......",
	"##.#..#...#",
	"...#.#..###",
	"...#.#..###",
	"##.#..#...#",
	".####......",
	"#....#...##",
	".#.##......",
	"..##.#.#..#",
	"..##.#.#..#",
	".#.##......",
	"#....#...##",
	".####...#..",
	"##.#..#...#",
	"",
	"##.#...",
	"##.#...",
	".#.#.##",
	"#.###.#",
	"......#",
	"..#.##.",
	"#.#..##",
	"#.#...#",
	"#.#...#",
	"#.#..##",
	"..#.##.",
	"......#",
	"#.#.#.#",
	".#.#.##",
	"##.#...",
	"",
	".#...###.##..",
	"#..#.##..#...",
	"#.#.#..###.#.",
	".#.#####....#",
	"..##..#..#.##",
	"..##..#..#.##",
	".#.#####....#",
	"#.#.#..###.#.",
	"#..#.##..#...",
	".#...###.##..",
	".#...###..#..",
	"#..#.##..#...",
	"#.#.#..###.#.",
	".#.#####....#",
	"..##..#..#.##",
	"",
	"#..#.#.##...#.#..",
	"..##....####.##.#",
	".###.###...#.####",
	"###....#.......#.",
	".#.#######..#.###",
	"#..#..##..##.#..#",
	"#..#..##..##.#..#",
	".#.#######..#.###",
	"###....#.......#.",
	".###.###...#.####",
	"..##....####.##.#",
	"#..#.#.###..#.#..",
	"#..#.#.###..#.#..",
	"..##....####.##.#",
	".###.###...#.####",
	"",
	"######....##..##.",
	"..#..#####.#.###.",
	".#####...########",
	"######...########",
	"..#..#####.#.###.",
	"######....##..##.",
	".##...#.#....#...",
	"##..##.......###.",
	".###..###.###...#",
	"....#.##.###.##..",
	"..##.#...####...#",
	"..##.#...#.###.#.",
	"#.#...#.#....##.#",
	"##.#.##..#......#",
	"##.#.##.##.##..#.",
	"...#...#.#.###.##",
	"...#...#.#.###.##",
	"",
	"..#.##..#",
	"##.#..###",
	"##..#...#",
	"###.##.##",
	"###.##.##",
	"##..#...#",
	"##.#..###",
	"..#.##..#",
	".##..##.#",
	"###.###..",
	"#..##....",
	"#.#...#..",
	"..####..#",
	"..####..#",
	"#.#...#..",
	"#..##....",
	"###.###.#",
	"",
	"..#.##.#.##.#..",
	"##..##.#.####.#",
	".....#...#...#.",
	".####....######",
	".####....######",
	".....#...#...#.",
	"##..##.#.####.#",
	"..#.##.#.##.#..",
	"#..#####.#.#.#.",
	"####..#.###..#.",
	".....##...##..#",
	"..#.#.###..#...",
	"#...#.##..#...#",
	".##.##..##.#...",
	".##.##..##.#...",
	"....#.##..#...#",
	"..#.#.###..#...",
	"",
	"......#..#.",
	".###.######",
	"###...#..#.",
	".#.###.##.#",
	"#..##......",
	"##..#.#..#.",
	"##..#.#..#.",
	"#..##......",
	".#.###.##.#",
	"###...#..#.",
	".##########",
	"......#..#.",
	"...###....#",
	"..#...####.",
	".#.#.#.##.#",
	"",
	"#.##..##.##",
	"###.##.####",
	"#..####..#.",
	"#..####..#.",
	"###.##.####",
	"#.##..##.##",
	".........#.",
	"#...##...##",
	"###.##.###.",
	"....##....#",
	"#..####..#.",
	"",
	".#.##.....#.##.#.",
	".##...####.####.#",
	"......###........",
	".#...##..#......#",
	".#...#.#.##....##",
	"#.#.#..###......#",
	".#.##.....######.",
	".#.##.....######.",
	"#.#.#...##......#",
	"",
	"###..##..",
	"######..#",
	"...##....",
	"##.#.#.##",
	"###..#...",
	"###..#...",
	"##.#.#.##",
	"...##....",
	"###.##..#",
	"",
	".....#....#.#",
	".##.#.##.####",
	"#####..###..#",
	".##.###.#.##.",
	"....#.#.....#",
	"#..#..#....#.",
	"......##.#...",
	"#######..#.##",
	"###.###.####.",
	"####.##...#.#",
	"...........##",
	"....#.#.##.#.",
	"#####.#.##.#.",
	"#####..#.#.#.",
	"#..##..####..",
	"#..##..####..",
	"#####..#.#.#.",
	"",
	".#.#....#.###.#",
	"####...#..##.##",
	"#.#...#..#...##",
	"#.####..#.#..#.",
	"..#..##.#.#...#",
	"##..#..###.####",
	"##..#..###.####",
	"..#..##.#.#...#",
	"#.####..#.#..#.",
	"#.#......#...##",
	"####...#..##.##",
	".#.#....#.###.#",
	".#.#....#.###.#",
	"",
	"##..#..#..#",
	"##.########",
	"###..##..##",
	"..###..###.",
	"##.######.#",
	"##........#",
	"..##.##.##.",
	"....####...",
	"##..####..#",
	"...#....#..",
	"..##.##.##.",
	"",
	"##...##.####.",
	"#...#..##.##.",
	".##.##.#.#..#",
	"..###..#####.",
	"..###..#####.",
	".##.##.#.#..#",
	"#...#..##.##.",
	"##...##.####.",
	"##...##.####.",
	"#..##..##.##.",
	".##.##.#.#..#",
	"",
	".........#.",
	"##.#.#..###",
	"##.#.#..###",
	".........#.",
	"#.#.##.#.#.",
	"#.##.#.#..#",
	".##.##.#.#.",
	"###.###....",
	"#.#.####.##",
	"#.#.####.#.",
	"###.###....",
	".##.##.#.#.",
	"#.##.#.#..#",
	"#.#.##.#.#.",
	".........#.",
	"",
	"###....",
	"####.##",
	"###..##",
	"...##..",
	"...####",
	"....#.#",
	"##.#.##",
	"..#..##",
	"....###",
	"",
	".#..#.###.#",
	"..##....###",
	"#.##.#..#..",
	"##..##.#...",
	"##..##.##.#",
	"##..####.#.",
	"##..####.#.",
	"##..##.#..#",
	"##..##.#...",
	"#.##.#..#..",
	"..##....###",
	"",
	".##...#",
	".##....",
	"####...",
	".##..#.",
	"....#.#",
	"#####..",
	"#..##..",
	"....#.#",
	".##..#.",
	"",
	".....###..###",
	"#..#.########",
	"####.#......#",
	"#.###########",
	"#.###########",
	"#.##.#......#",
	"#..#.########",
	".....###..###",
	".##.##.####.#",
	"..#....#..#..",
	"..##.#.####.#",
	"#.###........",
	"....#.#....#.",
	".#######..###",
	".###.#.####.#",
	"#.#.#...##...",
	"##...#.####.#",
	"",
	"#....##.#.##.",
	"..#.#..#.#..#",
	"..##....#....",
	"###..####.##.",
	"#..##.....##.",
	".#...#.######",
	".#...#.######",
	"#..##.....##.",
	"###..###..##.",
	"..##....#....",
	"..#.#..#.#..#",
	"",
	".##....",
	"##.####",
	"#.#####",
	"#..#..#",
	".######",
	"...####",
	"...####",
	"###....",
	".#.####",
	"##.....",
	"#.#####",
	"...####",
	"#..####",
	"...####",
	"#......",
	"",
	"....###.#",
	"#..###...",
	"####.##.#",
	".....#..#",
	"####...##",
	"#..#..#..",
	"....##...",
	".....#..#",
	".##....#.",
	".##..#...",
	"#####....",
	".##...#.#",
	".##......",
	".##..#...",
	".##...#.#",
	"#####....",
	".##..#...",
	"",
	"###.##.##########",
	"...####...#..#...",
	".#.#..#.#......#.",
	"##..##..########.",
	"#...##...##...#..",
	"###.##.##########",
	"##......###..###.",
	".##....##......##",
	"..######..####..#",
	"##.####.###..###.",
	".###..###..##..##",
	"....##...........",
	"##########....###",
	".#.#..#.#.#..#.#.",
	".#..##..#.#..#.#.",
	"..##..##...##...#",
	"..#....#........#",
	"",
	"#..#....#..####",
	"....#..#.....#.",
	"##.##..##.##.#.",
	"...######......",
	"#..#....#..#.##",
	"#..#....#..#.##",
	"...######......",
	"##.##..##.##.#.",
	"....#..#.....#.",
	"#..#....#..##.#",
	"..#..##..#....#",
	"#...####...###.",
	"#.#......#.#..#",
	"....#..#....#..",
	"...#.##.#....#.",
	"",
	"##..##..##.",
	".#..#....#.",
	"######..###",
	".#..#.##.#.",
	"#.##.####.#",
	".####....##",
	"#.##.#..#.#",
	".####....##",
	".#..#.##.#.",
	"######..###",
	"######..###",
	"######..###",
	"......##...",
	".##.#....#.",
	"#.##.####.#",
	"",
	"#...#.#.#........",
	"#.#.####.########",
	"####.#.#..##..##.",
	"#.#######.######.",
	"#.###.###.##..##.",
	"#.###.#.#########",
	"##.#.############",
	".#.....##........",
	"..#.###..########",
	"",
	"#..##.#...#",
	"#...#..###.",
	"##...#.####",
	".##.##.#.##",
	".##.##.#.##",
	"##...#.####",
	"#......###.",
	"#..##.#...#",
	".##..#.#.##",
	".##..#.#.##",
	"#..##.#...#",
	"",
	".###.##.##.",
	"#.##..#....",
	"#..#.###..#",
	"#..#.###..#",
	"#.###.#....",
	".###.##.##.",
	".#.#####..#",
	"...#.#.....",
	"##.....####",
	"",
	"##..#..#..#",
	"##..#..#..#",
	"...#..#.#..",
	".#.##..##.#",
	"###......##",
	"#.###..###.",
	"###.####.##",
	"###.#..#.##",
	"#..#.##.#..",
	"##........#",
	"..#.####.#.",
	"..#..##..#.",
	"....#..#...",
	"",
	"###...####....#",
	"..#..#.#.#.##.#",
	"..#..#.#.#.##.#",
	"###.#.####....#",
	"#.#..##........",
	"..##.##########",
	"##..#..#.......",
	"#...#.#.#######",
	"#..###.###.##.#",
	"##..###########",
	"#..#..#.##.##.#",
	"#...###..##..##",
	".##....#..####.",
	"#.###...#.#..#.",
	"###.....#..##..",
	"",
	"....####.",
	".#.......",
	"....####.",
	"#........",
	"#........",
	"....####.",
	".#.......",
	".#..####.",
	"#.###..##",
	"",
	".#.##.##.....##..",
	".######.#..######",
	"#....#..#.##....#",
	"..####.###.##..##",
	"..#..##..#...###.",
	".##.##.#.....##..",
	".##.##.#.....##..",
	"",
	"..........#####",
	".#..##..#..#.##",
	"#..#..##.##.#..",
	"####..####.####",
	"#.##..##.##.###",
	"...#..#...#..##",
	".#.####.#.#####",
	"...#..#...#.#..",
	".#......#......",
	"...#..#...#..##",
	"#..####..##..##",
	"",
	"...##.#..",
	"##.#...##",
	"...#..#.#",
	"...#..#.#",
	"##.#..###",
	"...##.#..",
	"..##...##",
	"#####.###",
	"####.##..",
	"..#.#####",
	"####..#.#",
	"....##...",
	"##.###.#.",
	"##..###.#",
	"####.#.##",
	"",
	"..#.#.#",
	".##.##.",
	"#####..",
	"####.#.",
	"##.#..#",
	"##.#..#",
	"####.#.",
	"",
	"..##..#.##.#.",
	"#.###.######.",
	"##..##.#..#.#",
	".#.##...##...",
	"..####.####.#",
	"...#.###..###",
	"...#.###..###",
	"#.####.####.#",
	".#.##...##...",
	"##..##.#..#.#",
	"#.###.######.",
	"",
	"..#####..#.....",
	".##.#.#......##",
	".#####.###..###",
	"##......###.#..",
	"###..##.#..#.##",
	"..##.###....#..",
	".#.##...####...",
	".#.##...###....",
	"..##.###....#..",
	"###..##.#..#.##",
	"##......###.#..",
	"",
	".##......##.#....",
	"#####..######....",
	"##.#....#.####..#",
	"..##.##.##..#.##.",
	"..#......#..##..#",
	"#.#..##..#.######",
	"...#....#.....##.",
	"...######...#....",
	"...##..##....#..#",
	"#.###..###.#.....",
	".#...##...#...##.",
	"#.#.#..#.#.#.....",
	".#...##...#.##...",
	"##...##...###....",
	"..#.#..#.#....##.",
	"..##.##.##..#####",
	".##.#..#.##.#.##.",
	"",
	"#....#.#.#..#.#..",
	"..####..#.#######",
	"..####..#.#######",
	"#....#.#.#..#.#..",
	".#.#####..#####..",
	"#..###.#..#.##..#",
	"#..#.#...###.###.",
	"#.#.##.##....##.#",
	"#.#.#....##...#..",
	"..##.##.......#..",
	"..##.##.......#..",
	"#.#.#....##...#..",
	"###.##.##....##.#",
	"#..#.#...###.###.",
	"#..###.#..#.##..#",
	"",
	".##......##..",
	".##..##..##..",
	"#..#.##.#..##",
	"#..######..#.",
	"#..#....#..##",
	".##.#..#.##..",
	".##.####.##.#",
	"##.#....#.###",
	"####....#####",
	"",
	"#..#####.",
	".....#...",
	"#..#..#..",
	"#..#..#..",
	".....#...",
	"#..#####.",
	".####....",
	"######.##",
	"#..#.#.##",
	"....##.#.",
	"....##...",
	"####.#..#",
	"#..#.....",
	"######.#.",
	".##.##..#",
	"",
	"......#",
	"#....##",
	"....#..",
	"##..##.",
	"......#",
	"##..###",
	"......#",
	"..##...",
	"..##..#",
	"#....##",
	"#....##",
	"",
	".#..####..#..##",
	"..#.#..#.#..###",
	"##..#..#..##...",
	"...#.##.#...###",
	"##..####..#####",
	".#...##...#.###",
	".##..##..##.###",
	"..#......#..###",
	"#.#......#.####",
	"############...",
	"#.#.##.#.#.#.##",
	"#..........####",
	"..##.##.##...##",
	".##.####.##.#..",
	"#..######..#...",
	"",
	"##.###..#.###.###",
	".###.##......#.##",
	"..#..##..##.##.##",
	"..#..##..##.##.##",
	".###.##......#.##",
	"##.###..#.###.###",
	"..#.#.#..#..###..",
	".#.#.#.#.#.#####.",
	".#.#.#.#.#.#####.",
	"..#.#.#..#..###..",
	"######..#.###.###",
	".###.##......#.##",
	"..#..##..##.##.##",
	"",
	"##..##..####..#",
	".#.#.....##....",
	"..##.#.##..##.#",
	".#..##.##..##.#",
	".###..#.#..#.#.",
	"##...###....###",
	"#...####....###",
	"#..#..#......#.",
	"...#..#......#.",
	"",
	"#.#..#.#.##",
	"#########..",
	"..####..#..",
	"#......#.##",
	"#......#...",
	"#.####.####",
	"..#..#..###",
	"#########..",
	"..#.....###",
	".#.##.#.###",
	"#.#..#.##..",
	"",
	"####..##.",
	"####..##.",
	"..###..##",
	"#...#...#",
	"...#.##..",
	"#...#.###",
	"#####...#",
	"..###.#..",
	"..###.#..",
	"#####...#",
	"#...#.###",
	"...#.##..",
	"#...#...#",
	"..###..##",
	"####...#.",
	"",
	".######....",
	"..#........",
	"###.#.#####",
	"###..#.#..#",
	"##..#.#....",
	".#.###.#..#",
	"##.####....",
	"#..##..#..#",
	".#..####..#",
	"....####..#",
	"#..##..#..#",
	"##.####....",
	".#.###.#..#",
	"##..#.#....",
	"###..#.#..#",
	"###.#.#####",
	"..#........",
	"",
	"#.####.###..###.#",
	"..####..#.......#",
	"#.#..#.###..###.#",
	"..####...#..#...#",
	"########......###",
	"##.##.###....###.",
	".##..##..#..#..##",
	"..####..#....#..#",
	"#......##.##.##..",
	"###..#####..#####",
	"########.#..#.###",
	"",
	"#.#.#.#",
	"..####.",
	"##.####",
	"#..##.#",
	"#.#..#.",
	".#.....",
	"####.##",
	".##.#.#",
	".#..#.#",
	"..#####",
	"..#####",
	".#..#.#",
	".##.#.#",
	"",
	"..#....",
	"...####",
	"..##..#",
	"##.#..#",
	"...#..#",
	"##.....",
	"....##.",
	".##....",
	"##.....",
	"####..#",
	"###....",
	"",
	".##.##.##..#..#..",
	".........#.#..#.#",
	"#...##...##....##",
	"#.#....#.#.#..#.#",
	"#.######.#.####.#",
	"...#..#...######.",
	"##......#########",
	"",
	"..#.###..",
	"....####.",
	"..#.####.",
	"###...##.",
	".....#.##",
	"#####.##.",
	"###......",
	"#.#.#.##.",
	"##.#.##..",
	"##......#",
	"##.#.##..",
	"..#.#...#",
	"...##...#",
	"...##...#",
	"..#.#...#",
	"",
	".####.#.####.##",
	".##...##.#..#.#",
	"#..##.##.###..#",
	".##........####",
	"#..##..#.#..##.",
	"....##...###...",
	".##...###...###",
	".##...###...###",
	"....##...###...",
	"",
	"..#.#..##.#..##",
	"#.#.##.###.##..",
	"#.#####.....#.#",
	"..#....########",
	".#.#.#.##.#..##",
	".#.#.#.##.#..##",
	"..#....########",
	"",
	"##.#####..#######",
	"###..#.#...##..##",
	"#....#..#.#......",
	"#.#..#...####..##",
	"....#.###########",
	".....####..######",
	".#..###.....#..#.",
	"##.###.#.#.##..##",
	"###.#.##.#.##..##",
	"",
	"#..#..#...##.",
	"#..###...#..#",
	"###.##..#####",
	".#..#.#......",
	".####.###.##.",
	"#.##..#.#....",
	"####..#.#....",
	"",
	"#.#..####..#.##",
	"..#...##...#...",
	"#.##########.##",
	"..####..####..#",
	"#####.##.######",
	".#.##.##.##.#.#",
	"##..######..##.",
	".#...#####..#.#",
	".###..##..###.#",
	"#.#.#.##.#.#.##",
	".#.#..##..#.#..",
	"##..######..##.",
	".#.##.##.##.#..",
	".#..#....#..#.#",
	".#..#....#..#.#",
	"",
	".#####...",
	"##....#.#",
	"#..#.#.#.",
	"..#.##...",
	"#.#.##...",
	"#..#.#.#.",
	"##....#.#",
	".#####...",
	".#.###.#.",
	"...###...",
	"...###...",
	"",
	"##.######.###",
	"####.##.####.",
	".##.####.##..",
	"###..##...###",
	"#####..#####.",
	"###.####.###.",
	"#...####...#.",
	"##.#.##.#.##.",
	".#.#.##.#.#..",
	"##.##..##.###",
	"##........###",
	"##........###",
	"##.##..##.###",
	".#.#.##.#.#..",
	"##.#.##.#.##.",
	"#...####...#.",
	"###.####.###.",
	"",
	"#.##.##..####.#",
	"......##....#..",
	"..##..#...##..#",
	"......#####....",
	"..........#.##.",
	"..##..####.####",
	".#....##......#",
	"##..##.##..####",
	".......#####...",
	"##..####..#.#.#",
	".......#.##.#..",
	"..##....##...#.",
	".#..#......####",
	"#######.#.##...",
	"#######.#.##...",
	"",
	"#.##.##.#",
	"##..#..#.",
	".####..##",
	"##..####.",
	".#..#..#.",
	"#....##..",
	"#.##.##.#",
	"",
	"##.##.#..",
	"##..##.##",
	"###.#....",
	"##.......",
	"##.#..#..",
	"..#.###..",
	"..#######",
	"......#..",
	"###.#....",
	"....#.###",
	"..###..##",
	".......##",
	"...######",
	"..#....#.",
	"..#..#.##",
	"",
	"##...........",
	"......#..#...",
	"##...........",
	"#..#.#.##.#.#",
	"#..#.#.##.#.#",
	"##...........",
	"......#..#...",
	"##...........",
	"#.#.##.##.##.",
	".....######..",
	"..###......##",
	"#.##.#....#.#",
	".##.##.##.##.",
	".....######..",
	"####.#.##.#.#",
	"#####.####.#.",
	"##.###....###",
	"",
	"#.#..#..#.#",
	"..##.##.###",
	"..##.##.###",
	"#.#..#..#.#",
	"##...#.##.#",
	"###...#.##.",
	"##.###.....",
	"#..#.##.#.#",
	"#....##.#.#",
	"##.###.....",
	"###...#.##.",
	"",
	"####.#.##...#",
	".##...##.####",
	"#####...#....",
	".....#.##.###",
	"########.#...",
	"####.#.##.###",
	"....###.####.",
	".....#.##..#.",
	"######..##.##",
	"",
	"..#.##.#.##",
	".#....#....",
	"####...##..",
	"####...##..",
	".#.........",
	"..#.##.#.##",
	".......##..",
	"###.#.##.##",
	"...#.#.....",
	"",
	"..#..#.####.#",
	"#####.##..###",
	"##.#.##.##.##",
	"..#..###..###",
	"###.###.##.##",
	"..#...##..##.",
	"###.#........",
	"##.#.........",
	"####..#....#.",
	"...#...#..#..",
	"##...##....##",
	"...#..#.##.#.",
	"....#.######.",
	"###.#..#..#..",
	"..#...#....#.",
	"",
	"####.#...",
	"....#....",
	".##..#.#.",
	".##..#.#.",
	"....#...#",
	"####.#...",
	"########.",
	"",
	"..####.#.",
	".#..##...",
	".#..##...",
	"..####.#.",
	".###...#.",
	"##.#..#.#",
	"##.#..#..",
	"##.#..#..",
	"##.#..#.#",
	".###...#.",
	".#####.#.",
	"",
	"..######.",
	"##.####.#",
	"...####..",
	"##..##..#",
	"####..###",
	"###.##.##",
	"..#.##.#.",
	"##......#",
	"..#.##.#.",
	"...####..",
	"####..###",
	"...#..#..",
	"###.##.##",
	"###.##.##",
	"##.####.#",
	"...###...",
	"###....##",
	"",
	"#..#..#..#.###...",
	".##....##..####.#",
	"...#..#...##..#..",
	"####..####..#...#",
	".########..#..#.#",
	"###.##.####...#.#",
	"...####......#..#",
	"##..##..###..#...",
	".#.#..#.#..#.#...",
	"#........###.##.#",
	"#..#..#..####.###",
	".##.##.##.#..#.#.",
	"###....###.##..##",
	"###....###.##..##",
	".##.##.##....#.#.",
	"#..#..#..####.###",
	"#........###.##.#",
	"",
	"####.#.##",
	"#..###.##",
	"#..#.####",
	".##.##..#",
	".##...##.",
	"####..#..",
	"....#.##.",
	"#..#..###",
	"####.....",
	".....#...",
	".##..##.#",
	"....#.#..",
	"....#.#..",
	".##..##.#",
	".....#.#.",
	"####.....",
	"#..#..###",
	"",
	"###.###",
	"##...#.",
	"....#.#",
	"##.#..#",
	"##.#..#",
	"....###",
	"##...#.",
	"###.###",
	"....##.",
	"..#.###",
	"###..##",
	"...###.",
	"###.#..",
	"",
	"###....##",
	"...#..#..",
	"###.##.##",
	"##.####.#",
	"##.#..#.#",
	"...#..#..",
	"###....##",
	"###....##",
	"...#####.",
	"####..###",
	"##.####.#",
	"",
	".#..##.#.##..###.",
	".#...##.#.#.#...#",
	".#...#.#....#.#.#",
	"..#..#####...##.#",
	"..#..#####...##.#",
	".#...#.#....#.#.#",
	".#...##.#.#.#...#",
	".#..##.#.##..###.",
	"##..#.##.#.###..#",
	"#...#..###...###.",
	"..#..#.#.####.###",
	"..##.#.#.##.##...",
	"..##.#.#.##.##...",
	"..#..###.####.###",
	"#...#..###...###.",
	"##..#.##.#.###..#",
	".#..##.#.##..###.",
	"",
	"##.############.#",
	"###.#.######.#.##",
	"..##.#..##..####.",
	"##.#..#.##.#..#.#",
	"..#####....#####.",
	"##...#..##..#...#",
	"..#.#.#....#.#.#.",
	"...#..#.##.#..#..",
	"..#.##..##..##.#.",
	"........##.......",
	"###..##.##.##..##",
	"##...#.####.#...#",
	"..#.##.####.##.#.",
	".....#......#....",
	"#####.#....#.####",
	"..#..########..#.",
	"#####.######.####",
	"",
	".###.#.#....#",
	"..##.#.#....#",
	"#.#.#.##....#",
	"###.#.##....#",
	".##.###..##..",
	"#.#..#.######",
	"#..###.......",
	"",
	"....##..##...",
	"##..#....#..#",
	"###.######.##",
	"##..##..##..#",
	"######..#####",
	".....####....",
	"#####....####",
	"#...#....#...",
	".....#..#....",
	"...#.####.#..",
	"##.#..##..#.#",
	"",
	"..###..####",
	"..#..##.###",
	".###..#....",
	".###..#....",
	"..#..##.###",
	"..####.####",
	"#..##.#....",
	"",
	"##.##...#....",
	"###.#.#.##..#",
	"........##.#.",
	"######....##.",
	"#..#.....##..",
	"###.##.#.#..#",
	"##...#..#..#.",
	"##.##..#.#.##",
	"##.##..######",
	"....#..#####.",
	"...#.....###.",
	"...#.....###.",
	"....#..#####.",
	"##.##..######",
	"##.##..#.#.##",
	"",
	"...#...#####.",
	"..##.##..##..",
	"..##.##..##..",
	"...#...#####.",
	"##...#.###..#",
	"###....#.#...",
	"#.###.#..##..",
	"#.##.#..#....",
	"#.##.#..#.#..",
	"",
	"#....###..#.#",
	"#....###..#.#",
	".####.##.....",
	".####..##.#..",
	".#..#..#.###.",
	"..##....#...#",
	".####.#...#..",
	"#....#.##...#",
	".........#.##",
	"..##..#..#..#",
	"........#.#..",
	"#.##.#....##.",
	"#....#.##.###",
	".####..#.##.#",
	".#..#....##..",
	"#.##.##.....#",
	"#.#..#..###..",
	"",
	"...###.",
	"..#####",
	"...#.##",
	"####.##",
	"##.###.",
	"#.#...#",
	"....##.",
	"##...##",
	"##...##",
	"",
	"..######.",
	"##..##..#",
	"##..##..#",
	"..#....#.",
	"..#.##.#.",
	"###.##.##",
	"..#....#.",
	"....##...",
	".##....##",
	"##.#..#.#",
	"##.#..#.#",
	"###.##.##",
	"...#..#..",
	"##..##..#",
	"..#.##.#.",
	"...####..",
	"..#....#.",
	"",
	"####.##..####",
	"...#.#.#..#.#",
	"##.###.....#.",
	"##.###.....#.",
	".....#.#..#.#",
	"####.##..####",
	"###.#...##.#.",
	"",
	"#.#............",
	"###.#..........",
	"...######..####",
	"###.#..#.##.#..",
	".#....##.##.##.",
	"###...##.##.##.",
	"..#.#....##....",
	"",
	"########.##.#",
	"#.####.#####.",
	"##########.##",
	"######.#.###.",
	"..#..#..##...",
	"#.#..#.#..###",
	"..#..#...###.",
	"...##...###..",
	"#..##..##...#",
	"##.##.###.##.",
	"##.##.###.##.",
	"",
	"###.#####.#",
	"###.#####.#",
	"#..##.####.",
	"####.#...#.",
	"#.....#..##",
	".##.#..##.#",
	"##.##.#...#",
	"##.##.#...#",
	".#..#..##.#",
	"#.....#..##",
	"####.#...#.",
	"",
}
