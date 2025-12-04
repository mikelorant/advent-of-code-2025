package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fn string) (Warehouse, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	var wh Warehouse
	var idx int

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		row := lineToRow(scanner.Text(), idx)
		wh = append(wh, row)
		idx++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	wh.setNeighbours()

	return wh, nil
}

func lineToRow(line string, y int) Row {
	var row Row

	for _, char := range line {
		var kind Kind

		if char == '@' {
			kind = Paper
		}

		row = append(row, &Cell{
			Kind: kind,
		})
	}

	return row
}

func (w *Warehouse) setNeighbours() {
	rows := len(*w)
	cols := len((*w)[0])

	for y := range rows {
		for x := range cols {
			cell := (*w)[y][x]
			cell.Neighbour = make(Neighbour)

			for _, dir := range allDirections() {
				nx, ny := toDirection(x, y, dir)

				if nx < 0 || ny < 0 || nx >= cols || ny >= rows {
					continue
				}

				cell.Neighbour[dir] = (*w)[ny][nx]
			}
		}
	}
}

func toDirection(x, y int, dir Direction) (int, int) {
	switch dir {
	case TopLeft:
		x--
		y--
	case Top:
		y--
	case TopRight:
		x++
		y--
	case Left:
		x--
	case Right:
		x++
	case BottomLeft:
		x--
		y++
	case Bottom:
		y++
	case BottomRight:
		x++
		y++
	}

	return x, y
}
