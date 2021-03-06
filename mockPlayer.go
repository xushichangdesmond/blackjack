package blackjack

import "github.com/stretchr/testify/mock"

type MockPlayer interface {
	Player
	Mock() *mock.Mock
}

type mockPlayer struct {
	bal  Balance
	bet  uint
	mock *mock.Mock
}

func (m *mockPlayer) Balance() Balance {
	return m.bal
}

func (m *mockPlayer) Decide(ph Hand, dcv CardValue) PlayDecision {
	args := m.mock.Called(ph, dcv)
	return args.Get(0).(PlayDecision)
}

func (m *mockPlayer) Surrender(ph Hand, dcv CardValue) bool {
	args := m.mock.Called(ph, dcv)
	return args.Get(0).(bool)
}

func (m *mockPlayer) PlaceBet() uint {
	return m.bet
}

func (m *mockPlayer) Mock() *mock.Mock {
	return m.mock
}

func NewMockPlayer(bet uint) MockPlayer {
	return &mockPlayer{NewBalance(), bet, &mock.Mock{}}
}
