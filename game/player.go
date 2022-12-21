package game

// inv: player.currentRound is the active round and is not included in the score
type player struct {
	name         string
	currentRound int
	rounds       []playerRound
}

func NewPlayer(playerName string, startingHandSize int) player {
	rounds := startingHandSize + startingHandSize - 1
	return player{name: playerName, rounds: make([]playerRound, rounds), currentRound: 0}
}

// *player.NewRound initializes a new player
func (p *player) NewRound(estimate int) {
	p.rounds[p.currentRound] = NewPlayerRound(estimate)
}

func (p *player) FinishRound(actual int) {
	p.rounds[p.currentRound].SetActualTricks(actual)
	p.currentRound++
}

func (p player) CalculatePlayerScore() int {
	score := 0

	for index := 0; index < p.currentRound; index++ {
		score += p.rounds[index].CalculateRoundScore()
	}
	return score
}


func (p player) GetPlayerSummary () []playerRoundSummary {
	result := make([]playerRoundSummary, len(p.rounds))
	for index := range p.rounds {
		result[index] = p.rounds[index].GetRoundSummary()
	}

	return result
}