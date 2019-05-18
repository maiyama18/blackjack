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
			h := &Hand{cards: test.cards}
			got := h.String()
			if got != test.want {
				t.Fatalf("hand string wrong. want=%s, got=%s", test.want, got)
			}
		})
	}
}

func TestHand_Add(t *testing.T) {
	h := &Hand{cards: []Card{{suit: "Spades", rank: "A"}}}

	added := Card{suit: "Diamond", rank: "J"}
	h.Add(added)

	if len(h.cards) != 2 {
		t.Fatalf("card not added")
	}
	if h.cards[1] != added {
		t.Fatalf("added card wrong. want=%v, got=%v", added, h.cards[1])
	}
}

func TestHand_Score(t *testing.T) {
	tests := []struct {
		name  string
		cards []Card
		want  Score
	}{
		{
			name:  "empty",
			cards: []Card{},
			want:  nil,
		},
		{
			name:  "one_card",
			cards: []Card{{suit: "Spades", rank: "3"}},
			want:  Score([]int{3}),
		},
		{
			name:  "one_card_A",
			cards: []Card{{suit: "Spades", rank: "A"}},
			want:  Score([]int{1, 11}),
		},
		{
			name:  "multiple_cards_with_one_A",
			cards: []Card{{suit: "Spades", rank: "7"}, {suit: "Hearts", rank: "3"}},
			want:  Score([]int{10}),
		},
		{
			name:  "multiple_cards_with_one_A",
			cards: []Card{{suit: "Spades", rank: "A"}, {suit: "Hearts", rank: "3"}},
			want:  Score([]int{4, 14}),
		},
		{
			name: "multiple_cards_with_multiple_A",
			cards: []Card{
				{suit: "Spades", rank: "A"},
				{suit: "Hearts", rank: "3"},
				{suit: "Diamonds", rank: "A"},
			},
			want: Score([]int{5, 15, 15, 25}),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := &Hand{cards: test.cards}
			got := h.Score()
			testScoreEqual(t, got, test.want)
		})
	}
}

func TestHand_BestScore(t *testing.T) {
	tests := []struct {
		name  string
		cards []Card
		want  int
	}{
		{
			name:  "empty",
			cards: []Card{},
			want:  0,
		},
		{
			name:  "one_not_burst",
			cards: []Card{{suit: "Spades", rank: "3"}},
			want:  3,
		},
		{
			name:  "multiple_not_burst",
			cards: []Card{{suit: "Spades", rank: "3"}, {suit: "Hearts", rank: "A"}},
			want:  14,
		},
		{
			name: "one_burst",
			cards: []Card{
				{suit: "Spades", rank: "J"},
				{suit: "Hearts", rank: "8"},
				{suit: "Diamonds", rank: "5"},
			},
			want: 23,
		},
		{
			name: "multiple_burst",
			cards: []Card{
				{suit: "Spades", rank: "J"},
				{suit: "Hearts", rank: "Q"},
				{suit: "Diamonds", rank: "5"},
				{suit: "Clubs", rank: "A"},
			},
			want: 26,
		},
		{
			name: "burst_and_not_burst",
			cards: []Card{
				{suit: "Spades", rank: "J"},
				{suit: "Hearts", rank: "8"},
				{suit: "Clubs", rank: "A"},
			},
			want: 19,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := &Hand{cards: test.cards}
			got := h.BestScore()
			if got != test.want {
				t.Errorf("best score wrong. want=%d, got=%d", test.want, got)
			}
		})
	}
}
