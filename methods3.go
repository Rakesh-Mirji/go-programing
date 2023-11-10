package main

import ("fmt"
		"math"
		"image/color")

type Point struct {X,Y float64}
type ColoredPoint struct{
	Point
	Color color.RGBA
}

//Method functions
func (p Point) Distance(q Point) float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) scaleBy(factor float64){
	p.X += factor
	p.Y += factor
}

func (p ColoredPoint) Distance(q Point) float64{
	return p.Point.Distance(q)
}

func (p *ColoredPoint) scaleBy(factor float64){
	p.Point.scaleBy(factor)
}

func main(){
	// Composing types by struct embedding

	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255,0,0,255}
	blue := color.RGBA{0,0,255,255}

	cp = ColoredPoint{Point{1,2},red}
	var cq = ColoredPoint{Point{5,4},blue}

	fmt.Println(cp.Point.Distance(cq.Point))
	cp.scaleBy(2)
	cp.scaleBy(6)
	fmt.Println(cp.Distance(cq.Point))	
}