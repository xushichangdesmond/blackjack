package blackjack

import "github.com/golang/glog"

type RunningCounter interface {
	RunningCount() int
	Reset()
}

type hiLoRunningCounter int

func (r *hiLoRunningCounter) RunningCount() int {
	return int(*r)
}

func (r *hiLoRunningCounter) Reset() {
	*r = 0
}

func NewHiLoRunningCounter() (RunningCounter, CardListener) {
	counter := hiLoRunningCounter(0)
	return &counter, func(c Card) {
		if c.Value() == Ace || c.Value() > Nine {
			counter--
		} else if c.Value() < 7 {
			counter++
		}
	}
}

type csmCounter [3]int

func (r *csmCounter) RunningCount() int {
	return (*r)[0] + (*r)[1] + (*r)[2]
}

func (r *csmCounter) Reset() {
	(*r)[0] = 0
	(*r)[1] = 0
	(*r)[2] = 0
}

func NewCSMCounter() (RunningCounter, CardListener, ShuffleListener) {
	counter := csmCounter([3]int{0, 0})
	return &counter,
		func(c Card) {
			if c.Value() == Ace || c.Value() > Nine {
				counter[0]--
			} else if c.Value() < 7 {
				counter[0]++
			}
			if glog.V(50) {
				glog.Infoln("Count", counter)
			}
		},
		func() {
			counter[0], counter[1], counter[2] = 0, counter[0], counter[1]
		}
}

type red7RunningCounter int

func (r *red7RunningCounter) RunningCount() int {
	return int(*r)
}

func (r *red7RunningCounter) Reset() {
	*r = 0
}

func NewRed7RunningCounter() (RunningCounter, CardListener, ShuffleListener) {
	counter := red7RunningCounter(0)
	return &counter,
		func(c Card) {
			if c.Value() == Ace || c.Value() > Nine {
				counter--
			} else if c.Value() < 6 {
				counter++
			} else if c.Value() == 7 && (c.Suit() == Diamonds || c.Suit() == Hearts) {
				counter++
			}
		},
		func() {
			counter.Reset()
		}
}

type blackHoleCounter struct{}

func (r *blackHoleCounter) RunningCount() int {
	return 0
}

func (r *blackHoleCounter) Reset() {
}

func NewBlackHoleCounter() RunningCounter {
	return &blackHoleCounter{}
}
