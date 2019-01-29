package deck

type Discard struct {
	Cards []Card
}

func InitDiscard() *Discard {
	return &Discard{
		Cards: make([]Card, 0),
	}
}

func (d *Discard) Discard(card ...Card) {
	d.Cards = append(d.Cards, card...)
}

func (d *Discard) ReShuffle(deck *Deck) {
	deck.Cards = append(deck.Cards, d.Cards...)
	d.Cards = make([]Card, 0)
}

func (d *Discard) Draw() Card {
	card := d.Cards[len(d.Cards)-1]

	d.Cards = d.Cards[:len(d.Cards)-1]

	return card
}
