package main

import "fmt"

func main(){
	ab:=5
	fmt.Printf("%08b\n",ab)

	medals := []string{"gold","silver","bronze"}

	fmt.Println(medals,len(medals))

	for i:= len(medals)-1; i >= 0; i--{
		fmt.Println(medals[i])
	}

	x := "hello"

	for _,x := range x{
		x := x + 'A' - 'a'
		fmt.Printf("%c,%v|\n",x,x)
	}

    for i, c := range "go" {
        fmt.Printf("%d\t%q\t%d\n ", i, c, c)
    }

	var a [3]int = [3]int{1,2,3}


	q := [...]int{1,2,3}

	fmt.Println(a,q)
	fmt.Printf("%T\n",q)

}