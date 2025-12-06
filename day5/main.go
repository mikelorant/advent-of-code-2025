package main

import (
	"log"
)

type Database struct {
	Ranges      []Range
	Ingredients []int
}

type Range struct {
	Low  int
	High int
}

func main() {
	do := Task("input1.txt", 1)
	log.Println("Part 1:", do)

	do = Task("input1.txt", 2)
	log.Println("Part 2:", do)
}

func Task(file string, part int) int {
	db, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return countFresh(db)
	case 2:
		return countFreshRanges(db)
	}

	return 0
}

func countFresh(db Database) int {
	var fresh int

	for _, id := range db.Ingredients {
		for _, r := range db.Ranges {
			if id >= r.Low && id <= r.High {
				fresh++
				break
			}
		}
	}

	return fresh
}

func countFreshRanges(db Database) int {
	var cnt int

	for {
		o1, o2 := findOverlapIndexes(db.Ranges)

		if o1 == -1 && o2 == -1 {
			break
		}

		db.Ranges = append(db.Ranges, mergeRanges(db.Ranges[o1], db.Ranges[o2]))
		db.Ranges = deleteIndexes(db.Ranges, []int{o1, o2})
	}

	for _, r := range db.Ranges {
		cnt += r.High - r.Low + 1
	}

	return cnt
}

func findOverlapIndexes(rs []Range) (a, b int) {
	for idx1, r1 := range rs {
		for idx2, r2 := range rs {
			if idx1 == idx2 {
				continue
			}

			if isOverlap(r1, r2) {
				return idx1, idx2
			}
		}
	}

	return -1, -1
}

func mergeRanges(a, b Range) Range {
	return Range{
		Low:  min(a.Low, b.Low),
		High: max(a.High, b.High),
	}
}

func isOverlap(a, b Range) bool {
	switch {
	// a.High is inside b
	case a.High >= b.Low && a.High <= b.High:
		return true
	// a.Low is inside b
	case a.Low >= b.Low && a.Low <= b.High:
		return true
	default:
		return false
	}
}

func deleteIndexes(rs []Range, dis []int) []Range {
	var ranges []Range

	for idx, r := range rs {
		var delete bool

		for _, di := range dis {
			if idx == di {
				delete = true
				break
			}
		}

		if delete {
			continue
		}

		ranges = append(ranges, r)
	}

	return ranges
}
