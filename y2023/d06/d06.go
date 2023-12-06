package d06

func WaysToWin(time, distance int) int {
	var speed int
	var wins int

	for i := 0; i <= time; i++ {
		timeRemaining := time - i
		if speed*timeRemaining > distance {
			wins++
		}
		speed++
	}

	return wins
}
