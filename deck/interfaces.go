package deck

type Decking interface {
	Draw() Card
}

type Discarding interface {
	Decking

	Discard(...Card)
	ReShuffle(*Deck)
}
