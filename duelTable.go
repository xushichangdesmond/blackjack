package blackjack

import (
	"context"
	"errors"
	"sync"

	"github.com/golang/glog"
)

var ErrNoEmptyBox = errors.New("No empty bos available")

type DuelingTable interface {
	Table

	Join(Player) error

	SubscribeCardListener(CardListener)
	SubscribeShuffleListener(ShuffleListener)
}

// 1 player (1 box) vs the dealer
type duelingTable struct {
	r *Rules
	s Shoe

	dh  *Hand
	phs []HandWithBet

	p          chan Player
	havePlayer bool
	pJoinMutex sync.Mutex

	bet    uint
	rounds uint64
	bal    Balance

	shoesSofar int

	burnt *Hand

	shuffler Shuffler

	cardListeners    listeners
	shuffleListeners listeners
}

func (t *duelingTable) SubscribeCardListener(l CardListener) {
	t.cardListeners.Subscribe(l)
}

func (t *duelingTable) SubscribeShuffleListener(l ShuffleListener) {
	t.shuffleListeners.Subscribe(l)
}

func (t *duelingTable) Rules() *Rules {
	return t.r
}

func (t *duelingTable) Join(p Player) error {
	t.pJoinMutex.Lock()
	defer t.pJoinMutex.Unlock()

	if t.havePlayer {
		return ErrNoEmptyBox
	}

	t.havePlayer = true
	t.p <- p
	return nil
}

func (t *duelingTable) Open(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case p := <-t.p:
				t.bet = t.askForBet(p)

				if t.bet < 0 {
					t.pJoinMutex.Lock()
					t.havePlayer = false
					t.pJoinMutex.Unlock()
					if glog.V(50) {
						glog.Infoln("Player ", p, " has left the table")
					}
					continue
				}
				go func() {
					defer func() {
						t.p <- p
					}()
					if glog.V(50) {
						glog.Infoln("Player has placed bet of ", t.bet, " units")
					}
					t.playOneRound(p)
					t.shuffleIfRequired(p)
				}()
			}
		}
	}()
}

func (t *duelingTable) askForBet(p Player) uint {
	b := p.PlaceBet()
	if b < t.r.MinBetUnits {
		return t.r.MinBetUnits
	}
	if b > t.r.MaxBetUnits {
		return t.r.MaxBetUnits
	}
	return b
}

func (t *duelingTable) shuffleIfRequired(p Player) {
	if len(t.s) < t.r.ShufflingPoint {
		t.s = t.shuffler(t.s, *t.burnt)
		for _, l := range t.shuffleListeners.ls {
			(l.(ShuffleListener))()
		}
		*t.burnt = Hand([]Card{})
		t.shoesSofar++
		if t.shoesSofar%100000 == 0 && glog.V(40) {
			glog.Infoln("Shuffling after shoe #", t.shoesSofar, "; rounds so far=", t.rounds, "; player balance=", p.Balance())
		}
	}
}

