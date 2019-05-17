package card

import (
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.cards) != 13*4 {
		t.Fatalf("number of cards wrong. want=%d, got=%d", 13*4, len(deck.cards))
	}
}

func TestDeck_Draw(t *testing.T) {
	deck := NewDeck()
	for i := 0; i < 13*4; i++ {
		_, err := deck.Draw()
		if err != nil {
			t.Errorf("got error in draw: %s", err)
		}
	}
	_, err := deck.Draw()
	if !strings.Contains(err.Error(), "empty") {
		t.Errorf("expect %q to contain %q", err, "empty")
	}
}
