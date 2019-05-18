package card

import "testing"

func TestHand_String(t *testing.T) {
	tests := []struct {
		name  string
		cards []Card
		want  string
	}{
		{
			name:  "empty",
			cards: []Card{},
			want:  "{}",
		},
		{
			name:  "one_card",
			cards: []Card{{suit: "Spades", rank: "A"}},
			want:  "{♠A}",
		},
		{
			name:  "multiple_cards",
			cards: []Card{{suit: "Spades", rank: "A"}, {suit: "Diamonds", rank: "J"}},
			want:  "{♠A, ♦J}",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := NewHand()
			h.cards = test.cards
			got := h.String()
			if got != test.want {
				t.Fatalf("hand string wrong. want=%s, got=%s", test.want, got)
			}
		})
	}
}
