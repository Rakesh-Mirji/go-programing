package main

import ("fmt")

func main(){
	ages1 := make(map[int]string)

	ages2 := map[int]string{
		1:"anil",
		2:"pavan",
		3:"ravi",
	}

	ages1[0] = "sumit"

	delete(ages2, 1 )

	ages1[3] = ages2[2]+" bob"

	ages1[3] += " and "

	fmt.Println(ages1,ages2)
	fmt.Printf("%T\t%T\n",ages1,ages2)


	for name, age := range ages1{
		fmt.Printf("%d\t%s\n", name, age)
		// fmt.Println(name,age)
	}

	if ages1[9]==""{
		fmt.Println("not Present")
	}


	if args,ok := ages2[3]; ok{
		fmt.Println(args,ok)
	}

	
}