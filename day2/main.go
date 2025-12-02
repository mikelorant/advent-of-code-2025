package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ID struct {
	First int
	Last  int
}

func main() {
	do := Task("input1.txt", 1)
	log.Println("Part 1:", do)

	do = Task("input1.txt", 2)
	log.Println("Part 2:", do)
}

func Task(file string, part int) int {
	ids, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return addInvalidIDs(ids, false)
	case 2:
		return addInvalidIDs(ids, true)
	}

	return 0
}

func addInvalidIDs(ids []ID, hasRepeat bool) int {
	var cnt int

	for _, id := range ids {
		cnt += checkRange(id, hasRepeat)
	}

	return cnt
}

func checkRange(id ID, hasRepeat bool) int {
	var cnt int

	for i := id.First; i <= id.Last; i++ {
		id := strconv.Itoa(i)
		div := len(id) / 2

		switch hasRepeat {
		case true:
			for div > 0 {
				if !isValid(id, div) {
					cnt += i

					break
				}

				div--
			}
		default:
			if len(id)%2 == 1 {
				continue
			}

			if !isValid(id, div) {
				cnt += i
			}
		}
	}

	return cnt
}

func isValid(id string, div int) bool {
	if len(id)%div != 0 {
		return true
	}

	if isRepeatN(id, id[:div], div) {
		return false
	}

	return true
}

func isRepeatN(str, subStr string, i int) bool {
	return strings.Count(str, subStr) == len(str)/i
}

func (id ID) String() string {
	return fmt.Sprintf("%d-%d", id.First, id.Last)
}
