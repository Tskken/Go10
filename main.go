package main

import (
	"Go10/deck"
	"Go10/phase"
	"log"
)

func main() {
	d, err := deck.InitDeck()
	if err != nil {
		log.Println(err)
	}

	//discard := deck.InitDiscard()

	phases := phase.Init()

	hand := make([]deck.Card, 0)

	for i:= 0; i < 50; i++ {
		hand = append(hand, d.Draw())
	}

	log.Println(hand)

	log.Println(phases[0].Check(hand))

	log.Println(hand)

	log.Println(phases[1].Check(hand))
}
