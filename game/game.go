package game

import (
	"fmt"
	"io"
	"os"

	"github.com/mui87/blackjack/log"

	"github.com/mui87/blackjack/card"
)

func Run() {
	g := New(os.Stdout)

	playerScore, burst := g.PlayersTurn()
	if burst {
		fmt.Printf("You lose due to burst with score: %d\n", playerScore)
		return
	}

	dealerScore, burst := g.DealersTurn()
	if burst {
		fmt.Printf("You win due to burst of dealer with score: %d\n", g.dealer.hand.BestScore())
		return
	}

	if playerScore > dealerScore {
		fmt.Printf("You win. you: %d, dealer: %d\n", playerScore, dealerScore)
	} else if playerScore < dealerScore {
		fmt.Printf("You lose. you: %d, dealer: %d\n", playerScore, dealerScore)
	} else {
		fmt.Printf("Even. you: %d, dealer: %d\n", playerScore, dealerScore)
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
	fmt.Printf("Your score is %d. Draw more? (y/n): ", p.hand.BestScore())
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
	return d.hand.BestScore() < 17
}

type Game struct {
	deck   *card.Deck
	player *Player
	dealer *Dealer

	logger *log.Logger
}

func New(out io.Writer) *Game {
	deck := card.NewDeck()

	player := NewPlayer()
	dealer := NewDealer()

	logger := log.NewLogger(out)

	return &Game{
		deck:   deck,
		player: player,
		dealer: dealer,
		logger: logger,
	}
}

func (g *Game) Init() {
	for i := 0; i < 2; i++ {
		cp, _ := g.deck.Draw()

		g.logger.Draw("you", cp, false)
		g.player.hand.Add(cp)

		cd, _ := g.deck.Draw()

		secret := false
		if i == 1 {
			secret = true
		}
		g.logger.Draw("dealer", cd, secret)
		g.dealer.hand.Add(cd)
	}
}

// PlayersTurn make the player draw cards.
// This method ends when the player declare to stop to draw cards or the player's hand bursts.
func (g *Game) PlayersTurn() (int, bool) {
	for g.player.doesDraw() {
		c, _ := g.deck.Draw()
		g.player.hand.Add(c)
		g.logger.Draw("you", c, false)
		if g.player.hand.Burst() {
			return g.player.hand.BestScore(), true
		}
	}
	return g.player.hand.BestScore(), false
}

// DealersTurn make the dealer draw cards.
func (g *Game) DealersTurn() (int, bool) {
	c, _ := g.dealer.hand.LastCard()
	fmt.Printf("Dealer's second card was: %s\n", c)
	fmt.Printf("Dealer's score: %d\n", g.dealer.hand.BestScore())

	for g.dealer.doesDraw() {
		c, _ := g.deck.Draw()
		g.dealer.hand.Add(c)
		fmt.Printf("Dealer draw: %s\n", c)
		fmt.Printf("Dealer's score: %d\n", g.dealer.hand.BestScore())
		if g.dealer.hand.Burst() {
			return g.dealer.hand.BestScore(), true
		}
	}
	return g.dealer.hand.BestScore(), false
}
