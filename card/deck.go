package card

import "errors"

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	var cards []Card
	for _, s := range []string{"Spades", "Diamonds", "Clubs", "Hearts"} {
		for _, r := range []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} {
			cards = append(cards, Card{suit: s, rank: r})
		}
	}

	return &Deck{cards: shuffle(cards)}
}

func (d *Deck) Draw() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}

	c := d.cards[0]
	d.cards = d.cards[1:]

	return c, nil
}
