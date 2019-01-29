package main

import (
	"github.com/Tskken/Go10/deck"
	"github.com/Tskken/Go10/phase"
	"log"
)

func main() {
	d, err := deck.InitDeck()
	if err != nil {
		log.Println(err)
	}

	//discard := deck.InitDiscard()

	phases := phase.Init()

	for _, p := range phases {
		hand := make([]deck.Card, 0)

		for i := 0; i < 10; i++ {
			hand = append(hand, d.Draw())
		}

		log.Println(hand)

		log.Println(p.Check(hand))
	}
}
