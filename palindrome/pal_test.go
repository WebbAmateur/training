package palindrome

import "testing"

func TestPalindrome(t *testing.T) {

	//verifySingle(t, "test", false)
	verifySingle(t, "A nut for a jar of tuna", true)
}

func verifySingle(t *testing.T, candidate string, expected bool) {
	actual := IsPalindrome(candidate)
	if actual != expected {
		t.Errorf("Candidate:%s Expected:%t Actual:%t", candidate, expected, actual)
	}
}
