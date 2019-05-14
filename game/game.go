package game

import (
	"fmt"

	"github.com/mui87/blackjack/card"
)

type Player struct {
	hand *card.Hand
}

func NewPlayer() *Player {
	return &Player{
		hand: card.NewHand(),
	}
}

type Dealer struct {
	hand *card.Hand
}

func NewDealer() *Dealer {
	return &Dealer{
		hand: card.NewHand(),
	}
}

type Game struct {
	deck   *card.Deck
	player *Player
	dealer *Dealer
}

func New() *Game {
	fmt.Println("game start")

	deck := card.NewDeck()

	player := NewPlayer()
	dealer := NewDealer()

	for i := 0; i < 2; i++ {
		cp, _ := deck.Draw()
		cd, _ := deck.Draw()

		fmt.Printf("You draw: %s\n", cp)
		player.hand.Add(cp)

		if i == 0 {
			fmt.Printf("Dealer draw: %s\n", cd)
		} else {
			fmt.Print("Dealer draw: ?\n")
		}
		dealer.hand.Add(cd)
	}

	return &Game{
		deck:   deck,
		player: player,
		dealer: dealer,
	}
}
