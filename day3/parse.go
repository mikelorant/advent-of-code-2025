package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fn string) ([]string, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	var banks []string

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		banks = append(banks, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan file: %w", err)
	}

	return banks, nil
}
