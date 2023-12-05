package day5

import (
	"strconv"
	"strings"

	"goAoc2023/utils"
)

type translationRange struct {
	start      int
	end        int
	conversion int // add this to start or end to translate into next category; subtract this to translate backwards
}

func runB(input []string) int {
	seedsInts := getSeeds(input[0])
	sRanges := getSeedsTranslationRanges(seedsInts)

	mapsUnsorted := parseRanges(input[2:])
	maps := [][]translationRange{}
	for _, unsorted := range mapsUnsorted {
		maps = append(maps, sortRanges(unsorted))
	}

	translationOfSeedsToLocation := updateSeedsTranslationWithAllMaps(sRanges, maps)
	return getMinLocation(translationOfSeedsToLocation)
}

func getSeedsTranslationRanges(ints []int) []translationRange {
	ranges := make([]translationRange, len(ints)/2)
	for index := 0; index < len(ints); index += 2 {
		newRange := translationRange{
			start:      ints[index],
			end:        ints[index] + ints[index+1] - 1,
			conversion: 0,
		}
		ranges[index/2] = newRange
	}

	return sortRanges(ranges)
}

func sortRanges(ranges []translationRange) []translationRange {
	sorted := []translationRange{}
	for _, r := range ranges {
		if len(sorted) == 0 {
			sorted = append(sorted, r)
			continue
		}

		inserted := false
		for i, other := range sorted {
			if r.start < other.start {
				sorted = append(sorted[:i+1], sorted[i:]...)
				sorted[i] = r
				inserted = true
				break
			}
		}
		if !inserted {
			sorted = append(sorted, r)
		}
	}
	return sorted
}

func parseRanges(input []string) [][]translationRange {
	maps := make([][]translationRange, 7)
	index := 0
	newMap := []translationRange{}
	for j := 0; j < len(input); j++ {
		if input[j] == "" {
			maps[index] = newMap
			newMap = []translationRange{}
			index++
			continue
		}

		if _, err := strconv.Atoi(input[j][0:1]); err != nil {
			continue
		}

		dest_source_length := strings.Split(input[j], " ")
		dest := utils.Atoi(dest_source_length[0])
		source := utils.Atoi(dest_source_length[1])
		length := utils.Atoi(dest_source_length[2])

		newRange := translationRange{
			start:      source,
			end:        source + length - 1,
			conversion: dest - source,
		}
		newMap = append(newMap, newRange)
	}
	return maps
}

func updateSeedsTranslationWithAllMaps(sRanges []translationRange, maps [][]translationRange) []translationRange {
	for _, nextMap := range maps {
		sRanges = updateSeedsTranslation(sRanges, nextMap)
	}
	return sRanges
}

func updateSeedsTranslation(sRanges []translationRange, nextMap []translationRange) []translationRange {
	newRanges := []translationRange{}
	for _, sRange := range sRanges {
		start, end := translateFromSeedToLastCategory(sRange)
		hasRemainder := true
		for mIndex, mRange := range nextMap {
			if mRange.end < start || mRange.start > end {
				continue
			}

			nextMStart := end + 1 // bigger than this range cares about
			if len(nextMap) > mIndex+1 {
				nextMStart = nextMap[mIndex+1].start
			}

			sRange, newRanges = updateSeedsRange(sRange, mRange, nextMStart, newRanges)

			if sRange.start > sRange.end {
				hasRemainder = false
				break
			}
		}
		if hasRemainder {
			newRanges = append(newRanges, sRange)
		}
	}

	return newRanges
}

func translateFromSeedToLastCategory(sRange translationRange) (int, int) {
	start := sRange.start + sRange.conversion
	end := sRange.end + sRange.conversion
	return start, end
}

func translateFromLastCategoryToSeed(mRange translationRange, conversion int) (int, int) {
	start := mRange.start - conversion
	end := mRange.end - conversion
	return start, end
}

func updateSeedsRange(sRange translationRange, mRange translationRange, nextMStart int, newRanges []translationRange) (translationRange, []translationRange) {
	newRanges = addRangeBeforeM(sRange, mRange, newRanges)
	newRanges = addRangeFromOverlap(sRange, mRange, newRanges)
	sRange = getRemainingSRange(sRange, mRange)
	return sRange, newRanges
}

