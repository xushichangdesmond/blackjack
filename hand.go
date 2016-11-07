package blackjack

import (
	"errors"
	"fmt"
)

type Hand []Card

var ErrCannotSplit = errors.New("Cannot split")

func (h Hand) HasAce() bool {
	for _, c := range h {
		if c.Value() == Ace {
			return true
		}
	}
	return false
}

func (h Hand) Value() int {
	v := 0
	for _, c := range h {
		v += int(c.ValueTenBased())
	}
	return v
}

func (h Hand) IsSoft() bool {
	return h.HasAce() && h.Value() < 12
}

func (h *Hand) Split() (newHand Hand, err error) {
	if len(*h) != 2 {
		return nil, ErrCannotSplit
	}
	if (*h)[0].ValueTenBased() != (*h)[1].ValueTenBased() {
		return nil, ErrCannotSplit
	}
	newHand = Hand([]Card{(*h)[1]})
	*h = (*h)[:1]
	return newHand, nil
}

type HandWithBet struct {
	*Hand
	OriginalBet    uint
	AdditionalBet  uint
	Blackjack      bool
	NumberOfSplits *int
	AceSplit       *bool
}

func (h HandWithBet) String() string {
	return fmt.Sprint(h.Hand)
}
