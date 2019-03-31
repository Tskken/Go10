package deck

import (
	"github.com/Tskken/Go10/configs"
	"image/color"
	"math/rand"
	"sync"
	"time"
)

type Go10Deck struct {
	DrawPile    []Card
	DiscardPile []Card

	sync.Mutex
}

func NewDeck() (*Go10Deck, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}

	cards := make([]Card, 0)

	for _, c := range config.CardData {
		for count := c.Count; count > 0; count-- {
			if c.HasColor {
				for _, cl := range config.ColorData {
					cards = append(cards, Card{
						Name: c.Name,
						Color: color.RGBA64{
							R: cl.Color.R,
							G: cl.Color.G,
							B: cl.Color.B,
							A: cl.Color.A,
						},
						Score: c.Score,
					})
				}
			} else {
				cards = append(cards, Card{
					Name:  c.Name,
					Color: color.RGBA64{},
					Score: c.Count,
				})
			}
		}
	}

	return &Go10Deck{
		DrawPile:    cards,
		DiscardPile: make([]Card, 0),
	}, nil
}

func (d *Go10Deck) DrawFromDeck() (c Card) {
	d.Lock()

	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(d.DrawPile))

	c = d.DrawPile[i]

	d.DrawPile = append(d.DrawPile[:i], d.DrawPile[i+1:]...)

	d.Unlock()

	return
}

func (d *Go10Deck) DrawFromDiscard() (c Card) {
	d.Lock()

	c = d.DiscardPile[len(d.DiscardPile)-1]

	d.DiscardPile = d.DiscardPile[:len(d.DiscardPile)-2]

	d.Unlock()

	return
}

func (d *Go10Deck) DiscardCard(card ...Card) {
	d.DiscardPile = append(d.DiscardPile, card...)
}

func (d *Go10Deck) ReShuffle() {
	d.DrawPile = append(d.DrawPile, d.DiscardPile...)
	d.DiscardPile = make([]Card, 0)
}
