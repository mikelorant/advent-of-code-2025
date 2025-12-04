package main

import (
	"log"
)

type Warehouse []Row

type Row []*Cell

type Cell struct {
	Kind      Kind
	Processed bool
	Neighbour Neighbour
}

type Neighbour map[Direction]*Cell

type Direction int

const (
	TopLeft Direction = iota
	Top
	TopRight
	Left
	Right
	BottomLeft
	Bottom
	BottomRight
)

type Kind int

const (
	Empty Kind = iota
	Paper
)

func main() {
	do := Task("input1.txt", 1)
	log.Println("Part 1:", do)

	do = Task("input1.txt", 2)
	log.Println("Part 2:", do)
}

func Task(file string, part int) int {
	wh, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return count(wh, false)
	case 2:
		return count(wh, true)
	}

	return 0
}

func count(wh Warehouse, iter bool) int {
	var cnt int

	for {
		scan(wh)
		proc := process(wh)

		cnt += proc

		if proc == 0 || iter == false {
			return cnt
		}

		remove(wh)
	}
}

func scan(wh Warehouse) {
	for y, row := range wh {
		for x, cell := range row {
			if cell.Kind != Paper {
				continue
			}

			wh.check(x, y)
		}
	}
}

func process(wh Warehouse) int {
	var cnt int

	for _, row := range wh {
		for _, cell := range row {
			if cell.Kind != Paper {
				continue
			}

			if cell.Processed != true {
				continue
			}

			cnt++
		}
	}

	return cnt
}

func remove(wh Warehouse) {
	for _, row := range wh {
		for _, cell := range row {
			if cell.Kind != Paper {
				continue
			}

			if !cell.Processed {
				continue
			}

			cell.Kind = Empty
		}
	}
}

func (wh *Warehouse) check(x, y int) {
	cell := (*wh)[y][x]

	var cnt int

	for _, dir := range allDirections() {
		if cell.Neighbour[dir] == nil {
			continue
		}

		if cell.Neighbour[dir].Kind != Paper {
			continue
		}

		cnt++
	}

	if cnt < 4 {
		cell.Processed = true
	}
}

func allDirections() []Direction {
	all := []Direction{
		TopLeft,
		Top,
		TopRight,
		Left,
		Right,
		BottomLeft,
		Bottom,
		BottomRight,
	}

	return all
}
