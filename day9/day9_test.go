package day9_test

import (
	"testing"

	"github.com/merkata/go-aoc-2022/day9"
)

type test struct {
	description string
	input       string
	option      bool
	want        int
}

func TestGetRopeTailVisits(t *testing.T) {
	tests := []test{
		{description: "part 1 validation",
			option: false,
			input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			want: 13,
		},
		{description: "part 1 solution",
			option: false,
			input: `L 2
D 2
U 1
R 1
U 2
D 1
L 1
D 1
R 2
D 2
U 1
R 2
U 2
R 2
U 2
L 2
U 2
L 2
U 1
D 1
L 1
R 1
U 2
R 2
D 1
U 2
L 1
U 2
L 2
D 2
R 1
U 2
R 2
U 1
D 1
L 1
U 2
L 2
D 1
U 1
R 2
D 2
R 1
D 1
L 1
R 2
U 1
L 1
D 2
U 2
D 1
L 1
R 1
U 1
D 1
R 2
D 2
R 1
D 2
L 1
D 2
L 1
U 2
R 1
L 1
U 2
L 2
D 2
R 2
L 1
U 2
R 1
U 2
R 2
L 2
D 2
R 1
L 1
D 1
L 1
U 2
L 2
D 2
L 2
U 1
R 2
U 1
R 2
D 2
U 1
R 1
U 2
R 2
L 1
U 2
L 1
U 2
L 1
U 1
R 2
U 2
L 1
D 2
U 1
L 2
U 2
D 2
U 1
D 1
R 1
D 2
U 2
R 2
L 2
R 2
U 1
L 2
R 3
D 2
L 1
R 1
D 3
U 3
R 2
L 3
U 1
L 3
U 2
L 2
U 3
D 2
R 1
U 1
R 3
U 2
D 2
U 3
D 3
L 2
U 2
L 3
U 1
R 1
D 3
L 1
R 2
U 1
L 2
D 1
U 2
L 2
U 1
D 1
U 1
R 3
U 3
R 2
U 1
D 1
U 2
D 1
R 1
U 3
R 1
U 3
R 1
U 2
D 1
L 1
D 1
L 1
D 1
U 1
L 2
R 2
U 2
L 3
U 1
D 1
R 3
U 1
L 2
R 1
D 1
R 3
D 2
U 2
D 1
R 1
U 1
D 1
U 1
D 3
U 1
L 3
D 1
R 2
U 2
R 2
U 1
R 2
D 2
L 3
U 1
D 1
L 3
U 3
R 1
U 3
D 3
L 3
D 3
L 2
R 2
U 2
R 1
U 2
D 2
U 2
L 2
R 1
U 3
R 4
L 2
R 4
D 2
R 2
L 2
R 2
D 1
U 2
D 3
R 1
U 1
D 4
R 4
L 1
U 1
D 2
L 2
R 3
U 2
L 3
R 4
U 3
L 3
D 4
R 3
L 2
U 3
D 1
L 2
U 3
L 1
D 2
R 1
D 4
L 2
R 4
U 4
L 2
R 2
D 4
L 3
U 3
L 2
U 2
L 4
D 2
R 4
D 1
L 4
U 4
R 1
U 4
D 2
R 2
U 3
L 1
R 4
U 3
R 3
U 4
L 2
D 3
L 4
R 2
L 4
D 4
U 1
L 2
D 4
R 1
U 3
R 1
U 4
L 1
D 4
U 2
L 4
U 1
R 2
U 4
L 2
D 3
L 2
U 4
D 2
L 3
D 1
U 4
D 1
L 4
U 3
R 3
L 2
R 2
U 2
D 2
L 1
U 4
R 2
L 3
R 1
U 4
L 2
U 4
L 2
U 3
L 2
R 3
U 3
L 2
D 4
R 3
U 2
L 3
D 3
R 4
U 1
R 3
D 5
U 1
L 1
R 5
D 2
L 3
U 5
D 3
L 3
D 1
U 3
R 2
U 5
L 2
D 2
L 5
U 2
D 5
R 1
U 1
R 5
L 4
D 4
R 2
U 5
D 5
L 1
R 2
U 4
L 2
D 1
L 5
D 1
R 3
L 2
R 4
D 1
R 4
U 4
L 1
U 2
D 5
U 4
L 3
D 3
R 3
D 5
L 2
U 5
R 3
L 5
R 4
U 4
R 3
L 2
D 4
L 2
U 1
R 3
L 2
R 3
L 5
U 1
L 2
R 1
D 3
R 4
L 2
D 2
R 4
D 3
L 5
D 2
L 4
R 3
L 3
D 2
R 5
U 4
R 4
U 3
L 4
U 1
R 4
D 4
U 1
L 3
U 1
L 1
R 4
D 1
U 3
D 2
R 3
L 3
U 2
R 4
D 4
L 1
R 5
U 1
L 2
U 2
R 4
U 4
R 2
D 2
U 3
L 5
D 4
L 5
D 1
L 2
D 4
R 6
L 3
D 5
R 3
L 5
R 5
D 1
R 3
L 5
R 5
L 6
R 4
L 2
R 5
D 3
U 1
R 4
D 2
R 2
L 6
U 2
R 5
U 5
L 3
R 6
L 6
D 3
L 1
R 1
L 1
D 5
L 4
U 3
D 5
L 1
U 6
R 5
D 6
U 1
D 5
L 3
D 5
R 6
U 2
D 1
L 4
D 5
R 2
U 5
R 4
L 1
R 2
L 1
R 2
L 6
U 2
D 3
R 5
D 3
L 4
R 4
D 1
R 6
D 1
R 4
L 3
D 6
L 5
D 2
R 1
L 3
R 2
L 6
D 4
U 5
R 6
U 3
D 2
R 4
L 6
U 4
L 2
U 2
L 1
U 5
D 1
L 2
U 1
D 2
L 1
R 4
U 4
L 4
D 4
R 3
D 6
L 2
U 4
R 1
L 2
U 3
L 6
U 6
R 3
D 7
U 4
R 6
L 6
D 7
L 5
U 1
R 2
U 7
D 5
L 7
R 2
U 3
R 7
D 5
L 2
U 2
R 3
U 5
R 3
D 4
R 4
U 3
L 6
R 2
D 7
U 6
D 1
U 3
R 3
D 4
R 6
U 4
L 3
D 3
R 4
L 2
D 1
R 3
U 6
R 5
U 3
L 5
R 5
D 4
U 3
L 7
U 7
D 3
U 1
D 1
U 1
D 3
L 5
R 4
U 6
R 7
D 2
L 3
R 2
U 3
D 2
L 5
D 7
R 6
D 1
U 5
D 7
U 6
D 4
L 6
D 4
L 2
D 7
R 6
L 3
D 2
L 3
D 2
U 6
L 1
R 2
U 5
L 5
R 2
U 6
R 3
L 7
R 4
D 7
L 7
D 6
L 4
D 2
R 6
L 1
D 3
U 6
D 7
U 1
L 5
D 7
U 3
R 3
L 7
D 2
L 7
U 5
D 4
R 3
D 8
R 7
U 1
R 8
U 4
D 5
R 6
L 5
D 5
L 7
R 7
U 4
D 2
L 7
U 8
D 6
L 5
D 2
L 6
U 6
R 5
U 5
L 6
U 5
L 3
U 4
D 7
R 3
L 4
R 7
D 1
R 5
D 2
R 8
L 8
U 8
L 6
U 5
L 1
U 6
R 2
U 4
L 5
R 2
D 8
L 1
U 4
R 6
U 1
L 4
R 7
L 8
R 5
U 5
R 3
L 3
U 2
D 8
L 8
R 4
U 3
D 3
L 1
R 8
U 3
L 8
D 8
U 7
D 8
L 3
U 5
D 7
L 7
U 5
L 3
R 6
D 7
R 1
D 7
L 5
U 3
D 5
U 1
L 2
D 4
U 5
R 1
L 6
U 3
L 1
D 6
R 1
L 1
R 3
L 4
U 6
R 3
U 8
L 2
U 3
R 1
D 5
R 2
U 2
D 8
L 3
D 2
R 8
D 1
L 3
U 1
D 8
U 1
L 2
D 4
R 3
L 4
D 3
L 2
D 4
L 6
U 6
D 9
U 5
R 9
D 1
R 8
L 9
R 7
D 5
U 1
L 1
R 8
U 3
R 4
L 2
R 1
D 6
U 5
D 6
L 9
R 2
D 2
R 8
D 2
R 3
U 1
L 1
U 1
L 1
U 3
D 9
R 2
U 1
D 1
R 6
D 7
L 9
U 4
L 7
U 8
D 8
L 3
D 9
L 4
D 8
U 9
L 9
D 7
U 7
D 2
L 4
D 5
L 5
U 6
D 8
L 5
U 9
L 4
U 3
D 4
R 9
U 4
D 7
L 6
R 7
L 5
U 5
L 8
D 3
L 5
U 7
D 6
U 3
R 1
D 8
U 3
L 8
U 9
R 1
U 1
L 3
R 6
U 5
R 7
U 4
R 4
L 8
U 4
L 9
R 7
D 6
R 8
U 9
D 8
L 3
R 7
U 1
L 9
R 2
L 3
D 6
R 9
D 1
L 3
D 4
U 2
D 9
R 9
U 3
R 2
D 3
U 7
L 4
D 6
R 3
D 8
U 4
R 8
U 6
R 5
L 5
U 2
D 1
R 5
U 9
D 3
L 5
R 8
U 2
D 9
R 5
U 4
L 7
U 4
D 6
U 7
D 1
L 6
U 7
D 6
L 3
D 8
U 1
L 1
U 9
L 2
U 4
D 2
R 8
U 3
R 8
L 4
U 3
L 7
U 7
D 7
R 7
L 2
U 6
R 1
U 7
D 6
L 1
R 5
L 7
R 1
D 7
U 9
D 9
L 1
D 10
U 10
L 2
R 9
L 1
R 3
U 8
L 1
R 2
L 2
U 2
D 1
U 5
R 5
D 7
L 6
D 2
L 10
U 10
D 4
L 10
R 10
U 5
L 7
R 10
D 1
L 8
D 3
U 7
R 2
D 7
U 4
L 3
R 4
L 10
U 3
L 5
R 4
D 3
R 9
L 9
R 6
U 5
L 5
R 11
U 6
L 11
R 9
D 5
U 7
L 8
R 5
U 11
D 6
L 9
D 5
L 2
D 2
R 4
L 1
D 10
L 9
R 6
D 2
L 11
D 4
R 5
D 2
R 8
D 3
L 1
R 11
D 1
L 11
D 7
L 5
U 7
D 4
L 7
U 11
L 3
U 7
L 2
D 6
R 1
D 3
R 1
L 1
D 11
L 6
U 1
R 9
U 8
D 8
L 2
R 6
U 10
R 7
L 5
U 1
R 4
L 2
D 11
L 10
U 7
L 10
D 1
R 6
U 2
D 7
R 7
D 2
R 10
U 8
D 5
U 2
L 11
R 9
D 7
U 3
R 6
U 6
D 7
U 6
L 2
R 11
L 10
D 7
U 10
D 3
R 10
L 4
D 8
R 5
L 5
D 9
U 8
R 3
D 10
R 11
U 5
D 11
L 2
D 8
L 7
U 4
D 5
U 9
D 3
U 11
D 4
L 1
U 4
L 3
R 7
L 12
D 10
U 8
L 6
D 11
U 5
L 2
D 6
L 5
U 8
L 4
U 6
R 9
D 2
L 6
U 12
R 2
L 10
R 7
U 7
L 5
R 6
U 5
R 7
D 6
R 4
U 6
L 4
U 4
L 4
U 3
D 3
L 9
D 12
R 10
U 3
L 1
U 4
L 8
R 4
D 6
L 2
D 7
L 9
U 3
L 11
R 9
D 8
L 1
D 11
R 8
L 6
U 7
R 5
U 3
D 4
R 4
U 5
L 6
U 10
D 2
R 8
U 7
D 4
L 11
R 9
U 9
D 12
U 8
R 4
U 5
D 11
L 5
R 5
D 2
L 12
U 4
L 1
R 9
L 5
D 1
R 2
L 2
D 7
L 12
D 2
R 7
U 10
L 1
U 2
L 8
D 5
L 7
U 8
L 8
U 10
D 7
R 8
D 7
R 9
L 7
D 1
R 3
D 4
R 11
D 3
L 5
D 7
U 8
L 2
D 10
L 12
R 7
U 13
R 4
U 10
L 10
D 6
U 7
D 12
U 2
D 8
L 1
R 5
D 9
L 1
R 11
L 2
D 7
L 13
U 1
L 8
D 2
U 12
L 6
R 10
L 10
R 9
U 3
R 1
L 10
D 2
L 4
D 8
U 7
L 11
D 8
U 1
D 1
L 2
D 2
L 5
U 6
D 1
U 6
R 2
L 2
D 5
U 2
R 10
L 4
D 3
L 11
D 11
U 4
D 4
R 12
D 11
U 2
R 11
U 3
R 8
U 8
L 5
R 12
U 1
L 10
R 11
L 3
U 6
D 2
R 7
D 13
L 11
R 8
L 8
R 1
L 5
D 13
R 8
L 12
D 13
R 10
U 5
D 13
U 6
L 2
D 2
U 7
D 4
L 4
R 5
L 12
U 6
D 8
U 7
D 12
R 11
L 11
D 9
R 4
D 3
L 5
R 6
D 9
R 3
U 10
R 10
U 3
D 11
L 3
D 5
R 14
U 7
R 3
D 1
L 4
U 6
D 6
R 10
L 12
R 10
L 4
U 10
R 8
D 14
L 6
U 7
D 4
L 9
R 10
D 10
L 4
R 6
U 8
D 10
L 12
U 11
L 7
R 1
U 1
R 12
U 7
D 3
U 4
R 3
D 6
U 2
R 2
U 14
D 10
R 5
L 12
U 5
D 6
R 7
L 13
D 11
R 6
D 11
R 6
D 6
U 10
L 14
D 4
R 14
D 7
U 12
L 3
U 2
D 3
R 1
D 3
U 2
D 13
L 14
R 13
U 3
D 12
L 14
U 12
D 6
R 2
U 6
L 11
D 3
L 7
D 7
R 11
D 11
R 10
L 2
U 12
L 14
U 11
R 12
L 4
D 2
R 13
L 1
U 7
L 2
D 7
L 14
R 11
U 1
L 7
U 11
L 6
U 1
D 11
L 14
R 4
L 1
D 14
U 7
L 6
R 3
L 2
U 10
D 3
R 4
U 7
R 11
U 7
R 13
U 2
L 12
R 9
L 15
R 4
U 5
D 7
L 10
R 9
D 5
L 14
U 9
R 3
L 15
R 5
L 10
D 5
U 3
D 9
U 13
D 15
R 10
L 9
R 9
U 4
D 2
L 3
R 5
L 11
R 7
L 15
D 12
L 13
D 4
R 2
U 14
D 9
L 1
D 1
L 6
U 9
D 14
U 2
L 1
R 10
U 14
R 4
D 11
R 6
D 11
U 9
R 2
L 6
R 10
L 4
D 7
R 11
U 5
L 2
D 5
U 6
R 8
U 4
R 9
D 15
R 8
L 10
U 5
L 4
U 7
L 3
U 5
R 2
U 13
L 9
U 4
D 13
R 1
D 15
U 4
L 7
D 12
L 8
U 6
L 10
U 13
R 12
D 13
U 6
D 13
U 2
L 13
U 7
L 8
U 6
D 2
U 13
L 2
D 13
R 3
D 2
L 11
D 13
U 1
D 12
L 9
U 2
R 9
U 16
D 6
L 16
D 15
R 7
U 3
D 11
L 8
U 3
L 15
R 7
D 14
L 8
R 7
U 12
R 11
U 2
R 13
D 9
L 4
U 14
L 6
D 7
L 10
U 14
D 8
L 6
D 3
L 6
D 12
U 13
R 9
L 6
U 9
R 9
D 13
R 13
U 13
L 4
D 1
R 8
L 1
R 15
U 15
R 14
D 9
L 5
D 2
L 14
U 14
L 15
R 10
U 11
L 7
R 12
U 7
L 7
D 9
U 1
L 11
R 11
D 13
U 7
R 7
U 7
R 8
D 12
U 11
D 14
U 16
L 5
R 8
D 2
L 10
D 14
R 16
D 13
U 14
L 6
R 11
D 4
R 15
D 2
R 3
L 2
U 8
L 4
D 5
U 3
L 11
U 7
R 14
U 4
L 5
D 10
U 16
L 1
U 5
R 11
D 16
U 5
D 11
R 5
D 9
R 3
U 13
D 6
L 10
D 2
U 4
D 7
U 1
R 10
D 6
U 12
L 2
D 8
U 5
L 5
U 7
L 3
U 16
D 12
L 12
R 16
L 5
R 17
L 7
U 17
D 15
R 14
L 7
U 16
R 7
D 7
R 14
D 6
L 7
U 6
R 5
D 14
L 8
R 3
L 1
U 12
L 8
D 11
R 3
D 11
L 16
R 4
U 2
R 17
U 10
R 3
U 13
R 6
D 16
U 15
R 3
U 7
D 6
U 2
L 12
D 11
L 6
D 9
R 7
U 10
R 14
D 10
R 10
D 17
R 9
L 11
R 7
D 14
L 10
U 17
R 5
D 12
U 11
D 11
R 7
D 12
U 10
D 10
U 8
D 11
R 8
U 9
L 17
D 15
R 2
D 14
L 10
U 2
L 9
R 15
L 8
D 9
L 7
R 10
D 6
L 15
U 1
L 11
R 14
L 7
R 17
U 17
L 8
U 16
D 4
R 9
L 13
R 5
L 1
D 14
U 7
L 8
R 10
U 10
D 16
U 6
R 13
D 12
U 10
L 3
U 2
R 7
D 12
R 16
D 1
L 14
R 9
D 18
R 1
L 17
R 9
U 16
L 5
U 12
R 7
D 6
L 17
U 3
R 5
U 13
D 9
U 12
L 13
D 5
R 1
U 7
L 4
R 13
D 1
L 15
U 11
L 5
D 2
R 12
D 12
U 5
L 1
R 5
U 3
R 5
L 9
R 15
U 3
D 18
L 10
R 2
D 7
R 13
L 8
R 5
L 2
R 15
D 12
L 14
U 15
L 8
U 13
R 13
L 15
R 16
D 10
L 15
U 1
R 1
D 3
U 17
D 2
L 5
R 17
U 6
R 10
L 3
D 11
R 13
D 5
U 3
D 9
U 10
R 12
D 10
R 15
U 11
L 6
R 3
U 4
D 7
R 12
L 5
R 7
D 17
L 14
U 9
D 7
U 2
R 1
D 13
R 3
D 12
L 3
U 17
D 1
U 6
L 8
U 17
D 1
R 6
L 8
R 8
U 5
L 12
D 8
R 5
L 11
U 17
L 3
U 11
L 1
R 15
D 15
L 10
D 10
L 19
R 11
L 4
D 11
L 13
U 4
D 6
L 3
U 11
R 3
L 9
D 4
U 10
L 14
D 10
R 11
L 1
R 1
D 8
R 19
D 6
L 16
D 1
U 13
R 19
U 19
D 3
R 8
D 16
R 2
D 2
L 19
D 7
R 2
U 4
L 5
R 2
U 9
L 5
D 13
L 5
R 14
L 13
D 5
U 10
R 12
L 11
D 18
U 7
L 12
R 11
D 16
U 15
L 12
R 12
D 1
R 2
D 6
L 15
D 8
R 19
U 9
R 12
L 10
R 3
D 19
R 15
U 12
R 6
L 3
D 17
R 2
U 11
D 12
L 15
D 16
R 18
D 16
U 12
R 16
U 17
R 15
L 10
R 4
D 12
L 5
R 11
L 1
U 11
R 3
D 5`,
			want: 5981,
		},
	}
	for _, tc := range tests {
		got := day9.GetRopeTailVisits(tc.input, tc.option)
		if tc.want != got {
			t.Errorf("failed test: %s, wanted %d got %d\n", tc.description, tc.want, got)
		}
	}
}
