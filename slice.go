package main

import "fmt"


func main(){

	arr1 := []int{1,2,3,4}

	var arr2 []int = []int{5,6,7,8}

	fmt.Println(arr1,arr2)

	arr3 := append(arr1, arr2...)
	fmt.Println(arr3)


	// nil -nil means the arr is not initialized
	// if not nil then the array may be empty of
	var s []int
	fmt.Println("len->",len(s),"arr=",s, s==nil)

	s=nil
	fmt.Println("len->",len(s),"arr=",s, s==nil)

	s=[]int(nil)
	fmt.Println("len->",len(s),"arr=",s, s==nil)	

	s=[]int{}
	fmt.Println("len->",len(s),"arr=",s, s==nil)


	fmt.Printf("types are %T\t%T\t%T\t%T\n",arr1,arr2,arr3,s)
	// make a copy of arr to c
	c:= make([]int, len(arr1))
	fmt.Println(c)
	//copy the contents of arr1 to c 
	copy(c,arr1)
	fmt.Println(c)
}