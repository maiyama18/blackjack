package card

import (
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

func (c Card) Score() int {
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

func shuffle(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := len(cards) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}
