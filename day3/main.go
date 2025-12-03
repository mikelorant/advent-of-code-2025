package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	do := Task("input1.txt", 1)
	log.Println("Part 1:", do)

	do = Task("input1.txt", 2)
	log.Println("Part 2:", do)
}

func Task(file string, part int) int {
	banks, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return sumBanks(banks, 2)
	case 2:
		return sumBanks(banks, 12)
	}

	return 0
}

func sumBanks(banks []string, digits int) int {
	var cnt int

	for _, bank := range banks {
		cnt += maxJoltage(bank, digits)
	}

	return cnt
}

func maxJoltage(bank string, digits int) int {
	var jolts string

	low := 0
	high := len(bank) - (digits - 1)

	for {
		substr := bank[low:high]
		char, idx := maxIndex(substr)

		jolts += char

		if len(jolts) == digits {
			break
		}

		low += idx + 1
		high++
	}

	jolt, _ := strconv.Atoi(jolts)

	return jolt
}

func maxIndex(str string) (string, int) {
	for i := 9; i > 0; i-- {
		idx := strings.Index(str, strconv.Itoa(i))

		if idx != -1 {
			return strconv.Itoa(i), idx
		}
	}

	return "", -1
}
