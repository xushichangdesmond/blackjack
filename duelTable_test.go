package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDealTo(t *testing.T) {
	dt := NewDuelingTable(NewSolaireRules(), PerfectShuffler).(*duelingTable)

	shoeSize := len(dt.s)
	firstCardInShoe := dt.s[0]
	secondCardInShoe := dt.s[1]
	thirdCardInShoe := dt.s[2]

	h := Hand([]Card{})
	dt.dealTo(&h)
	dt.dealTo(&h)

	assert.Equal(t, shoeSize-2, len(dt.s))
	assert.Equal(t, 2, len(h))

	assert.Equal(t, thirdCardInShoe, dt.s[0])

	assert.Equal(t, firstCardInShoe, h[0])
	assert.Equal(t, secondCardInShoe, h[1])
}

func TestRareCases(t *testing.T) {
	r := NewSolaireRules()
	r.MinBetUnits = 1
	r.BetUnit = 10

	dt := NewDuelingTable(r, PerfectShuffler).(*duelingTable)

	p := NewMockPlayer(r.MinBetUnits)
	p.Mock().On("Decide", mock.Anything, mock.Anything).Twice().Return(SplitElseHit)
	p.Mock().On("Decide", mock.Anything, mock.Anything).Once().Return(Stand)

	dt.s = Shoe([]Card{
		newCardUnchecked(Ace),
		newCardUnchecked(Ace),
		newCardUnchecked(Ace),
		newCardUnchecked(Ace),
		newCardUnchecked(Ten),
		newCardUnchecked(Seven),
	})
	dt.bet = r.MinBetUnits
	dt.playOneRound(p)
	p.Mock().AssertExpectations(t)

	assert.Equal(t, 0, p.Balance().Amount())

	p = NewMockPlayer(r.MinBetUnits)
	p.Mock().On("Decide", mock.Anything, mock.Anything).Twice().Return(SplitElseHit)
	p.Mock().On("Decide", mock.Anything, mock.Anything).Once().Return(Stand)

	dt.s = Shoe([]Card{
		newCardUnchecked(Ace),
		newCardUnchecked(Ace),
		newCardUnchecked(Ace),
		newCardUnchecked(Ace),
		newCardUnchecked(Ten),
		newCardUnchecked(Ten),
	})
	dt.bet = r.MinBetUnits
	dt.playOneRound(p)
	p.Mock().AssertExpectations(t)

	assert.Equal(t, -10, p.Balance().Amount())

	p = NewMockPlayer(r.MinBetUnits)
	p.Mock().On("Surrender", mock.Anything, mock.Anything).Once().Return(false)
	p.Mock().On("Decide", mock.Anything, mock.Anything).Once().Return(SplitElseHit)
	p.Mock().On("Decide", mock.Anything, mock.Anything).Once().Return(DoubleDownElseHit)
	p.Mock().On("Decide", mock.Anything, mock.Anything).Once().Return(Stand)

	dt.s = Shoe([]Card{
		newCardUnchecked(Eight),
		newCardUnchecked(Ten),
		newCardUnchecked(Eight),
		newCardUnchecked(Three),
		newCardUnchecked(Ten),
		newCardUnchecked(Ten),
		newCardUnchecked(Ace),
	})
	dt.bet = r.MinBetUnits
	dt.playOneRound(p)
	p.Mock().AssertExpectations(t)

	assert.Equal(t, -10, p.Balance().Amount())
}
