package day

var allRotations = [][3]int{
	{0, 1, 2},
	{0, 2, 1},
	{1, 0, 2},
	{1, 2, 0},
	{2, 0, 1},
	{2, 1, 0},
}

var allFlips = [][3]int{
	{1, 1, 1},    // no flip
	{-1, 1, 1},   // flip x
	{1, -1, 1},   // flip y
	{1, 1, -1},   // flip z
	{-1, -1, 1},  // flip x and y
	{-1, 1, -1},  // flip x and z
	{1, -1, -1},  // flip y and z
	{-1, -1, -1}, // flip all
}
