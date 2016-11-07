package blackjack

import (
	"fmt"
	"sync"
)

type Receiver func(amount uint)

type Balance interface {
	Amount() int
	Receiver() Receiver
	Pay(amount uint, r Receiver)
}

type balance struct {
	sync.Mutex
	amount int
	high   int
	low    int
}

func (b *balance) Amount() int {
	return b.amount
}

func (b *balance) Receiver() Receiver {
	return func(amount uint) {
		b.Lock()
		defer b.Unlock()

		b.amount += int(amount)
		if b.high < b.amount {
			b.high = b.amount
		}
	}
}

func (b *balance) Pay(amount uint, r Receiver) {
	b.Lock()
	defer b.Unlock()

	b.amount -= int(amount)
	r(amount)
	if b.low > b.amount {
		b.low = b.amount
	}
}

func (b *balance) String() string {
	return fmt.Sprintf("Amount %d, High %d, Low %d", b.amount, b.high, b.low)
}

func NewBalance() Balance {
	return &balance{}
}
