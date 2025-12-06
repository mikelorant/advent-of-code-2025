package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mode int

const (
	Ranges Mode = iota
	Available
)

func parse(fn string) (Database, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return Database{}, fmt.Errorf("unable to open file: %w", err)
	}

	var db Database
	var mode Mode

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			mode = Available
			continue
		}

		switch mode {
		case Ranges:
			rs := strings.Split(line, "-")
			low, _ := strconv.Atoi(rs[0])
			high, _ := strconv.Atoi(rs[1])
			db.Ranges = append(db.Ranges, Range{
				Low:  low,
				High: high,
			})
		case Available:
			id, _ := strconv.Atoi(line)
			db.Ingredients = append(db.Ingredients, id)
		}
	}

	if err := scanner.Err(); err != nil {
		return Database{}, fmt.Errorf("unable to scan file: %w", err)
	}

	return db, nil
}
