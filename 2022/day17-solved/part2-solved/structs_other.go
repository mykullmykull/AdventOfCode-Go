package day

// this struct allows us to save these counters as values in a single cache
type counters struct {
	droppedCount int
	height       int
}

// this struct allows us to save this state as a key in a single cache
type state struct {
	towerCols columns
	rockIndex int
	windIndex int
}
