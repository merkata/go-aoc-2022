package day8

import (
	"strconv"
	"strings"
)

type Trees []int
type Forrest []Trees

func VisibleTrees(input string, option bool) int {
	forrest := Forrest{}
	for _, line := range strings.Split(input, "\n") {
		trees := Trees{}
		for _, tree := range line {
			tree, _ := strconv.Atoi(string(tree))
			trees = append(trees, tree)
		}
		forrest = append(forrest, trees)
	}
	visible := 0
	score := 0
	for i := 0; i < len(forrest); i++ {
		for j := 0; j < len(forrest[0]); j++ {
			isVisible, scenicScore := check(forrest, i, j)
			visible += isVisible
			if score < scenicScore {
				score = scenicScore
			}
		}
	}
	if !option {
		return visible
	}
	return score
}

func check(forrest Forrest, i, j int) (int, int) {
	val := forrest[i][j]

	scores := []int{}
	score := 0
	visible := 4
	for m := i + 1; m < len(forrest); m++ {
		score++
		if forrest[m][j] >= val {
			visible--
			break
		}
	}
	scores = append(scores, score)
	score = 0
	for m := i - 1; m >= 0; m-- {
		score++
		if forrest[m][j] >= val {
			visible--
			break
		}
	}
	scores = append(scores, score)
	score = 0
	for m := j + 1; m < len(forrest[0]); m++ {
		score++
		if forrest[i][m] >= val {
			visible--
			break
		}
	}
	scores = append(scores, score)
	score = 0
	for m := j - 1; m >= 0; m-- {
		score++
		if forrest[i][m] >= val {
			visible--
			break
		}
	}
	scores = append(scores, score)
	scenicScore := mul(scores)
	if visible > 0 {
		return 1, scenicScore
	}
	return 0, scenicScore
}

func mul(a []int) int {
	result := 1
	for _, val := range a {
		result *= val
	}
	return result
}
