package blackjack

type betSizeOnRunningCountPlayer struct {
	counter  RunningCounter
	delegate Player
	betSizes []uint
}

func (p *betSizeOnRunningCountPlayer) Balance() Balance {
	return p.delegate.Balance()
}

func (p *betSizeOnRunningCountPlayer) Decide(ph Hand, dcv CardValue) PlayDecision {
	return p.delegate.Decide(ph, dcv)
}

func (p *betSizeOnRunningCountPlayer) Surrender(ph Hand, dcv CardValue) bool {
	return p.delegate.Surrender(ph, dcv)
}

func (p *betSizeOnRunningCountPlayer) PlaceBet() uint {
	i := p.counter.RunningCount()
	if i < 0 {
		i = 0
	} else if i >= len(p.betSizes) {
		i = len(p.betSizes) - 1
	}
	return p.betSizes[i]
}

func NewBetSizeOnRunningCountPlayer(counter RunningCounter, delegate Player, betSizes []uint) Player {
	return &betSizeOnRunningCountPlayer{counter, delegate, betSizes}
}
