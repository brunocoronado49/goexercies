package practices

import "fmt"

func Score(playerOne []int, playerTwo []int) string {
	firstScore := 0
	secondScore := 0

	if len(playerOne) != len(playerTwo) {
		return "Error, array len not equals."
	}

	for i := range len(playerOne) {
		if playerOne[i] > playerTwo[i] {
			firstScore++
		} else if playerOne[i] < playerTwo[i] {
			secondScore++
		}
	}

	return fmt.Sprintf("Final Score (p1-p2): %v - %v", firstScore, secondScore)
}
