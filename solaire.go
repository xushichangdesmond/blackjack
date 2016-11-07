package blackjack

func NewSolaireRules() *Rules {
	return &Rules{
		MaxSplitsPerHand: 2,
		MaxAceSplits:     1,

		DoublingRule:      DoubleOnAny,
		DoubleAfterSplits: true,

		DealerHitsOnSoft17: false,
		BlackjackPayout:    1.5,

		SurrenderRule:  SurrenderExceptDelaerAce,
		EarlySurrender: true,

		MinBetUnits: 1,
		MaxBetUnits: 100,
		BetUnit:     50,

		NumberOfDecks: 6,

		ShufflingPoint: 6 * 52, // perfect CSM

		PlayerLosesOriginalBetOnlyOnDealerBJ: true,
		CanHitAfterAceSplit:                  false,
	}
}
