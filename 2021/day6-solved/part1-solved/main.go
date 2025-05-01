package day

func main(fishes []int) int {
	for days := 1; days <= 80; days++ {
		println("Day", days, "Fishes", len(fishes))
		updated := []int{}
		for len(fishes) > 0 {
			f := fishes[0]
			fishes = fishes[1:]
			if f == 0 {
				updated = append(updated, 8)
				updated = append(updated, 6)
			} else {
				updated = append(updated, f-1)
			}
		}
		fishes = updated
	}
	return len(fishes)
}
