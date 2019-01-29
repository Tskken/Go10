package phase

import (
	"Go10/player"
)

type Phasing interface {
	Check(hand player.Hand, usWild ...string) bool
}
