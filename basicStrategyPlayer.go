package blackjack

type bsPlayer struct {
	bet uint
	bal Balance
}

/* only optimized against below rules
- S17 (dealer stands on soft 17)
- dealer BJ only wins original bets
- double after splits allowed
- double any
- early surrender
not optimized in surrendering vs dealer Ace yet*/
func NewBasicStrategyPlayer(bet uint) Player {
	return &bsPlayer{bet, NewBalance()}
}

func (p *bsPlayer) BalanceReceiver() Receiver {
	return p.bal.Receiver()
}

func (p *bsPlayer) BalanceAmount() int {
	return p.bal.Amount()
}

func (p *bsPlayer) PayTo(amount uint, r Receiver) {
	p.bal.Pay(amount, r)
}

func (p *bsPlayer) Decide(ph Hand, dcv CardValue) PlayDecision {
	v := ph.Value()

	if ph.HasAce() && v < 12 {
		// soft hand

		switch {
		case v == 2:
			// pair of aces
			if dcv < 4 || dcv > 6 {
				return SplitElseHit
			}
			return SplitElseStand

		case v == 3 || v == 4:
			if dcv == 5 || dcv == 6 {
				return DoubleDownElseHit
			}
			return Hit

		case v == 5 || v == 6:
			if dcv < 4 || dcv > 7 {
				return Hit
			}
			return DoubleDownElseHit
		case v == 7:
			if dcv < 3 || dcv > 7 {
				return Hit
			}
			return DoubleDownElseHit
		case v == 8:
			if dcv == 1 || dcv > 8 {
				return Hit
			}
			if dcv == 2 || dcv == 7 || dcv == 8 {
				return Stand
			}
			return DoubleDownElseStand
		default:

			return Stand
		}
	}

	// hard hand
	switch {
	case v == 4 || v == 6:
		if dcv > 7 || dcv == 1 {
			return Hit
		}
		return SplitElseHit
	case v == 5 || v == 7:
		return Hit
	case v == 8:
		if dcv == 5 || dcv == 6 {
			return SplitElseHit
		}
		return Hit
	case v == 9:
		if dcv > 6 || dcv < 3 {
			return Hit
		}
		return DoubleDownElseHit
	case v == 10:
		if dcv == 10 || dcv == 1 {
			return Hit
		}
		return DoubleDownElseHit
	case v == 11:
		if dcv == 1 {
			return Hit
		}
		return DoubleDownElseHit
	case v == 12:
		if dcv == 2 || dcv == 3 {
			return SplitElseHit
		}
		if dcv > 6 || dcv == 1 {
			return Hit
		}
		return SplitElseStand
	case v == 13 || v == 15:
		if dcv > 6 || dcv == 1 {
			return Hit
		}
		return Stand
	case v == 14:
		if dcv > 7 || dcv == 1 {
			return Hit
		}
		if dcv == 7 {
			return SplitElseHit
		}
		return SplitElseStand
	case v == 16:
		if dcv > 6 || dcv == 1 {
			return SplitElseHit
		}
		return SplitElseStand
	case v == 18:
		if dcv == 7 || dcv == 10 || dcv == 1 {
			return Stand
		}
		return SplitElseStand
	default:
		// 17, 19, 20
		return Stand
	}
}

func (p *bsPlayer) PlaceBet() uint {
	return p.bet
}

func (p *bsPlayer) Surrender(ph Hand, dcv CardValue) bool {
	switch dcv {
	case 1:
	// TODO: put in surrender values against dealer Ace
	case 10:
		v := ph.Value()
		if v == 15 || v == 16 || v == 14 {
			return true
		}
	case 9:
		return ph.Value() == 16
	}
	return false

}
