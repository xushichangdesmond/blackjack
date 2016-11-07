package events

type PlaceBets interface {
	BetOnMyBox(amount int)
	BoxTaken(position int) bool
	PlaceBet(position int, amount int)
	TakeSeat(position int)
}
