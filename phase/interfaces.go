package phase

import (
	"github.com/Tskken/Go10/player"
)

type Phasing interface {
	Check(hand player.Hand, usWild ...string) bool
}
