package main

import "fmt"


func zero(ptr *[2]byte){
	for i:= range ptr{
		ptr[i] = byte(i)
	}

	// *ptr = [2]byte{3,4}
 }


func main(){
	//slice
	arr1 := []int{1,2,3,4}
	var arr2 []int = []int{5,6,7,8}

	//array
	var a [3]int = [3]int{1,2}
	
	q := [...]int{1,2}
	
	//slice
	var arr3 []int = append(arr1,90)

	fmt.Println(arr1,arr2,arr3)
	
	fmt.Println(a,q)
	fmt.Printf("type-> %T\ttype-> %T\ttype-> %T\n",a,q,arr1)

	const(
		USD int = iota
		EUR
		YEN
	)

	//array
	symbol := [...]string{USD:"$", EUR:"€", YEN:"¥"}
	fmt.Println(YEN ,symbol)

	r:= [...]string{10:"1","2","3","4","5",9:"hello"}
	fmt.Println(r)
 
	bytes := [...]byte{255,12} 

	zero(&bytes)
	fmt.Println(bytes)
	
	//slice
	months := []string{1:"january", 2:"Frbruary", 3:"March", 4:"April", 5:"May", 6:"June", 7:"July", 8:"Auguest", 9:"September", 10:"October", 11:"November", 12:"December" }	

	summer := months[4:7]
	summer[2] = "-----"
	fmtf.Println(months,summer[:5])



}