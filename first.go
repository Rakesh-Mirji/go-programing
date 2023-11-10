package main

import (
	"fmt"
	"reflect"
	"time"
)

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	fmt.Println("hello world")

	for {
		fmt.Println("qwee")
		break
	}

	if num := 2; num > 0 {
		fmt.Println(num, "is correct")
	} else {
		fmt.Println(num, "not correct")
	}

	a, b := 1, 2
	a = 4
	fmt.Println(a, b)

	t := time.Now().Weekday()

	fmt.Println(reflect.TypeOf(t), t, time.Saturday, time.Sunday)

	// var x error
	// fmt.Println(x)

	x := 10
	p := &x

	fmt.Println(*p)

	*p = 2

	fmt.Println(x)

	fmt.Println(f() == f())

	v := 2
	fmt.Println(incr(&v))

	var r rune = 'A'
	fmt.Println(r)

}	 