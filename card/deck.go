package card

import "errors"

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
