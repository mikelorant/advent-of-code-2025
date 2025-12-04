package main

import (
	"fmt"
	"strings"
)

func (w Warehouse) String() string {
	var ss []string

	for _, row := range w {
		ss = append(ss, fmt.Sprint(row))
	}

	return strings.Join(ss, "\n")
}

func (r Row) String() string {
	var ss []string

	for _, cell := range r {
		ss = append(ss, fmt.Sprint(cell))
	}

	return strings.Join(ss, "")
}

func (c Cell) String() string {
	return fmt.Sprint(c.Kind)
}

func (k Kind) String() string {
	KindName := map[Kind]string{
		Empty: ".",
		Paper: "@",
	}

	return KindName[k]
}
