package card

import (
	"math/rand"
	"strconv"
	"time"
)

type Score []int

func NewScore(ints ...int) Score {
	var score Score
	for _, i := range ints {
		score = append(score, i)
	}
	return score
}

func (s Score) Plus(os Score) Score {
	if s == nil {
		return os
	}

	var score Score
	for _, ss := range s {
		for _, oss := range os {
			score = append(score, ss+oss)
		}
	}
	return score
}

func (s Score) Min() int {
	if len(s) == 0 {
		return 0
	}

	min := s[0]
	for _, ss := range s[1:] {
		if ss < min {
			min = ss
		}
	}
	return min
}

func (s Score) Max() int {
	if len(s) == 0 {
		return 0
	}

	max := s[0]
	for _, ss := range s[1:] {
		if ss > max {
			max = ss
		}
	}
	return max
}

type Card struct {
	suit string
	rank string
}

func (c Card) String() string {
	// https://blog.pokercopilot.com/2017/01/how-to-use-card-suit-symbols-%E2%99%A5%E2%99%A6%E2%99%A0%E2%99%A3-on-your-computer
	switch c.suit {
	case "Spades":
		return "♠" + c.rank
	case "Diamonds":
		return "♦" + c.rank
	case "Clubs":
		return "♣" + c.rank
	case "Hearts":
		return "♥" + c.rank
	}
	return ""
}

func (c Card) Score() Score {
	switch c.rank {
	case "A":
		return NewScore(1, 11)
	case "J", "Q", "K":
		return NewScore(10)
	default:
		p, _ := strconv.Atoi(c.rank)
		return NewScore(p)
	}
}

func shuffle(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := len(cards) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}
