//go:generate stringer -type=PlayDecision
package blackjack

type PlayDecision int8

const (
	Stand PlayDecision = iota
	Hit
	DoubleDownElseHit
	DoubleDownElseStand
	SplitElseStand
	SplitElseHit
)

var PlayDecisions = []PlayDecision{
	Stand,
	Hit,
	DoubleDownElseHit,
	DoubleDownElseStand,
	SplitElseStand,
	SplitElseHit,
}
