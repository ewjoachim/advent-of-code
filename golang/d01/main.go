package d01

import (
	"aoc/boilerplate"
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func pt1() {
	scanner := bufio.NewScanner(boilerplate.Fetch(2023, 1))

	sum := 0
	for scanner.Scan() {
		firstValue := 0
		lastValue := 0
		for _, c := range scanner.Text() {
			if !unicode.IsDigit(c) {
				continue
			}
			if firstValue == 0 {
				firstValue = int(c - '0')
			}
			lastValue = int(c - '0')
		}
		sum += firstValue*10 + lastValue
	}
	fmt.Println(sum)
}
func pt2() {
	digits := map[string]uint{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	scanner := bufio.NewScanner(boilerplate.Fetch(2023, 1))

	sum := int(0)
	for scanner.Scan() {
		firstValue := -1
		lastValue := -1
		line := scanner.Text()
		for i, c := range line {
			value := -1
			if unicode.IsDigit(c) {
				value = int(c - '0')
			}
			for word, v := range digits {
				if strings.HasPrefix(line[i:], word) {
					value = int(v)
					break
				}
			}
			if value != -1 {
				if firstValue < 1 {
					firstValue = value
				}
				lastValue = value
			}

		}
		sum += firstValue*10 + lastValue
	}
	fmt.Println(sum)
}

func main_01() {
	pt1()
	pt2()
}
