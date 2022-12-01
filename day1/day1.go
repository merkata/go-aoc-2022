package day1

import (
	"fmt"
	"strconv"
)

func CalculateCalories(input []string, sum int) int {

	var elf int
	max := make([]int, sum)

	for _, line := range input {
		if line != "" {
			calories, _ := strconv.Atoi(line)
			elf += calories
		} else {
			if compare(elf, max) {
				insert(elf, max)
			}
			elf = 0
		}
	}
	if compare(elf, max) {
		insert(elf, max)
	}

	return total(max)
}

func compare(elf int, elems []int) bool {
	fmt.Printf("comparing elf calories %d against max of %v\n", elf, elems)
	for _, elem := range elems {
		if elf > elem {
			return true
		}
	}
	return false
}

func insert(elf int, elems []int) {
	var pos int
	lowest := elems[pos]
	for idx, elem := range elems {
		if lowest > elem {
			lowest, pos = elem, idx
		}
	}
	elems[pos] = elf
}

func total(elems []int) int {
	var sum int
	for _, elem := range elems {
		sum += elem
	}
	return sum
}
