package d02

import (
	"aoc/boilerplate"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Set struct {
	red   uint
	green uint
	blue  uint
}

func NewSet(s string) *Set {
	var red, green, blue uint
	for _, part := range strings.Split(s, ", ") {
		valStr, _, _ := strings.Cut(part, " ")
		valInt, _ := strconv.Atoi(valStr)
		val := uint(valInt)
		switch {
		case strings.HasSuffix(part, "red"):
			red = val
		case strings.HasSuffix(part, "green"):
			green = val
		case strings.HasSuffix(part, "blue"):
			blue = val
		}
	}
	return &Set{red, green, blue}
}

func Combine(sets ...*Set) *Set {
	s := &Set{}
	for _, set := range sets {
		if set.red > s.red {
			s.red = set.red
		}
		if set.green > s.green {
			s.green = set.green
		}
		if set.blue > s.blue {
			s.blue = set.blue
		}
	}
	return s
}
func pt1() {
	// scanner := bufio.NewScanner(bytes.NewBufferString(`Game 96: 2 red, 2 green, 1 blue; 1 red, 4 green; 1 green`))
	scanner := bufio.NewScanner(boilerplate.Fetch(2023, 2))
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		before, after, _ := strings.Cut(s, ": ")
		_, gameStr, _ := strings.Cut(before, " ")
		game, _ := strconv.Atoi(gameStr)

		setStr := strings.Split(after, "; ")
		max := &Set{}
		for _, setStr := range setStr {
			set := NewSet(setStr)
			max = Combine(max, set)
		}
		cond := max.red <= 12 && max.green <= 13 && max.blue <= 14
		if cond {
			sum += game
		}
	}
	fmt.Println(sum)
}
func pt2() {
	// scanner := bufio.NewScanner(bytes.NewBufferString(`Game 96: 2 red, 2 green, 1 blue; 1 red, 4 green; 1 green`))
	scanner := bufio.NewScanner(boilerplate.Fetch(2023, 2))
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		_, after, _ := strings.Cut(s, ": ")
		// _, gameStr, _ := strings.Cut(before, " ")
		// game, _ := strconv.Atoi(gameStr)

		setStr := strings.Split(after, "; ")
		max := &Set{}
		for _, setStr := range setStr {
			set := NewSet(setStr)
			max = Combine(max, set)
		}
		sum += (int(max.red) * int(max.green) * int(max.blue))
	}
	fmt.Println(sum)
}

func main() {
	pt1()
	pt2()
}
