package phase

import (
	"github.com/Tskken/Go10/player"
)

type Phases []Phasing

func Init() Phases {
	return []Phasing{
		new(One),
		new(Two),
		new(Three),
		new(Four),
		new(Five),
		new(Six),
		new(Seven),
		new(Eight),
		new(Nine),
		new(Ten),
	}
}

type One struct{}

/*
	TODO: One.Check():
		- Comment code
		- May be possible to simplify wUsed math
*/
func (One) Check(hand player.Hand, usWild ...string) bool {
	pass, _ := hand.SetCount(2, usWild...)

	return pass
}

type Two struct{}

/*
	TODO: Two.Check()
		- Change code to handle 2 different types.
		- Current setup will make set and run overlap with the same cards which could create issues.
*/
func (Two) Check(hand player.Hand, usWild ...string) bool {
	pass1, hand := hand.SetCount(1, usWild...)

	/*wCount := 0

	if len(usWild) > 0 && len(usWild) <= wCount {
		for _, c := range usWild {
			hand = append(hand,
				deck.Card{
					Name: c,
					Color: map[string]color.RGBA64{},
				})
		}
	}*/

	pass2, hand := hand.RunCount(1)

	return pass1 && pass2
}

type Three struct{}

func (Three) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}

type Four struct{}

func (Four) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}

type Five struct{}

func (Five) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}

type Six struct{}

func (Six) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}

type Seven struct{}

func (Seven) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}

type Eight struct{}

func (Eight) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}

type Nine struct{}

func (Nine) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}

type Ten struct{}

func (Ten) Check(hand player.Hand, usWild ...string) bool {
	panic("implement me")
}
