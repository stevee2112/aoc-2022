package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"strings"
)

type Round struct {
	OpponentChoice string
	YourChoice string
}

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")


	defer input.Close()
	scanner := bufio.NewScanner(input)

	rounds := []Round{}
	for scanner.Scan() {
		valString := scanner.Text()
		parts := strings.Split(valString, " ")
		rounds = append(rounds, Round{parts[0], parts[1]})
	}

	fmt.Printf("Part 1: %d\n", part1Strategy(rounds))
	fmt.Printf("Part 2: %d\n", part2Strategy(rounds))
}

func part1Strategy(rounds []Round) int {

	total := 0
	for _, round := range rounds {
		total += pointsPerChoice(round.YourChoice) + checkWin(round)
	}

	return total
}

func part2Strategy(rounds []Round) int {

    total := 0
    for _, round := range rounds {
        switch round.YourChoice {
        case "X": // lose
			determined := needed(round)
            total += pointsPerChoice(determined.YourChoice) + checkWin(determined)
        case "Y": // draw
			determined := needed(round)
            total += pointsPerChoice(determined.YourChoice) + checkWin(determined)
        case "Z": // win
			determined := needed(round)
            total += pointsPerChoice(determined.YourChoice) + checkWin(determined)
        }
    }

    return total
}

func pointsPerChoice(choice string) int {
	switch choice {
	case "X": // Rock
		return 1
	case "Y": // Paper
		return 2
	case "Z": // Scissor
		return 3
	}

	return 0
}

func needed(round Round) Round {

	shouldChoose := ""
	them := round.OpponentChoice
	switch round.YourChoice {
	case "X": // lose
		switch them {
		case "A": // Rock
			shouldChoose = "Z" // Scissor
		case "B": // Paper
			shouldChoose = "X" // Rock
		case "C": //Scissor
			shouldChoose = "Y" // Paper"
		}
	case "Y": // draw
		switch them {
		case "A": // Rock
			shouldChoose = "X" // Rock
		case "B": // Paper
			shouldChoose = "Y" // Paper
		case "C": //Scissor
			shouldChoose = "Z" // Scissor
		}
	case "Z": // win
		switch them {
		case "A": // Rock
			shouldChoose = "Y" // Paper
		case "B": // Paper
			shouldChoose = "Z" // Scissor
		case "C": //Scissor
			shouldChoose = "X" // Rock
		}
	}

	return Round{them, shouldChoose}
}

func checkWin(round Round) int {
	you := round.YourChoice
	them := round.OpponentChoice
	winPoints := 6
	losePoints := 0
	tiePoints := 3
	switch you {
	case "X": // Rock
		switch them {
		case "A": // Rock
			return tiePoints
		case "B": // Paper
			return losePoints
		case "C": //Scissor
			return winPoints
		}
	case "Y": // Paper
		switch them {
		case "A": // Rock
			return winPoints
		case "B": // Paper
			return tiePoints
		case "C": //Scissor
			return losePoints
		}
	case "Z": // Scissor
		switch them {
		case "A": // Rock
			return losePoints
		case "B": // Paper
			return winPoints
		case "C": //Scissor
			return tiePoints
		}
	}

	return 0
}
