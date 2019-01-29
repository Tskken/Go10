package phase

import (
	"Go10/deck"
	"image/color"
	"testing"
)

func TestOne_Check(t *testing.T) {
	one := new(One)

	hand := make([]deck.Card, 0)

	hand = append(hand,
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "2",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "2",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "wild",
			Color: color.RGBA64{},
		},
		)

	t.Run("Basic check", func(t *testing.T) {
		if !one.Check(hand) {
			t.Fail()
		}
	})

	t.Run("Pick card for wild check", func(t *testing.T) {
		if !one.Check(hand, "2") {
			t.Fail()
		}
	})

	hand = append(hand,
		deck.Card{
			Name: "3",
			Color: color.RGBA64{},
		},
	)

	t.Run("Pick card for wild check 2", func(t *testing.T) {
		if !one.Check(hand, "2") {
			t.Fail()
		}
	})

	hand = append(hand,
		deck.Card{
			Name: "wild",
			Color: color.RGBA64{},
		},
		)

	t.Run("Pick card for wild check 3", func(t *testing.T) {
		if !one.Check(hand, "2") {
			t.Fail()
		}
	})

	t.Run("Pick card for wild check 4", func(t *testing.T) {
		if !one.Check(hand, "2", "3") {
			t.Fail()
		}
	})

	t.Run("Pick card for wild check 5", func(t *testing.T) {
		if !one.Check(hand, "3") {
			t.Fail()
		}
	})

	t.Run("Pick card for wild check 6", func(t *testing.T) {
		if !one.Check(hand, "3", "3") {
			t.Fail()
		}
	})

	hand2 := make([]deck.Card, 0)

	hand2 = append(hand2,
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "2",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "3",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "wild",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "wild",
			Color: color.RGBA64{},
		},
	)

	t.Run("Pick card for wild check 7", func(t *testing.T) {
		if !one.Check(hand2, "3") {
			t.Fail()
		}
	})

	t.Run("Pick card for wild check 8", func(t *testing.T) {
		if one.Check(hand2, "3", "2") {
			t.Fail()
		}
	})

	t.Run("Pick more card then wilds", func(t *testing.T) {
		if !one.Check(hand2, "1", "2", "3", "3") {
			t.Fail()
		}
	})

	hand2 = append(hand2,
		deck.Card{
			Name:"skip",
			Color:color.RGBA64{},
		},
		deck.Card{
			Name:"wild",
			Color:color.RGBA64{},
		})

	t.Run("Check for 3 wild for non existing numbers", func(t *testing.T) {
		if !one.Check(hand2, "9", "9", "9") {
			t.Fail()
		}
	})

	t.Run("Base check 2", func(t *testing.T) {
		if !one.Check(hand2) {
			t.Fail()
		}

	})

	hand2 = append(hand2,
		deck.Card{
			Name:"1",
			Color:color.RGBA64{},
		})

	t.Run("Base check 3", func(t *testing.T) {
		if !one.Check(hand2) {
			t.Fail()
		}

	})

}

func TestTwo_Check(t *testing.T) {
	two := new(Two)

	hand := make([]deck.Card, 0)

	hand = append(hand,
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "1",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "2",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "3",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "4",
			Color: color.RGBA64{},
		},
		deck.Card{
			Name: "wild",
			Color: color.RGBA64{},
		},
	)

	t.Run("Basic check", func(t *testing.T) {
		if !two.Check(hand) {
			t.Fail()
		}
	})
}
