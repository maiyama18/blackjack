package card

import (
	"fmt"
	"strings"
)

type Hand struct {
	cards []Card
}

func NewHand() *Hand {
	return &Hand{cards: []Card{}}
}

func (h *Hand) String() string {
	var cardStrs []string
	for _, c := range h.cards {
		cardStrs = append(cardStrs, c.String())
	}

	return fmt.Sprintf("{%s}", strings.Join(cardStrs, ", "))
}

func (h *Hand) Add(card Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) Score() Score {
	var score Score
	for _, c := range h.cards {
		score = score.Plus(c.Score())
	}
	return score
}

func (h *Hand) BestScore() int {
	if len(h.Score()) == 0 {
		return 0
	}

	var valids, bursts Score
	for _, s := range h.Score() {
		if s <= 21 {
			valids = append(valids, s)
		} else {
			bursts = append(bursts, s)
		}
	}

	if len(valids) == 0 {
		return bursts.Min()
	}
	return valids.Max()
}

func (h *Hand) Burst() bool {
	return h.BestScore() > 21
}
