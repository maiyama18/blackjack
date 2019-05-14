package main

import (
	"fmt"

	"github.com/mui87/blackjack/card"
)

func main() {
	fmt.Println("game start")

	deck := card.NewDeck()
	c1, _ := deck.Draw()
	c2, _ := deck.Draw()
	c3, _ := deck.Draw()
	c4, _ := deck.Draw()

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
}
