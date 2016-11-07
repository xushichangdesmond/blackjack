package blackjack

type Player interface {
	Balance() Balance

	Decide(ph Hand, dcv CardValue) PlayDecision

	Surrender(ph Hand, dcv CardValue) bool

	//PlaceBet asks the player how many betting units he/she wishes to bet. Negative values signal a leaving of the table
	PlaceBet() uint
}
