//go:generate stringer -type=CardValue,Suit
package blackjack

import "fmt"
import "errors"

type CardValue int

// cardValues
const (
	Joker CardValue = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var CardValues []CardValue = []CardValue{
	Joker,
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

var cardValuesShortRepresentations = []string{
	"*", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

func (v CardValue) ShortRepresentation() string {
	return cardValuesShortRepresentations[v]
}

type Suit int

// suits
const (
	Diamonds Suit = iota
	Hearts
	Clubs
	Spades
)

var Suits []Suit = []Suit{
	Diamonds,
	Hearts,
	Clubs,
	Spades,
}

var suitsShortRepresentations = []rune{
	'♦', '♥', '♣', '♠',
}

func (s Suit) ShortRepresentation() rune {
	return suitsShortRepresentations[s]
}

type Card interface {
	// Value returns a 0 to 13 int8 representing the value of the card
	// where 0 is the Joker and 1 is the Ace
	Value() CardValue

	// ValueTenBased is same as Value, except that all face cards are ten-valued (returns 10)
	ValueTenBased() CardValue

	Suit() Suit
}

// card is a uint8 where first 2 bits are zeros, the next 2 bits represent the suit,
// and next 4 bits (0-13) represent the card value
type card uint8

var _ Card = (*card)(nil)
var _ fmt.Stringer = (*card)(nil)

var (
	ErrInvalidSuit      = errors.New("Invalid suit")
	ErrInvalidCardValue = errors.New("Invalid card value")
)

func NewCard(suit Suit, cv CardValue) (Card, error) {
	if (suit < 0) || (int(suit) >= len(Suits)) {
		return nil, ErrInvalidSuit
	}
	if (cv < 0) || (int(cv) >= len(CardValues)) {
		return nil, ErrInvalidCardValue
	}
	return card(int(suit)<<4 + int(cv)), nil
}

func (c card) Value() CardValue {
	return CardValue(c & 15)
}

func (c card) ValueTenBased() CardValue {
	if v := c.Value(); v < Ten {
		return v
	}
	return Ten
}

func (c card) Suit() Suit {
	return Suit(c >> 4)
}

func (c card) String() string {
	return c.Value().ShortRepresentation() + string(c.Suit().ShortRepresentation())
}
