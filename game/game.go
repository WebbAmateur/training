package game

const NumberOfCards int = 52

type Game struct {
	startingTricks int
	currentTricks  int
	players        []player
	countUp        bool
	gameOver       bool
}

// NewGame creates and returns a new Game
// 2 <= playerCount <=6
// startingTricks * playerCount <= 52
func NewGame(playerCount int, startingTricks int) Game {

	if startingTricks*playerCount > NumberOfCards {
		panic("startingTricks * playerCount > 52")
	}

	// invariant: startingTricks * playerCount <= 52

	return Game{startingTricks: startingTricks, currentTricks: startingTricks, players: make([]player, 0, playerCount), countUp: false}
}

func (g Game) CurrentTricks() int {
	return g.currentTricks
}

// Game.AddPlayer adds a new player to the game and fails if the total number of players is exceeded
func (g *Game) AddPlayer(name string) {

	// Make sure that we have room for another player
	if len(g.players) == cap(g.players) {
		panic("Too many players")
	}

	g.players = append(g.players, NewPlayer(name, g.startingTricks))
}

// Game.CurrentScores returns []int containing the current score for each player in order added to the game
func (g Game) CurrentScores() []int {

	scores := make([]int, len(g.players))

	for index := range g.players {
		scores[index] = g.players[index].CalculatePlayerScore()
	}

	
	return scores
}

// *Game.NewRound accepts estimates for a new round
// *Game.NewRound panics if all players have not been registered
// *Game.NewRound panics if the wrong number of estimates is received
// *Game.NewROund panics if an estimate < 0 or estimate > current number of tricks
func (g *Game) NewRound(estimates []int) {

	if len(g.players) != cap(g.players) {
		panic("Not all players have been registered")
	}

	if len(estimates) != len(g.players) {
		panic("Wrong number of estimates provided")
	}

	for index := 0; index < len(g.players); index++ {
		// Get the right player from players and tell it to begin a round with the right element of estimates
		g.players[index].NewRound(estimates[index])
	}
}

func (g *Game) FinishRound(actuals []int) {
	for index := 0; index < len(g.players); index++ {
		// Get the right player from players and tell it to finish a round with the right element of actia;s
		g.players[index].FinishRound(actuals[index])
	}

	// Update *c.currentTricks
	if !g.countUp {
		if g.currentTricks == 1 {
			g.countUp = true
		} else {
			g.currentTricks--
		}
	}

	if g.countUp {
		if g.currentTricks == g.startingTricks {
			g.gameOver = true
		} else {
			g.currentTricks++
		}
	}
}

func (g Game) GetGameSummary() [][]playerRoundSummary{
	result := make([][]playerRoundSummary, (len(g.players)))

	for index := range g.players {
		result[index] = g.players[index].GetPlayerSummary()
	}

	return result

}
