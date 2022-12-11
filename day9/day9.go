package day9

import (
	"strconv"
	"strings"
)

type Rope struct {
	Visits     map[[2]int]bool
	HeadCoords [2]int
	TailCoords [2]int
}

type Instruction struct {
	Direction string
	Moves     int
}

func (r *Rope) Move(direction string) {
	switch direction {
	case "L":
		r.HeadCoords[1]--
		if r.HeadCoords[0] == r.TailCoords[0] { // on the same line
			if r.TailCoords[1]-r.HeadCoords[1] > 1 {
				r.TailCoords[1]--
			}
		} else {
			if r.TailCoords[1]-r.HeadCoords[1] > 1 {
				r.TailCoords[1]--
				if r.HeadCoords[0] > r.TailCoords[0] {
					r.TailCoords[0]++
				} else {
					r.TailCoords[0]--
				}
			}
		}
	case "R":
		r.HeadCoords[1]++
		if r.HeadCoords[0] == r.TailCoords[0] { // on the same line
			if r.HeadCoords[1]-r.TailCoords[1] > 1 {
				r.TailCoords[1]++
			}
		} else { // different lines
			if r.HeadCoords[1]-r.TailCoords[1] > 1 {
				r.TailCoords[1]++
				if r.HeadCoords[0] > r.TailCoords[0] {
					r.TailCoords[0]++
				} else {
					r.TailCoords[0]--
				}
			}
		}
	case "U":
		r.HeadCoords[0]++
		if r.HeadCoords[1] == r.TailCoords[1] { // on the same line
			if r.HeadCoords[0]-r.TailCoords[0] > 1 {
				r.TailCoords[0]++
			}
		} else {
			if r.HeadCoords[0]-r.TailCoords[0] > 1 {
				r.TailCoords[0]++
				if r.HeadCoords[1] > r.TailCoords[1] {
					r.TailCoords[1]++
				} else {
					r.TailCoords[1]--
				}
			}
		}
	case "D":
		r.HeadCoords[0]--
		if r.HeadCoords[1] == r.TailCoords[1] { // on the same line
			if r.TailCoords[0]-r.HeadCoords[0] > 1 {
				r.TailCoords[0]--
			}
		} else {
			if r.TailCoords[0]-r.HeadCoords[0] > 1 {
				r.TailCoords[0]--
				if r.HeadCoords[1] > r.TailCoords[1] {
					r.TailCoords[1]++
				} else {
					r.TailCoords[1]--
				}
			}
		}
	}
	r.Visits[r.TailCoords] = true
}

func GetRopeTailVisits(input string, option bool) int {
	visits := make(map[[2]int]bool)
	visits[[2]int{0, 0}] = true
	rope := &Rope{Visits: visits, HeadCoords: [2]int{0, 0}, TailCoords: [2]int{0, 0}}
	instructions := []Instruction{}
	for _, line := range strings.Split(input, "\n") {
		parsed := strings.Split(line, " ")
		direction := parsed[0]
		moves, _ := strconv.Atoi(parsed[1])
		instructions = append(instructions, Instruction{Direction: direction, Moves: moves})
	}
	for _, instruction := range instructions {
		for i := instruction.Moves; i > 0; i-- {
			rope.Move(instruction.Direction)
		}
	}
	return len(rope.Visits)
}
