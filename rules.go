package blackjack

type DoublingRule int

const (
	DoubleOnAny = iota
	DoubleOnTenEleven
	DoubleOnNineTenEleven
	DoublingNotAllowed
)

type SurrenderRule int

const (
	SurrenderAny SurrenderRule = iota
	SurrenderExceptDelaerAce
	SurrenderingNotAllowed
)

type Rules struct {
	MaxSplitsPerHand    int
	MaxAceSplits        int
	CanHitAfterAceSplit bool

	DoublingRule      DoublingRule
	DoubleAfterSplits bool

	DealerHitsOnSoft17 bool
	BlackjackPayout    float64

	SurrenderRule  SurrenderRule
	EarlySurrender bool
	//dealerDrawsHoleCard bool
	//dealerPeeksBlackjackWithTenOrFace bool
	//dealerPeeksBlackjackWithAce bool

	PlayerLosesOriginalBetOnlyOnDealerBJ bool

	MinBetUnits uint
	MaxBetUnits uint
	BetUnit     uint

	NumberOfDecks int
	// ShufflingPoint is the minimum number of cards in the shoe before each round
	ShufflingPoint int
}
