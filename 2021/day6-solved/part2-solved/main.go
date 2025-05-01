package day

func main(fishes []int) int {
	fishMap := parse(fishes)
	for day := 0; day < 256; day++ {
		newMap := [9]int{}
		newMap[8] = fishMap[0]
		for age := 0; age < 8; age++ {
			newMap[age] = fishMap[age+1]
		}
		newMap[6] += fishMap[0]
		fishMap = newMap
	}
	return countAll(fishMap)
}

func parse(fishes []int) [9]int {
	fishMap := [9]int{}
	for _, age := range fishes {
		fishMap[age]++
	}
	return fishMap
}

func countAll(fishMap [9]int) int {
	count := 0
	for _, num := range fishMap {
		count += num
	}
	return count
}
