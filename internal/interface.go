package internal

import "github.com/hajimehoshi/ebiten/v2"

type Spline interface {
	DrawIt(img *ebiten.Image, selected int, agent float64)
	GetBasicSplinePoint(t float64) Point2D
	GetBasicSplineGradient(t float64) Point2D
	Length() int
}
