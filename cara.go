package blackjack

func NewCaraRules() *Rules {
	return &Rules{
		MaxSplitsPerHand: 2,
		MaxAceSplits:     1,

		DoublingRule:      DoubleOnAny,
		DoubleAfterSplits: true,

		DealerHitsOnSoft17: false,
		BlackjackPayout:    1.5,

		SurrenderRule:  SurrenderExceptDelaerAce,
		EarlySurrender: true,

		MinBetUnits: 3,
		MaxBetUnits: 100,
		BetUnit:     10,

		NumberOfDecks: 4,

		ShufflingPoint: 52,

		PlayerLosesOriginalBetOnlyOnDealerBJ: true,
		CanHitAfterAceSplit:                  false,
	}
}
