package deck

import (
	"encoding/json"
	"image/color"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	ConfigPath = "../Go10/configs/DeckConfig.json"
	ColorPath  = "../Go10/configs/CardColors.json"
)

type Deck struct {
	Cards []Card
}

func loadColorData() ([]color.RGBA64, error) {
	type Color struct {
		R uint16
		G uint16
		B uint16
		A uint16
	}

	type CardColor struct {
		Name  string
		Color Color
	}

	colorPath, err := filepath.Abs(ColorPath)
	if err != nil {
		return nil, err
	}

	colorFile, err := os.Open(colorPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := colorFile.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	colorDec := json.NewDecoder(colorFile)

	colors := make([]CardColor, 0)

	err = colorDec.Decode(&colors)
	if err != nil {
		log.Println(err)
	}

	rgbaColors := make([]color.RGBA64, 0, len(colors))

	for _, c := range colors {
		rgbaColors = append(rgbaColors, color.RGBA64{
			R: c.Color.R,
			G: c.Color.G,
			B: c.Color.B,
			A: c.Color.A,
		})
	}

	return rgbaColors, nil
}

func loadConfigData(rgbaColors []color.RGBA64) ([]Card, error) {
	type CardType struct {
		Name  string
		Color bool
		Count uint8
	}

	type ConfigDeck struct {
		CardCount int
		CardTypes []CardType
	}

	configPath, err := filepath.Abs(ConfigPath)
	if err != nil {
		return nil, err
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := configFile.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	configDec := json.NewDecoder(configFile)

	configDeck := new(ConfigDeck)

	err = configDec.Decode(&configDeck)
	if err != nil {
		return nil, err
	}

	cards := make([]Card, 0, configDeck.CardCount)

	for _, ct := range configDeck.CardTypes {
		colorIndex := 0
		for count := ct.Count; count > 0; count-- {
			if colorIndex > 3 {
				colorIndex = 0
			}

			if ct.Color {
				cards = append(cards, Card{
					Name:  ct.Name,
					Color: rgbaColors[colorIndex],
				})
				colorIndex++
			} else {
				cards = append(cards, Card{
					Name:  ct.Name,
					Color: rgbaColors[4],
				})
			}
		}
	}

	return cards, nil
}

func InitDeck() (*Deck, error) {
	rgbaColors, err := loadColorData()
	if err != nil {
		return nil, err
	}

	cards, err := loadConfigData(rgbaColors)
	if err != nil {
		return nil, err
	}

	return &Deck{
		Cards: cards,
	}, nil
}

func (d *Deck) Draw() Card {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	i := rng.Intn(len(d.Cards))

	card := d.Cards[i]

	d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)

	return card
}