func (t *duelingTable) playOneRound(p Player) {
	if glog.V(50) {
		glog.Infoln("New round #", t.rounds)
	}
	t.rounds++

	z := 0
	f := false
	ph := Hand([]Card{})
	t.phs = []HandWithBet{
		HandWithBet{
			Hand:           &ph,
			OriginalBet:    t.bet,
			NumberOfSplits: &z,
			AceSplit:       &f,
		},
	}

	dh := Hand([]Card{})
	t.dh = &dh

	t.dealTo(t.phs[0].Hand)
	t.dealTo(t.dh)
	t.dealTo(t.phs[0].Hand)

	// check for player blackjack
	t.phs[0].Blackjack = t.phs[0].Value() == 11 && t.phs[0].HasAce()

	dcv := (*t.dh)[0].ValueTenBased()

	// ask for surrenders
	var askForSurrenders bool
	switch t.r.SurrenderRule {
	case SurrenderAny:
		askForSurrenders = true
	case SurrenderingNotAllowed:
		askForSurrenders = false
	case SurrenderExceptDelaerAce:
		askForSurrenders = dcv != Ace
	default:
		panic("Invalid surrender rule")
	}
	if askForSurrenders {
		if p.Surrender(*t.phs[0].Hand, dcv) {
			if glog.V(50) {
				glog.Infoln("Player surrendered")
			}
			t.logHands()
			p.Balance().Pay(t.phs[0].OriginalBet*t.r.BetUnit/2, t.bal.Receiver())
			if glog.V(50) {
				glog.Infoln("Player balance", p.Balance().Amount())
			}
			return
		}
	}

	// players play first
	for handIndex := 0; handIndex < len(t.phs); {

		currentHand := t.phs[handIndex]
		v := currentHand.Hand.Value()

		if len(*currentHand.Hand) == 1 {
			// from a split
			t.dealTo(currentHand.Hand)
		}
		t.logHands()
		if v > 20 {
			handIndex++
			continue
		}
		if v == 11 && currentHand.HasAce() {
			handIndex++
			continue
		}

		switch logPlayerDecision(p.Decide(*currentHand.Hand, dcv)) {
		case Stand:
			handIndex++
		case Hit:
			if *currentHand.AceSplit && !t.r.CanHitAfterAceSplit {
				if glog.V(50) {
					glog.Infoln("Cannot hit...")
				}
				handIndex++
				continue
			}
			t.dealTo(currentHand.Hand)
		case DoubleDownElseHit:
			if *currentHand.AceSplit && !t.r.CanHitAfterAceSplit {
				if glog.V(50) {
					glog.Infoln("Cannot hit...")
				}
				handIndex++
				continue
			}

			double := true
			if len(*currentHand.Hand) != 2 {
				if glog.V(50) {
					glog.Infoln("Cannot double with ", len(*currentHand.Hand), " cards in hand; hitting instead...")
				}
				double = false
			}
			if !t.r.DoubleAfterSplits && *currentHand.NumberOfSplits != 0 {
				if glog.V(50) {
					glog.Infoln("Cannot double after splitting; hitting instead...")
				}
				double = false
			}
			t.dealTo(currentHand.Hand)
			if double {
				t.phs[handIndex].AdditionalBet += t.bet
				handIndex++
			}
		case DoubleDownElseStand:
			if *currentHand.AceSplit && !t.r.CanHitAfterAceSplit {
				if glog.V(50) {
					glog.Infoln("Cannot hit...")
				}
				handIndex++
				continue
			}

			if len(*currentHand.Hand) != 2 {
				if glog.V(50) {
					glog.Infoln("Cannot double with ", len(*currentHand.Hand), " cards in hand; standing instead...")
				}
				handIndex++
				continue
			}
			if !t.r.DoubleAfterSplits && *currentHand.NumberOfSplits != 0 {
				if glog.V(50) {
					glog.Infoln("Cannot double after splitting; standing instead...")
				}
				handIndex++
				continue
			}
			t.dealTo(currentHand.Hand)
			t.phs[handIndex].AdditionalBet += t.bet
			handIndex++
		case SplitElseHit:
			split := true
			if *currentHand.AceSplit && *currentHand.NumberOfSplits == t.r.MaxAceSplits {
				if glog.V(50) {
					glog.Infoln("Cannot split - reached max number of splits for this hand; hitting instead")
				}
				split = false
			} else if *currentHand.NumberOfSplits == t.r.MaxSplitsPerHand {
				if glog.V(50) {
					glog.Infoln("Cannot split - reached max number of splits for this hand; hitting instead")
				}
				split = false
			}

			if split {
				nh, err := currentHand.Split()
				if err != nil {
					if glog.V(50) {
						glog.Infoln("Split denied - ", err, "; hitting instead")
					}
					if *currentHand.AceSplit && !t.r.CanHitAfterAceSplit {
						if glog.V(50) {
							glog.Infoln("Cannot hit...")
						}
						handIndex++
						continue
					}
					t.dealTo(currentHand.Hand)
					continue
				}
				nhwb := HandWithBet{
					Hand:           &nh,
					AdditionalBet:  t.bet,
					NumberOfSplits: currentHand.NumberOfSplits,
					AceSplit:       currentHand.AceSplit,
				}
				*nhwb.NumberOfSplits = *nhwb.NumberOfSplits + 1
				*nhwb.AceSplit = nh.Value() == 1
				t.phs = append(t.phs, nhwb)
				t.logHands()
			} else {
				if *currentHand.AceSplit && !t.r.CanHitAfterAceSplit {
					if glog.V(50) {
						glog.Infoln("Cannot hit...")
					}
					handIndex++
					continue
				}
				t.dealTo(currentHand.Hand)
			}
		case SplitElseStand:
			if *currentHand.AceSplit && *currentHand.NumberOfSplits == t.r.MaxAceSplits {
				if glog.V(50) {
					glog.Infoln("Cannot split - reached max number of splits for this hand; standing instead")
				}
				handIndex++
				continue
			}
			if *currentHand.NumberOfSplits == t.r.MaxSplitsPerHand {
				if glog.V(50) {
					glog.Infoln("Cannot split - reached max number of splits for this hand; standing instead")
				}
				handIndex++
				continue
			}
			nh, err := currentHand.Split()
			if err != nil {
				if glog.V(50) {
					glog.Infoln("Split denied - ", err, "; standing instead")
				}
				handIndex++
				continue
			}
			nhwb := HandWithBet{
				Hand:           &nh,
				AdditionalBet:  t.bet,
				NumberOfSplits: currentHand.NumberOfSplits,
				AceSplit:       currentHand.AceSplit,
			}
			*nhwb.NumberOfSplits = *nhwb.NumberOfSplits + 1
			*nhwb.AceSplit = nh.Value() == 1
			t.phs = append(t.phs, nhwb)
			t.logHands()
		default:
			panic("Invalid player decision")
		}
	}

	if glog.V(50) {
		glog.Infoln("Players have finished their turns")
	}
	t.logHands()

	//dealers turn
	dealerDrawn := false
	dealerV := 0
	dealerBJ := false
	for _, hwb := range t.phs {
		if hwb.Blackjack {
			if !dealerDrawn && ((*t.dh)[0].Value() == Ace || (*t.dh)[0].Value() > Nine) {
				// draw one card to try for blackjack
				t.dealTo(t.dh)
				dealerDrawn = true
				if t.dh.Value() == 11 {
					dealerBJ = true
				}
			}
			if dealerBJ {
				// push
				continue
			}

			t.payPlayer(p, uint(float64(hwb.OriginalBet*t.r.BetUnit)*t.r.BlackjackPayout))
			continue
		}
		pv := hwb.Hand.Value()
		if pv > 21 {
			// player hand busted

			if t.r.PlayerLosesOriginalBetOnlyOnDealerBJ {
				// if player has split, dealer needs to try for blackjack if possible
				if *hwb.NumberOfSplits != 0 {

					if !dealerDrawn && ((*t.dh)[0].Value() == Ace || (*t.dh)[0].Value() > Nine) {
						// draw one card to try for blackjack
						t.dealTo(t.dh)
						dealerDrawn = true
						if t.dh.Value() == 11 {
							dealerBJ = true
						}
					}
					if dealerBJ {
						if hwb.OriginalBet != 0 {
							p.Balance().Pay(hwb.OriginalBet*t.r.BetUnit, t.bal.Receiver())
							if glog.V(50) {
								glog.Infoln("Player balance", p.Balance().Amount())
							}
						}
						continue
					}
				}
			}
			p.Balance().Pay((hwb.OriginalBet+hwb.AdditionalBet)*t.r.BetUnit, t.bal.Receiver())
			if glog.V(50) {
				glog.Infoln("Player balance", p.Balance().Amount())
			}
			continue
		}

		// player not busted
		// try for BJ if possible
		if !dealerDrawn && ((*t.dh)[0].Value() == Ace || (*t.dh)[0].Value() > Nine) {
			// draw one card to try for blackjack
			t.dealTo(t.dh)
			dealerDrawn = true
			if t.dh.Value() == 11 {
				dealerBJ = true
			}
		}

		if dealerBJ {
			if t.r.PlayerLosesOriginalBetOnlyOnDealerBJ {
				p.Balance().Pay(hwb.OriginalBet*t.r.BetUnit, t.bal.Receiver())
			} else {
				if hwb.OriginalBet != 0 && hwb.AdditionalBet != 0 {
					p.Balance().Pay((hwb.OriginalBet+hwb.AdditionalBet)*t.r.BetUnit, t.bal.Receiver())
				}
			}
			if glog.V(50) {
				glog.Infoln("Player balance", p.Balance().Amount())
			}
			continue
		}

		// dealer no BJ - to draw to 17
		if dealerV == 0 {
			for {
				v := t.dh.Value()
				if t.dh.HasAce() {
					if v < 7 { // soft and less than 17
						t.dealTo(t.dh)
						dealerDrawn = true
						continue
					}
					if v == 7 && !t.r.DealerHitsOnSoft17 { // soft 17 and to stand
						dealerV = 17
						break
					}

					if v < 12 { // soft and more than 17 so stand
						dealerV = v + 10
						break
					}
					// this is a hard hand
				}
				if v < 17 {
					t.dealTo(t.dh)
					dealerDrawn = true
					continue
				}

				// stand
				dealerV = v
				break
			}
		}

		if pv < 12 && hwb.Hand.HasAce() {
			pv += 10
		}

		switch {
		case dealerV > 21:
			// dealer busted
			t.payPlayer(p, (hwb.OriginalBet+hwb.AdditionalBet)*t.r.BetUnit)
		case dealerV > pv:
			p.Balance().Pay((hwb.OriginalBet+hwb.AdditionalBet)*t.r.BetUnit, t.bal.Receiver())
			if glog.V(50) {
				glog.Infoln("Player balance", p.Balance().Amount())
			}
		case dealerV < pv:
			t.payPlayer(p, (hwb.OriginalBet+hwb.AdditionalBet)*t.r.BetUnit)
		}
	}
}

