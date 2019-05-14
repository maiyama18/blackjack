package card

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	Suit string
	Rank string
}

func (c Card) String() string {
	switch c.Suit {
	case "Spades":
		return "♦" + c.Rank
	case "Diamonds":
		return "♦" + c.Rank
	case "Clubs":
		return "♣" + c.Rank
	case "Hearts":
		return "♥" + c.Rank
	}
	return ""
}

func (c Card) Point() int {
	switch c.Rank {
	case "A":
		return 1
	case "J", "Q", "K":
		return 10
	default:
		p, _ := strconv.Atoi(c.Rank)
		return p
	}
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	var cards []Card
	for _, s := range []string{"Spades", "Diamonds", "Clubs", "Hearts"} {
		for _, r := range []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}

	return &Deck{Cards: shuffle(cards)}
}

func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}

	c := d.Cards[0]
	d.Cards = d.Cards[1:]

	return c, nil
}

func shuffle(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := len(cards) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}
