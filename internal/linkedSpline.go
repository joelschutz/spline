package internal

// import (
// 	"image/color"
// 	"math"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
// )

// type LinkedSpline struct {
// 	parentSpline Spline
// 	offset       float64
// 	freedom      float64
// }

// func (ls LinkedSpline) Length() int {
// 	return (ls.parentSpline.Length() * 2) - 2
// }

// func NewLinkedSpline(looped bool, ps ...Point2D) LinkedSpline {
// 	return LinkedSpline{points: ps, looped: looped}
// }

// func (s LinkedSpline) DrawIt(img *ebiten.Image, selected int, agent float64) {
// 	// Draw Controls
// 	for i, v := range s.points {
// 		c := color.RGBA{255, 255, 255, 255}
// 		if selected == i {
// 			c = color.RGBA{255, 0, 0, 255}
// 		}
// 		ebitenutil.DrawRect(img, float64(v.x)-2, float64(v.y)-2, 4, 4, c)
// 		// ebitenutil.DebugPrintAt(img, fmt.Sprintf("%.0f:%.0f", v.x, v.y), int(v.x), int(v.y))
// 	}

// 	offset := 3
// 	if s.looped {
// 		offset = 0
// 	}
// 	// Draw Gradient
// 	for i := 0; i < len(s.points)-offset; i++ {
// 		g := s.GetLinkedSplineGradient(float64(i))
// 		r := math.Atan2(-g.y, g.x)
// 		// img.Set(int(g.x+s.points[i].x), int(g.y+s.points[i].y), color.RGBA{255, 255, 255, 255})
// 		ebitenutil.DrawLine(img, (10*math.Sin(r) + s.points[i+1].x), (10*math.Cos(r) + s.points[i+1].y), (-10*math.Sin(r) + s.points[i+1].x), (-10*math.Cos(r) + s.points[i+1].y), color.RGBA{0, 255, 0, 255})

// 	}
// 	// Draw LinkedSpline
// 	for i := 0.; i < float64(len(s.points)-offset); i += 0.001 {
// 		// fmt.Println(len(s.points))
// 		p := s.GetLinkedSplinePoint(i)
// 		img.Set(int(p.x), int(p.y), color.RGBA{0, 0, 255, 255})
// 	}

// 	// Draw Agent
// 	g := s.GetLinkedSplineGradient(agent)
// 	r := math.Atan2(-g.y, g.x) + math.Pi/2
// 	p := s.GetLinkedSplinePoint(agent)

// 	vx := (20*math.Sin(r) + p.x)
// 	vy := (20*math.Cos(r) + p.y)
// 	ebitenutil.DrawLine(img, (math.Sin(r) + p.x), (+p.y), vx, vy, color.RGBA{255, 255, 0, 255})
// 	ebitenutil.DrawLine(img, vx, vy, vx-5*math.Cos(r-(math.Pi/4)), vy+5*math.Sin(r-(math.Pi/4)), color.RGBA{255, 255, 0, 255})
// 	ebitenutil.DrawLine(img, vx, vy, vx+5*math.Cos(r+(math.Pi/4)), vy-5*math.Sin(r+(math.Pi/4)), color.RGBA{255, 255, 0, 255})
// 	return
// }

// func (s LinkedSpline) GetLinkedSplinePoint(t float64) Point2D {
// 	// point indexes
// 	var p0, p1, p2, p3 int
// 	if !s.looped {
// 		p1 = int(t) + 1
// 		p2 = p1 + 1
// 		p3 = p2 + 1
// 		p0 = p1 - 1
// 	} else {
// 		p1 = int(t)
// 		p2 = (p1 + 1) % len(s.points)
// 		p3 = (p2 + 1) % len(s.points)
// 		if p1 >= 1 {
// 			p0 = p1 - 1
// 		} else {
// 			p0 = len(s.points) - 1
// 		}
// 	}
// 	// fmt.Println(p1, p3)

// 	// cache of t squared and cubic
// 	t = t - float64(int(t))
// 	tt := t * t
// 	ttt := tt * t

// 	q1 := -ttt + 2*tt - t
// 	q2 := 3*ttt - 5*tt + 2
// 	q3 := -3*ttt + 4*tt + t
// 	q4 := ttt - tt

// 	return Point2D{
// 		x: (s.points[p0].x*q1 + s.points[p1].x*q2 + s.points[p2].x*q3 + s.points[p3].x*q4) / 2,
// 		y: (s.points[p0].y*q1 + s.points[p1].y*q2 + s.points[p2].y*q3 + s.points[p3].y*q4) / 2,
// 	}
// }

// func (s LinkedSpline) GetLinkedSplineGradient(t float64) Point2D {
// 	// point indexes
// 	var p0, p1, p2, p3 int
// 	if !s.looped {
// 		p1 = int(t) + 1
// 		p2 = p1 + 1
// 		p3 = p2 + 1
// 		p0 = p1 - 1
// 	} else {
// 		p1 = int(t)
// 		p2 = (p1 + 1) % len(s.points)
// 		p3 = (p2 + 1) % len(s.points)
// 		if p1 >= 1 {
// 			p0 = p1 - 1
// 		} else {
// 			p0 = len(s.points) - 1
// 		}
// 	}
// 	// fmt.Println(p1, p3)

// 	// cache of t squared and cubic
// 	t = t - float64(int(t))
// 	tt := t * t

// 	q1 := -3*tt + 4*t - 1
// 	q2 := 9*tt - 10*t
// 	q3 := -9*tt + 8*t + 1
// 	q4 := 3*tt - 2*t

// 	return Point2D{
// 		x: (s.points[p0].x*q1 + s.points[p1].x*q2 + s.points[p2].x*q3 + s.points[p3].x*q4) / 2,
// 		y: (s.points[p0].y*q1 + s.points[p1].y*q2 + s.points[p2].y*q3 + s.points[p3].y*q4) / 2,
// 	}
// }

// func (s LinkedSpline) Length() int {
// 	return len(s.points)
// }

// func (s *LinkedSpline) IncrementPointX(index int, rate float64) {
// 	s.points[index].x += rate
// }

// func (s *LinkedSpline) DecrementPointX(index int, rate float64) {
// 	s.points[index].x -= rate
// }

// func (s *LinkedSpline) IncrementPointY(index int, rate float64) {
// 	s.points[index].y += rate
// }

// func (s *LinkedSpline) DecrementPointY(index int, rate float64) {
// 	s.points[index].y -= rate
// }
