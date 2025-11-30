package main

import (
	"log"
)

func main() {
	do := Task("input1.txt", 1)
	log.Println("Part 1:", do)
}

func Task(file string, part int) int {
	_, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return 0
	case 2:
		return 0
	}

	return 0
}
