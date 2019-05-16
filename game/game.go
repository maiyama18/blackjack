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

	g.Init()

	playerScore, burst := g.PlayersTurn()
	if burst {
		g.logger.Burst(g.player.name, playerScore)
		g.logger.Logf("%s LOSE\n", g.player.name)
		return
	}

	dealerScore, burst := g.DealersTurn()
	if burst {
		g.logger.Burst("dealer", dealerScore)
		g.logger.Logf("%s WIN\n", g.player.name)
		return
	}

	g.logger.Scores(g.player.name, g.dealer.name, playerScore, dealerScore)

	if playerScore > dealerScore {
		g.logger.Logf("%s WIN\n", g.player.name)
	} else if playerScore < dealerScore {
		g.logger.Logf("%s LOSE\n", g.player.name)
	} else {
		g.logger.Logf("EVEN\n")
	}
}

type Game struct {
	deck   *card.Deck
	player *Player
	dealer *Player

	logger *log.Logger
}

func New(out io.Writer) *Game {
	deck := card.NewDeck()

	player := NewPlayer("you")
	dealer := NewPlayer("dealer")

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

		g.logger.Draw(g.player.name, cp, false)
		g.player.hand.Add(cp)

		cd, _ := g.deck.Draw()

		secret := false
		if i == 1 {
			secret = true
		}
		g.logger.Draw(g.dealer.name, cd, secret)
		g.dealer.hand.Add(cd)
	}
}

// PlayersTurn make the player draw cards.
// This method ends when the player declare to stop to draw cards or the player's hand bursts.
func (g *Game) PlayersTurn() (int, bool) {
	g.logger.Hand(g.player.name, g.player.hand)
	for g.doesPlayerDraw() {
		c, _ := g.deck.Draw()
		g.logger.Draw(g.player.name, c, false)

		g.player.hand.Add(c)
		g.logger.Hand(g.player.name, g.player.hand)

		if g.player.hand.Burst() {
			return g.player.hand.BestScore(), true
		}
	}
	return g.player.hand.BestScore(), false
}

// DealersTurn make the dealer draw cards.
func (g *Game) DealersTurn() (int, bool) {
	for g.doesDealerDraw() {
		c, _ := g.deck.Draw()
		g.dealer.hand.Add(c)

		g.logger.Draw(g.dealer.name, c, false)
		g.logger.Hand(g.dealer.name, g.dealer.hand)

		if g.dealer.hand.Burst() {
			return g.dealer.hand.BestScore(), true
		}
	}
	return g.dealer.hand.BestScore(), false
}

func (g *Game) doesPlayerDraw() bool {
	g.logger.Logf("Draw more? (y/n): ")
	for {
		var answer string
		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			g.logger.Logf("Draw more? (y/n): ")
			continue
		}

		switch answer {
		case "y":
			return true
		case "n":
			return false
		default:
			g.logger.Logf("Draw more? (y/n): ")
			continue
		}
	}
}
func (g *Game) doesDealerDraw() bool {
	return g.dealer.hand.BestScore() < 17
}
