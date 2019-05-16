package game

import (
	"github.com/mui87/blackjack/card"
)

type Player struct {
	name string
	hand *card.Hand
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
		hand: card.NewHand(),
	}
}
