package game

const CORRECT_BONUS int = 10

// playerRound stores the estimate and actual for a single player for a single round
// playerRound.estimatedTricks is always correct
type playerRound struct {
	estimatedTricks int
	actualTricks    int
}

type playerRoundSummary struct {
	playerRound
	score int
}

// NewPlayerRound constructs a playerRound and sets its estimatedTricks field
func NewPlayerRound(estimate int) playerRound {
	return playerRound{estimatedTricks: estimate, actualTricks: 0}
}

// playerRound.CalculateRoundScore calculates the score from pr.estimatedTricks and pr.actualTricks
func (pr playerRound) CalculateRoundScore() int {
	score := pr.actualTricks

	if pr.actualTricks == pr.estimatedTricks {
		score += CORRECT_BONUS
	}

	return score
}

// playerROUnd.SetActualTricks persistently sets pr.actualTricks
func (pr *playerRound) SetActualTricks(actual int) {
	pr.actualTricks = actual
}

func (pr playerRound) GetRoundSummary() playerRoundSummary {

	return playerRoundSummary{playerRound: pr, score: pr.CalculateRoundScore()}
}
