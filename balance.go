package blackjack

import "sync"

type Receiver func(amount uint)

type Balance interface {
	Amount() int
	Receiver() Receiver
	Pay(amount uint, r Receiver)
}

type balance struct {
	sync.Mutex
	amount int
}

func (b *balance) Amount() int {
	return b.amount
}

func (b *balance) Receiver() Receiver {
	return func(amount uint) {
		b.Lock()
		defer b.Unlock()

		b.amount += int(amount)
	}
}

func (b *balance) Pay(amount uint, r Receiver) {
	b.Lock()
	defer b.Unlock()

	b.amount -= int(amount)
	r(amount)
}
func NewBalance() Balance {
	return &balance{}
}
