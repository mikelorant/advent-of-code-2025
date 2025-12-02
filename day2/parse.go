package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(fn string) ([]ID, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	var ids []ID

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		rs := strings.Split(scanner.Text(), ",")

		for _, r := range rs {
			if r == "" {
				continue
			}

			ids = append(ids, toID(r))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return ids, nil
}

func toID(str string) ID {
	ids := strings.Split(str, "-")

	first, _ := strconv.Atoi(ids[0])
	last, _ := strconv.Atoi(ids[1])

	return ID{
		First: first,
		Last:  last,
	}
}
