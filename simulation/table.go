package simulation

import (
	"bytes"
	"context"
	"errors"

	"fmt"

	"time"

	"github.com/xushichangdesmond/blackjack"
)

type Table interface {
	GameInProgress() bool
	NumberOfBoxes() int
	WaitToTakeBox(ctx context.Context, p Player, position int) (PlaySession, error)
	TakeBox(p Player, position int) (PlaySession, error)
	NumberOfDesksUsed() int
}

type table struct {
	//sync.Mutex
	boxInProgress
	boxes             []box `final:"true"`
	shoe              blackjack.Shoe
	numberOfDecksUsed int `final:"true"`
}

func newTable(numBoxes int, decks int) *table {
	t := new(table)
	t.boxes = make([]box, numBoxes)
	for i := range t.boxes {
		t.boxes[i] = box{}
	}
	t.numberOfDecksUsed = decks
	t.shoe = blackjack.NewShoe(decks)
	return t
}

func NewTable(numBoxes int, decks int) (Table, TableController) {
	t := newTable(numBoxes, decks)
	return t, (*tableController)(t)
}

func (t *table) GameInProgress() bool {
	return t.gameInProgress
}

func (t *table) NumberOfBoxes() int {
	return len(t.boxes)
}

func (t *table) NumberOfDesksUsed() int {
	return t.numberOfDecksUsed
}

var SeatTakenError = errors.New("Seat taken")

func (t *table) TakeBox(p Player, position int) (PlaySession, error) {

	if position >= len(t.boxes) {
		return nil, fmt.Errorf("Invalid(out-of-bounds) position %v", position)
	}

	b := &t.boxes[position]
	//b.Lock()
	//defer b.Unlock()

	if b.player != nil {
		return nil, SeatTakenError
	}

	b.player = p

	return playSession{
		player: p,
		table:  *t,
	}, nil
}

func (t *table) WaitToTakeBox(ctx context.Context, p Player, position int) (PlaySession, error) {
	result := make(chan PlaySession, 1)
	errChan := make(chan error, 1)

	quit := make(chan struct{})
	go func() {
		for {
			ps, err := t.TakeBox(p, position)
			if err != nil {
				if err != SeatTakenError {
					errChan <- err
					return
				}
			} else {
				result <- ps
				return
			}
			select {
			case <-quit:
				return
			case <-time.After(time.Second):

			}

		}
	}()

	select {
	case ps := <-result:
		return ps, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		close(quit)
		return nil, ctx.Err()
	}
}

func (t *table) String() string {
	//t.Lock()
	//defer t.Unlock()

	buffer := bytes.Buffer{}
	buffer.WriteString("{\n")
	buffer.WriteString(fmt.Sprintf("Game in progress - %v\n", t.gameInProgress))
	buffer.WriteString(fmt.Sprintf("Decks in shoe - %v\n", t.numberOfDecksUsed))
	for i, b := range t.boxes {
		buffer.WriteString(fmt.Sprintf("Box#%v - %v\n", i, b))
	}
	buffer.WriteString("}\n")
	return buffer.String()
}
