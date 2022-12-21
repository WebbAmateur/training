package game

import "testing"

func TestCalculateScore(t *testing.T) {

	// Test estimatated == actual
	testPlayerRound := NewPlayerRound(5)
	testPlayerRound.SetActualTricks(5)

	score := testPlayerRound.CalculateRoundScore()

	summary := testPlayerRound.GetRoundSummary()

	VerifyRoundSummary(t, playerRoundSummary{playerRound: playerRound{estimatedTricks: 5, actualTricks: 5}, score: 15}, summary)

	if score != 15 {
		t.Errorf("TestCalculateScore Expected:15  Actual:%d", score)
	}

	// Test estimataed != actual

	testPlayerRound = NewPlayerRound(5)
	testPlayerRound.SetActualTricks(4)

	score = testPlayerRound.CalculateRoundScore()

	if score != 4 {
		t.Errorf("TestCalculateScore Expected:4  Actual:%d", score)
	}

}

func VerifyRoundSummary(t *testing.T, expected playerRoundSummary, actual playerRoundSummary) {

	if expected.playerRound.estimatedTricks != actual.playerRound.estimatedTricks {
		t.Errorf("Tricks   Expected:%d   Actual:%d", expected.playerRound.estimatedTricks, actual.playerRound.estimatedTricks)
	}
}
