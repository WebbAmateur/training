package game

import "testing"

func TestPlayerGame(t *testing.T) {
	p := NewPlayer("Todd", 10)
	p.NewRound(5)
	p.FinishRound(5)
	score := p.CalculatePlayerScore()

	if score != 15 {
		t.Errorf("TestPlayerGame Round:0 Expected:15 Actual:%d", score)
	}

	p.NewRound(4)
	p.FinishRound(5)

	score = p.CalculatePlayerScore()

	if score != 20 {
		t.Errorf("TestPlayerGame Round:0 Expected:20 Actual:%d", score)
	}


}
