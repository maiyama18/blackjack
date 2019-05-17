package card

import (
	"testing"
)

func TestNewScore(t *testing.T) {
	tests := []struct {
		name string
		ints []int
		want Score
	}{
		{
			name: "one_int",
			ints: []int{3},
			want: Score([]int{3}),
		},
		{
			name: "two_ints",
			ints: []int{3, 4},
			want: Score([]int{3, 4}),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewScore(test.ints...)
			testScoreEqual(t, test.want, got)
		})
	}
}

func TestScore_Plus(t *testing.T) {
	tests := []struct {
		name  string
		left  Score
		right Score
		want  Score
	}{
		{
			name:  "one_elem+one_elem",
			left:  NewScore(3),
			right: NewScore(4),
			want:  NewScore(7),
		},
		{
			name:  "nil+one_elem",
			left:  nil,
			right: NewScore(4),
			want:  NewScore(4),
		},
		{
			name:  "one_elem+two_elems",
			left:  NewScore(3),
			right: NewScore(4, 5),
			want:  NewScore(7, 8),
		},
		{
			name:  "one_elem+two_elems",
			left:  NewScore(3, 4),
			right: NewScore(5, 6),
			want:  NewScore(8, 9, 9, 10),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.left.Plus(test.right)
			testScoreEqual(t, test.want, got)
		})
	}
}

func TestScore_MaxMin(t *testing.T) {
	tests := []struct {
		name    string
		score   Score
		wantMax int
		wantMin int
	}{
		{
			name:    "nil",
			score:   Score([]int{}),
			wantMax: 0,
			wantMin: 0,
		},
		{
			name:    "one_elem",
			score:   NewScore(2),
			wantMax: 2,
			wantMin: 2,
		},
		{
			name:    "three_elems",
			score:   NewScore(5, 2, 8),
			wantMax: 8,
			wantMin: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.score.Max() != test.wantMax {
				t.Fatalf("max wrong. want=%d, got=%d", test.wantMax, test.score.Max())
			}
			if test.score.Min() != test.wantMin {
				t.Fatalf("min wrong. want=%d, got=%d", test.wantMin, test.score.Min())
			}
		})
	}
}

func testScoreEqual(t *testing.T, want, got Score) {
	if len(got) != len(want) {
		t.Fatalf("length of score wrong. want=%d, got=%d", len(want), len(got))
	}
	for i, w := range want {
		g := got[i]
		if g != w {
			t.Fatalf("%d-th element wrong. want=%d, got=%d", i, w, g)
		}
	}
}

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want string
	}{
		{
			name: "Spades_A",
			card: Card{suit: "Spades", rank: "A"},
			want: "♠A",
		},
		{
			name: "Diamonds_7",
			card: Card{suit: "Diamonds", rank: "7"},
			want: "♦7",
		},
		{
			name: "Hearts_10",
			card: Card{suit: "Hearts", rank: "10"},
			want: "♥10",
		},
		{
			name: "Clubs_K",
			card: Card{suit: "Clubs", rank: "K"},
			want: "♣K",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.card.String()
			if got != test.want {
				t.Fatalf("card string wrong. want=%s, got=%s", test.want, got)
			}
		})
	}
}

func TestCard_Score(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want Score
	}{
		{
			name: "A",
			card: Card{suit: "Spades", rank: "A"},
			want: Score([]int{1, 11}),
		},
		{
			name: "Q",
			card: Card{suit: "Spades", rank: "Q"},
			want: Score([]int{10}),
		},
		{
			name: "7",
			card: Card{suit: "Spades", rank: "7"},
			want: Score([]int{7}),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.card.Score()
			testScoreEqual(t, test.want, got)
		})
	}
}
