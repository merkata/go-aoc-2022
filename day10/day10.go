package day10

import (
	"strconv"
	"strings"
)

type CPU struct {
	X              int
	Cycle          int
	SignalStrength int
}

func (c *CPU) Exec(command string, value int) {
	var cycles int
	switch command {
	case "noop":
		cycles = 1
	case "addx":
		cycles = 2
	}
	for i := 0; i < cycles; i++ {
		c.UpdateSignal()
		c.Cycle++
	}
	c.X += value
}

func (c *CPU) UpdateSignal() {
	switch c.Cycle {
	case 20:
		c.SignalStrength += c.X * 20
	case 60:
		c.SignalStrength += c.X * 60
	case 100:
		c.SignalStrength += c.X * 100
	case 140:
		c.SignalStrength += c.X * 140
	case 180:
		c.SignalStrength += c.X * 180
	case 220:
		c.SignalStrength += c.X * 220
	}
}

func SignalStrengths(input string, option bool) int {
	cpu := &CPU{X: 1, Cycle: 1}
	for _, line := range strings.Split(input, "\n") {
		var val int
		args := strings.Split(line, " ")
		if len(args) > 1 {
			val, _ = strconv.Atoi(args[1])
		}
		cpu.Exec(args[0], val)
	}
	return cpu.SignalStrength
}
