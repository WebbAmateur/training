package game

import "testing"

func TestGame(t *testing.T) {

	game := NewGame(3, 10)

	game.AddPlayer("Todd")
	game.AddPlayer("Nick")
	game.AddPlayer("Thomas")

	VerifyCount(t, 10, game.CurrentTricks())

	game.NewRound([]int{5, 4, 4})
	game.FinishRound([]int{0, 5, 4})

	VerifyScores(t, []int{0, 5, 14}, game.CurrentScores())
	VerifyPlayerSummary(t, 3, 1, game.GetGameSummary())

	VerifyCount(t, 9, game.CurrentTricks())

	game.NewRound([]int{2, 3, 4})
	game.FinishRound([]int{2, 3, 4})

	VerifyScores(t, []int{12, 18, 28}, game.CurrentScores())
	VerifyPlayerSummary(t, 3, 2, game.GetGameSummary())

	VerifyCount(t, 8, game.CurrentTricks())
}

func VerifyCount(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Count Expected:%d  Actual:%d", expected, actual)
	}
}

func VerifyScores(t *testing.T, expected []int, actual []int) {
	for index := range expected {
		if expected[index] != actual[index] {
			t.Errorf("Score:%d Expected:%d  Actual:%d", index, expected[index], actual[index])
		}
	}
}

// We test the contents of playerRoundSummary elsewhere.
// Here we verify player count and round cound
func VerifyPlayerSummary(t *testing.T, players int, round int, gameSummary [][]playerRoundSummary) {

	// Verify that the length of [][]playerRoundSummary is correct
	if players != len(gameSummary) {
		t.Errorf("Players in Summary  Expected:%d   Actual:%d", players, len(gameSummary))
	}

	// Verify that the length of each []playerRoundSummary is correct
	for playerIndex := range gameSummary {
		if len(gameSummary[playerIndex]) != round {
			t.Errorf("Rounds for Player %d  Expected: %d   Actual: %d", playerIndex, round, len(gameSummary[playerIndex]))
		}
	}
}