func addRangeBeforeM(sRange translationRange, mRange translationRange, newRanges []translationRange) []translationRange {
	start, _ := translateFromSeedToLastCategory(sRange)
	_, end := translateFromLastCategoryToSeed(mRange, -1*sRange.conversion)
	if start < mRange.start {
		rangeBeforeM := translationRange{
			start:      sRange.start,
			end:        end - 1,
			conversion: sRange.conversion,
		}
		newRanges = append(newRanges, rangeBeforeM)
	}
	return newRanges
}

func addRangeFromOverlap(sRange translationRange, mRange translationRange, newRanges []translationRange) []translationRange {
	rangeFromOverlap := translationRange{
		conversion: sRange.conversion + mRange.conversion,
	}

	if mRange.start <= sRange.start+sRange.conversion {
		rangeFromOverlap.start = sRange.start
	} else {
		rangeFromOverlap.start = mRange.start - sRange.conversion
	}

	if mRange.end >= sRange.end+sRange.conversion {
		rangeFromOverlap.end = sRange.end
	} else {
		rangeFromOverlap.end = mRange.end - sRange.conversion
	}

	return append(newRanges, rangeFromOverlap)
}

func getRemainingSRange(sRange translationRange, mRange translationRange) translationRange {
	return translationRange{
		conversion: sRange.conversion,
		start:      mRange.end - sRange.conversion + 1,
		end:        sRange.end,
	}
}

func getMinLocation(sRanges []translationRange) int {
	minLocation := sRanges[0].start + sRanges[0].conversion
	for _, sRange := range sRanges {
		location := sRange.start + sRange.conversion
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

// --------------------------------- Part A ------------------------- //

type almanacMap struct {
	from      string
	to        string
	mapRanges []almanacRange
}

type almanacRange struct {
	sourceStart      int
	destinationStart int
	length           int
}

func runA(input []string) int {
	seeds := getSeeds(input[0])
	maps := parseMaps(input[2:])
	locations := translateSeedsToLocations(seeds, maps)
	return minimum(locations)
}

func getSeeds(str string) []int {
	seedStrings := strings.Split(strings.Split(str, ": ")[1], " ")
	seeds := make([]int, len(seedStrings))
	for i, s := range seedStrings {
		seeds[i] = utils.Atoi(s)
	}
	return seeds
}

func parseMaps(input []string) []almanacMap {
	maps := make([]almanacMap, 7)
	index := 0
	newMap := almanacMap{}
	for j := 0; j < len(input); j++ {
		if input[j] == "" {
			maps[index] = newMap
			newMap = almanacMap{}
			index++
			continue
		}

		_, err := strconv.Atoi(input[j][0:1])
		if err != nil {
			from_to := strings.Split(strings.Split(input[j], " ")[0], "-to-")
			newMap.from = from_to[0]
			newMap.to = from_to[1]
			continue
		}

		dest_source_length := strings.Split(input[j], " ")
		newMapRange := almanacRange{
			destinationStart: utils.Atoi(dest_source_length[0]),
			sourceStart:      utils.Atoi(dest_source_length[1]),
			length:           utils.Atoi(dest_source_length[2]),
		}
		newMap.mapRanges = append(newMap.mapRanges, newMapRange)
	}
	return maps
}

func translateSeedsToLocations(seeds []int, maps []almanacMap) []int {
	locations := make([]int, len(seeds))
	for i, s := range seeds {
		s = translateSeedToLocation(s, maps)
		locations[i] = s
	}
	return locations
}

func translateSeedToLocation(s int, maps []almanacMap) int {
	for _, m := range maps {
		for _, r := range m.mapRanges {
			if s < r.sourceStart || s >= r.sourceStart+r.length {
				continue
			}

			s = r.destinationStart + s - r.sourceStart
			break
		}
	}
	return s
}

func minimum(ints []int) int {
	min := ints[0]
	for _, n := range ints[1:] {
		if n < min {
			min = n
		}
	}
	return min
}
