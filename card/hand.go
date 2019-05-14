package card

type Hand struct {
	Cards []Card
}

func NewHand() *Hand {
	return &Hand{Cards: []Card{}}
}

func (h *Hand) Add(card Card) {
	h.Cards = append(h.Cards, card)
}

func (h *Hand) Score() int {
	score := 0
	for _, c := range h.Cards {
		score += c.Point()
	}
	return score
}
