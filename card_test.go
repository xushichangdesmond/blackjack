package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	for s := Suit(0); s < 10; s++ {
		for v := CardValue(0); v <= 20; v++ {
			tc, err := NewCard(s, v)

			if (int(s) < len(Suits)) && (int(v) < len(CardValues)) {
				// valid card
				require.NoError(t, err)
				c := tc.(Card)
				assert.Equal(t, s, c.Suit())
				assert.Equal(t, v, c.Value())

				v10 := c.ValueTenBased()
				if v < Jack {
					assert.Equal(t, v, v10)
				} else {
					assert.Equal(t, Ten, v10)
				}
			} else {
				assert.Error(t, err)
			}
		}
	}
}
