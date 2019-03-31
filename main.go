package main

import (
	"github.com/Tskken/Go10/deck"
	"github.com/Tskken/Go10/phase"
	"log"
)

func main() {
	//discard := deck.InitDiscard()

	phases := phase.Init()

	for _, p := range phases {
		d, err := deck.NewDeck()
		if err != nil {
			log.Println(err)
		}

		hand := make([]deck.Card, 0)

		for i := 0; i < 10; i++ {
			hand = append(hand, d.DrawFromDeck())
		}

		log.Println(hand)

		log.Println(p.Check(hand))
	}
}
