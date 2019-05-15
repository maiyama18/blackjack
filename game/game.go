package game

import (
	"fmt"

	"github.com/mui87/blackjack/card"
)

func Run() {
	g := New()

	g.PlayersTurn()
	if g.player.hand.Burst() {
		fmt.Printf("You lose with score: %d\n", g.player.hand.Score())
		return
	}

	g.DealersTurn()
	if g.dealer.hand.Burst() {
		fmt.Printf("You win. you: %d, dealer: %d\n", g.player.hand.Score(), g.dealer.hand.Score())
		return
	}

	ps := g.player.hand.Score()
	ds := g.dealer.hand.Score()
	if ps > ds {
		fmt.Printf("You win. you: %d, dealer: %d\n", ps, ds)
	} else if ps < ds {
		fmt.Printf("You lose. you: %d, dealer: %d\n", ps, ds)
	} else {
		fmt.Printf("Even. you: %d, dealer: %d\n", ps, ds)
	}
}

type Player struct {
	hand *card.Hand
}

func NewPlayer() *Player {
	return &Player{
		hand: card.NewHand(),
	}
}

func (p *Player) doesDraw() bool {
	fmt.Printf("Your score is %d. Draw more? (y/n): ", p.hand.Score())
	for {
		var answer string
		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			fmt.Print("Answer in y/n: ")
			continue
		}

		switch answer {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Print("Answer in y/n: ")
			continue
		}
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

func (d *Dealer) doesDraw() bool {
	return d.hand.Score() < 17
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

// PlayersTurn make the player draw cards.
// This method ends when the player declare to stop to draw cards or the player's hand bursts.
func (g *Game) PlayersTurn() {
	for g.player.doesDraw() {
		c, _ := g.deck.Draw()
		g.player.hand.Add(c)
		fmt.Printf("You draw: %s\n", c)
		if g.player.hand.Burst() {
			return
		}
	}
}

// DealersTurn make the dealer draw cards.
func (g *Game) DealersTurn() {
	c, _ := g.dealer.hand.LastCard()
	fmt.Printf("Dealer's second card was: %s\n", c)
	fmt.Printf("Dealer's score: %d\n", g.dealer.hand.Score())

	for g.dealer.doesDraw() {
		c, _ := g.deck.Draw()
		g.dealer.hand.Add(c)
		fmt.Printf("Dealer draw: %s\n", c)
		fmt.Printf("Dealer's score: %d\n", g.dealer.hand.Score())
		if g.dealer.hand.Burst() {
			return
		}
	}
}
