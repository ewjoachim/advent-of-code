package d03

import (
	"aoc/boilerplate"
	"bufio"
)

type Pos struct {
	x, y int
}
type PartNum struct {
	value     int
	positions []*Pos
	symbol    *Symbol
}
type Symbol struct {
	value string
	pos   *Pos
	parts []*PartNum
}

func pt1() {
	scanner := bufio.NewScanner(boilerplate.Fetch(2023, 2))
	var l1, l2, l3 string
	for scanner.Scan() {
		l1 = l2
		l2 = l3
		l3 = scanner.Text()
		if l1 == "" {
			continue
		}




}
func pt2() {

}

func Main() {
	pt1()
	pt2()
}
