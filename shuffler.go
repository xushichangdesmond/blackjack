package blackjack

import "github.com/xushichangdesmond/blackjack/shuffle"

type Shuffler func(s Shoe, burnt []Card) Shoe

var ShuffleMaster126 = func(s Shoe, burnt []Card) Shoe {
	cs := append(s, burnt...)
	shuffle.Shuffle(cs, 16, len(cs))
	return cs
}

var PerfectShuffler = func(s Shoe, burnt []Card) Shoe {
	cs := append(s, burnt...)
	shuffle.Shuffle(cs, 0, len(cs))
	return cs
}
