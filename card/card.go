package card

import (
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	suit string
	rank string
}

func (c Card) String() string {
	switch c.suit {
	case "Spades":
		return "♦" + c.rank
	case "Diamonds":
		return "♦" + c.rank
	case "Clubs":
		return "♣" + c.rank
	case "Hearts":
		return "♥" + c.rank
	}
	return ""
}

func (c Card) Score() int {
	switch c.rank {
	case "A":
		return 1
	case "J", "Q", "K":
		return 10
	default:
		p, _ := strconv.Atoi(c.rank)
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
