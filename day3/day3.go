package day3

import (
	"strings"
)

func PrioSum(input []string, option bool) int {
	prios := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var priosum int
	lookup := make(map[rune]bool)
	sublookup := make(map[rune]bool)

	for i, line := range input {
		if !option {
			lookup = make(map[rune]bool)
			for _, sub := range line[:len(line)/2] {
				lookup[sub] = true
			}
			for _, sub := range line[len(line)/2:] {
				if _, ok := lookup[sub]; ok {
					priosum += strings.IndexRune(prios, sub) + 1
					break
				}
			}
		} else {
			switch (i + 1) % 3 {
			case 1:
				for _, sub := range line {
					lookup[sub] = true
				}
			case 2:
				for _, sub := range line {
					if _, ok := lookup[sub]; ok {
						sublookup[sub] = true
					}
				}
			case 0:
				for _, sub := range line {
					if _, ok := sublookup[sub]; ok {
						priosum += strings.IndexRune(prios, sub) + 1
						lookup = make(map[rune]bool)
						sublookup = make(map[rune]bool)
						break
					}
				}
			}
		}
	}

	return priosum
}
