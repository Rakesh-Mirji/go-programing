package main

import (
	"fmt"
	"math"
	"image/color")

type Point struct{X, Y float64}

type Path []Point


//Function
func Distance(p,q Point)float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Method functions
func (p Point) Distance(q Point) float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) scaleBy(factor float64){
	p.X += factor
	p.Y += factor
}

func (path Path) Distance() float64 {
	sum:=0.0
	for i := range path{
		if i>0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main(){
	p := Point{1,2}
	q := Point{4,6}

	
	
	fmt.Println(Distance(p,q))  //FUNCTION CALL
	fmt.Println(p.Distance(q))	//METHOD CALL
	
	perim := Path{
		{1,1},{5,1},{5,4},{1,1},
	}
	fmt.Println(perim.Distance())
	

	r := Point{1,2}
	(&r).scaleBy(2)
	// r.scaleBy(3)
	fmt.Println(r)
	
	Point{1,2}.Distance(q)
	// Point{1,2}.scaleBy(3) //Error : cannot call pointer method scaleBy on Point
	

}