package deck

import "image/color"

const (
	Skip = "skip"
	Wild = "wild"
)

type Card struct {
	Name  string
	Color color.RGBA64
}
