package shuffle

import (
	"math/rand"
)

type Shufflable interface {
	Size() int // number of elements
	Swap(indexA int, indexB int)
}

func Shuffle(s Shufflable) {
	for i := s.Size() - 1; i > -1; i-- {
		s.Swap(i, rand.Intn(i+1))
	}
}
