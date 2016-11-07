package shuffle

import (
	"math/rand"
)

type Swapper interface {
	Swap(indexA int, indexB int)
}

func Shuffle(s Swapper, start int, end int) {
	n := end - start
	for i := end - 1; i >= start; i-- {
		s.Swap(i, start+rand.Intn(n))
	}
}
