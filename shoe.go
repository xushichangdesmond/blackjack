package blackjack

import "github.com/xushichangdesmond/blackjack/shuffle"

type Shoe Hand

var _ shuffle.Swapper = (Shoe)(nil)

func (s Shoe) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func NewShoe(numOfStandardDecks int) Shoe {
	s := make([]Card, numOfStandardDecks*52)

	n := 0

	for i := 0; i < numOfStandardDecks; i++ {
		for j := Suits[0]; int(j) < len(Suits); j++ {
			for k := Ace; int(k) < len(CardValues); k++ {
				c, err := NewCard(j, k)
				if err != nil {
					panic("unexpected")
				}
				s[n] = c
				n++
			}
		}
	}
	return s
}
