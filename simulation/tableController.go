package simulation

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/xushichangdesmond/blackjack"
)

type TableController interface {
	StartNewShoe() error

	EndGame()
	StartNewGame()

	NextHand() (boxPosition int, h blackjack.Hand)

	DealOneCard()
}

type tableController table

func (c *tableController) StartNewShoe() error {
	//c.Lock()
	//defer c.Unlock()

	if c.gameInProgress {
		return errors.New("Game is in progress")
	}

	c.shoe.Shuffle()

	return nil
}

func (c *tableController) String() string {
	//c.Lock()
	//defer c.Unlock()

	buffer := bytes.Buffer{}
	buffer.WriteString("{\n")
	buffer.WriteString(fmt.Sprintf("Game in progress - %v\n", c.gameInProgress))
	buffer.WriteString(fmt.Sprintf("Decks in shoe - %v\n", c.numberOfDecksUsed))
	for i, b := range c.boxes {
		buffer.WriteString(fmt.Sprintf("Box#%v - %v\n", i, b))
	}
	buffer.WriteString("}\n")
	return buffer.String()
}
