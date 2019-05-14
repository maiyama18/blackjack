package card

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
