package main

import "fmt"

type rect struct{
	height,width int
}


func (r *rect) area()int{
	r.height=20
	return r.height * r.width
}

func (r rect) param()int{
	r.height=40
	return 2*r.height + 2*r.width
}

type maps map[int]string

func(rec *maps) f() maps{
	temp := *rec
	temp[2]="sd"

	return *rec
}

func(rec maps) g() maps{
	rec[3]="main"
	return rec
}


func main(){
	r:= rect{width:10,height:2}

	obj := maps{1:"string"}

	fmt.Println(obj.f(),obj)
	fmt.Println(obj.g(),obj)

	fmt.Printf("%#v\n",r)
	fmt.Println("Param->",r.param(),"Area->", r.area())
	fmt.Printf("%#v\n",r)
}