package blackjack

import (
	"math/rand"
)

type randomPlayer struct {
	bet uint
	bal Balance
}

func NewRandomPlayer(bet uint) Player {
	return &randomPlayer{bet, NewBalance()}
}

func (p *randomPlayer) BalanceReceiver() Receiver {
	return p.bal.Receiver()
}

func (p *randomPlayer) BalanceAmount() int {
	return p.bal.Amount()
}

func (p *randomPlayer) PayTo(amount uint, r Receiver) {
	p.bal.Pay(amount, r)
}

func (p *randomPlayer) Decide(ph Hand, dcv CardValue) PlayDecision {
	return PlayDecisions[rand.Intn(len(PlayDecisions))]
}

func (p *randomPlayer) PlaceBet() uint {
	return p.bet
}

func (p *randomPlayer) Surrender(ph Hand, dcv CardValue) bool {
	return rand.Intn(10) == 5
}
