package day2

import (
	"strings"
)

type Play struct {
	Opponent string
	Player   string
}

var RPSpoints = map[string]int{"A": 1, "B": 2, "C": 3}
var RPStranslate = map[string]string{"X": "A", "Y": "B", "Z": "C"}

func RPSOutput(input []string, outcome bool) int {

	plays := []Play{}

	for _, line := range input {
		play := strings.Split(line, " ")
		plays = append(plays, Play{Opponent: play[0], Player: play[1]})
	}

	var score int
	for _, play := range plays {
		if outcome {
			score += calculateRound(play.Opponent, play.Player)
		} else {
			score += calculateHand(play.Opponent, RPStranslate[play.Player]) + RPSpoints[RPStranslate[play.Player]]
		}
	}

	return score
}

func calculateHand(opponent, player string) int {
	if opponent == player {
		return 3
	}
	if opponent == "A" && player == "B" {
		return 6
	}
	if opponent == "B" && player == "C" {
		return 6
	}
	if opponent == "C" && player == "A" {
		return 6
	}
	return 0
}

func calculateRound(opponent, player string) int {
	var result int
	switch player {
	//lose
	case "X":
		if opponent == "A" {
			result = RPSpoints["C"] + 0
		} else if opponent == "B" {
			result = RPSpoints["A"] + 0
		} else if opponent == "C" {
			result = RPSpoints["B"] + 0
		}
		//draw
	case "Y":
		result = RPSpoints[opponent] + 3
		//win
	case "Z":
		if opponent == "A" {
			result = RPSpoints["B"] + 6
		} else if opponent == "B" {
			result = RPSpoints["C"] + 6
		} else if opponent == "C" {
			result = RPSpoints["A"] + 6
		}
	}
	return result
}
