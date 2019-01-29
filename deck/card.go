package deck

import "image/color"

const (
	Skip = "skip"
	Wild = "wild"
)

/*
	TODO: Card Color
		- Change card color format with in deck
 */

type Card struct {
	Name  string
	Color Color
	Score uint8
}

type Color struct {
	Name string
	Color color.RGBA64
}