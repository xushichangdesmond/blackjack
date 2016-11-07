package blackjack

import "github.com/xushichangdesmond/blackjack/shuffle"

type Shoe Hand

var _ shuffle.Shufflable = (Shoe)(nil)

func (s Shoe) Shuffle() {
	shuffle.Shuffle(s)
}

func (s Shoe) Size() int {
	return len(s)
}

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

/*func (s Shoe) String() string {
	b := bytes.Buffer{}
	b.WriteRune('[')
	for _, c := range s {
		b.WriteString(c.String())
		b.WriteRune(',')
	}
	b.WriteRune(']')
	return b.String()
}*/
