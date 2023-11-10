//  Closure function

package main


import "fmt"

func intSeq() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

func square(i int) func() int{
	return func() int{
		i*=i
		return i
	}
}

func main(){

	nextInt := intSeq()

	nextSquare := square(2)
	
	fmt.Println(nextInt(),nextSquare())
    fmt.Println(nextInt(),nextSquare())
    fmt.Println(nextInt(),nextSquare())
	
	fmt.Println(intSeq()())
	
	newInts := intSeq()

    fmt.Println(newInts())
}