package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parse(fn string) ([]int, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	var rots []int

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()

		dir := line[:1]
		rot, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("unable to parse rotation: %w", err)
		}

		if dir == "L" {
			rot = -rot
		}

		rots = append(rots, rot)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return rots, nil
}
