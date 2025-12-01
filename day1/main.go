package main

import (
	"log"
)

func main() {
	do := Task("input1.txt", 1)
	log.Println("Part 1:", do)

	do = Task("input1.txt", 2)
	log.Println("Part 2:", do)
}

func Task(file string, part int) int {
	rots, err := parse(file)
	if err != nil {
		log.Fatal("unable to parse file:", err.Error())
	}

	switch part {
	case 1:
		return password(rots, false)
	case 2:
		return password(rots, true)
	}

	return 0
}

func password(rots []int, countClicks bool) int {
	var i int

	dial := 50

	for _, rot := range rots {
		if countClicks {
			i += clicks(dial, rot)
		}

		dial = rotate(dial, rot)

		if dial == 0 {
			i++
		}
	}

	return i
}

func rotate(dial, rot int) int {
	dial = (dial + rot) % 100

	if dial < 0 {
		dial = 100 + dial
	}

	return dial
}

func clicks(dial, rot int) int {
	// Calculate full revolutions
	cs := revolutions(rot)

	start := dial
	end := rotate(dial, rot)

	// If we start or end on 0 just return revolutions
	if (start == 0) || (end == 0) {
		return cs
	}

	// end - start is positive
	// rot is positive
	// rotated right without crossing zero
	if (end-start) > 0 && rot > 0 {
		return cs
	}

	// end - start is negative
	// rot is negative
	// rotated left without crossing zero
	if (end-start) < 0 && rot < 0 {
		return cs
	}

	// we rotated passed zero so add it to the revolutions
	return cs + 1
}

func revolutions(rot int) int {
	r := rot / 100

	if r < 0 {
		r = -r
	}

	return r
}
