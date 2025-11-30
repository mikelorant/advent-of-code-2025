package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fn string) (any, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		// scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return nil, nil
}
