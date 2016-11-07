package blackjack

import "testing"
import "github.com/stretchr/testify/assert"

func TestHiLoRunningCounter(t *testing.T) {
	c, l := NewHiLoRunningCounter()
	assert.Equal(t, 0, c.RunningCount())
	l(newCardUnchecked(Ace))
	assert.Equal(t, -1, c.RunningCount())
	l(newCardUnchecked(King))
	assert.Equal(t, -2, c.RunningCount())
	l(newCardUnchecked(Queen))
	assert.Equal(t, -3, c.RunningCount())
	l(newCardUnchecked(Jack))
	assert.Equal(t, -4, c.RunningCount())
	l(newCardUnchecked(Ten))
	assert.Equal(t, -5, c.RunningCount())
	l(newCardUnchecked(Nine))
	assert.Equal(t, -5, c.RunningCount())
	l(newCardUnchecked(Eight))
	assert.Equal(t, -5, c.RunningCount())
	l(newCardUnchecked(Seven))
	assert.Equal(t, -5, c.RunningCount())
	l(newCardUnchecked(Six))
	assert.Equal(t, -4, c.RunningCount())
	l(newCardUnchecked(Five))
	assert.Equal(t, -3, c.RunningCount())

	c.Reset()
	assert.Equal(t, 0, c.RunningCount())

	l(newCardUnchecked(Four))
	assert.Equal(t, 1, c.RunningCount())
	l(newCardUnchecked(Three))
	assert.Equal(t, 2, c.RunningCount())
	l(newCardUnchecked(Two))
	assert.Equal(t, 3, c.RunningCount())
}

func TestCSMCounter(t *testing.T) {
	c, cl, sl := NewCSMCounter()
	assert.Equal(t, 0, c.RunningCount())
	cl(newCardUnchecked(Ace))
	assert.Equal(t, -1, c.RunningCount())
	cl(newCardUnchecked(King))
	assert.Equal(t, -2, c.RunningCount())
	sl()
	assert.Equal(t, -2, c.RunningCount())
	cl(newCardUnchecked(Five))
	assert.Equal(t, -1, c.RunningCount())
	cl(newCardUnchecked(Five))
	assert.Equal(t, 0, c.RunningCount())
	cl(newCardUnchecked(Six))
	assert.Equal(t, 1, c.RunningCount())
	cl(newCardUnchecked(Seven))
	assert.Equal(t, 1, c.RunningCount())
	sl()
	assert.Equal(t, 1, c.RunningCount())
	sl()
	assert.Equal(t, 3, c.RunningCount())

	c.Reset()
	assert.Equal(t, 0, c.RunningCount())
}
