package blackjack

import "testing"
import "github.com/stretchr/testify/assert"

func TestSplit(t *testing.T) {
	h := Hand([]Card{newCardUnchecked(Six)})
	_, err := h.Split()
	assert.Error(t, err)

	h = Hand([]Card{newCardUnchecked(Six), newCardUnchecked(Seven)})
	_, err = h.Split()
	assert.Error(t, err)

	h = Hand([]Card{newCardUnchecked(Six), newCardUnchecked(Six)})
	nh, err := h.Split()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(h))
	assert.Equal(t, 1, len(nh))
	assert.Equal(t, newCardUnchecked(Six), h[0])
	assert.Equal(t, newCardUnchecked(Six), nh[0])

	h = Hand([]Card{newCardUnchecked(Ten), newCardUnchecked(King)})
	nh, err = h.Split()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(h))
	assert.Equal(t, 1, len(nh))
	assert.Equal(t, newCardUnchecked(Ten), h[0])
	assert.Equal(t, newCardUnchecked(King), nh[0])
}

func newCardUnchecked(cv CardValue) Card {
	c, _ := NewCard(Diamonds, cv)
	return c
}
