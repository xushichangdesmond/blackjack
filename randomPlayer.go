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

func (p *randomPlayer) Balance() Balance {
	return p.bal
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