func (t *duelingTable) payPlayer(p Player, amount uint) {
	t.bal.Pay(amount, p.Balance().Receiver())
	if glog.V(50) {
		glog.Infoln("Player balance", p.Balance().Amount())
	}
}

func (t *duelingTable) logHands() {
	if glog.V(50) {
		glog.Infoln("Player Hands ", t.phs, "; dealer hand ", t.dh)
	}
}

func (t *duelingTable) dealTo(h *Hand) {
	c := t.s[0]
	for _, l := range t.cardListeners.ls {
		(l.(CardListener))(c)
	}
	*t.burnt = append(*t.burnt, c)
	*h = append(*h, c)
	t.s = t.s[1:]
	t.logHands()
}

func NewDuelingTable(r *Rules, shuffler Shuffler) DuelingTable {
	burnt := Hand([]Card{})
	t := &duelingTable{
		r:        r,
		p:        make(chan Player),
		s:        PerfectShuffler(NewShoe(r.NumberOfDecks), []Card{}),
		bal:      NewBalance(),
		shuffler: shuffler,
		burnt:    &burnt,
	}
	return t
}

func logPlayerDecision(d PlayDecision) PlayDecision {
	if glog.V(50) {
		glog.Infoln("Player has decided on ", d)
	}
	return d
}
