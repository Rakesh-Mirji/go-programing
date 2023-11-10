package main

import "fmt"

func plus(a ,b int) int{
	return a+b
}

func square(p *int){
	*p = *p * *p
}

func sum(nums ...int){
fmt.Println(nums)
total :=0
for _,num := range nums{
	total+=num
}
fmt.Println(total)

}

func main(){

	num := 10
	square(&num)
	fmt.Println(num)

	res := plus(1,2)
	fmt.Println(res)


	// Variadic function

	sum(1,2,3)
	sum(1,2)

	nums := []int{1,2,4,63,2}

	fmt.Printf("the type of nums is %T\n",nums)
	sum(nums...)
}