package main

import (
	"fmt"
	"log"

	"spline/internal"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	keys []ebiten.Key
}

var s []internal.BasicSpline
var selectedNode int
var selectedSpline int
var agent float64

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	for _, v := range g.keys {
		if inpututil.IsKeyJustPressed(v) {
			switch v {
			case ebiten.KeyA:
				selectedNode--
			case ebiten.KeyD:
				selectedNode++
			case ebiten.KeyZ:
				selectedSpline--
			case ebiten.KeyC:
				selectedSpline++
			}
			if selectedSpline >= len(s) {
				selectedSpline = 0
			} else if selectedSpline < 0 {
				selectedSpline = len(s) - 1
			}
			if selectedNode >= s[selectedSpline].Length() {
				selectedNode = 0
			} else if selectedNode < 0 {
				selectedNode = s[selectedSpline].Length() - 1
			}
		} else {
			switch v {
			case ebiten.KeyArrowLeft:
				s[selectedSpline].DecrementPointX(selectedNode, 1)
			case ebiten.KeyArrowRight:
				s[selectedSpline].IncrementPointX(selectedNode, 1)
			case ebiten.KeyArrowUp:
				s[selectedSpline].DecrementPointY(selectedNode, 1)
			case ebiten.KeyArrowDown:
				s[selectedSpline].IncrementPointY(selectedNode, 1)
			case ebiten.KeyQ:
				agent -= 0.05
			case ebiten.KeyW:
				agent += 0.05
			}
			if agent >= float64(s[selectedSpline].Length())-3 {
				agent = 0
			} else if agent < 0 {
				agent = float64(s[selectedSpline].Length()) - 3.001
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range s {
		v.DrawIt(screen, selectedNode, agent)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Node: %v Spline: %v", selectedNode, selectedSpline))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	var ps1 []internal.Point2D
	for i := 1; i < 8; i++ {
		ps1 = append(ps1, internal.NewPoint2D(20*float64(i), 40))
	}
	s1 := internal.NewBasicSpline(false, ps1...)

	var ps2 []internal.Point2D
	for i := 1; i < 14; i++ {
		ps2 = append(ps2, internal.NewPoint2D(20*float64(i), 50))
	}
	s2 := internal.NewBasicSpline(false, ps2...)

	s = append(s, s1, s2)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
