package card

import "errors"

type Hand struct {
	cards []Card
}

func NewHand() *Hand {
	return &Hand{cards: []Card{}}
}

func (h *Hand) Add(card Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) Score() int {
	score := 0
	for _, c := range h.cards {
		score += c.Score()
	}
	return score
}

func (h *Hand) Burst() bool {
	return h.Score() > 21
}

func (h *Hand) LastCard() (Card, error) {
	if len(h.cards) == 0 {
		return Card{}, errors.New("card never drawn")
	}
	return h.cards[len(h.cards)-1], nil
}
