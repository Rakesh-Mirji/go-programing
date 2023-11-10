package main
import "fmt"
func largestNum(arrs [][]int)[]int{
    var retArr []int
    for _,c := range arrs{
        var largest int = c[0]
        for _,c := range c{
            if c > largest{
                largest = c
            }
        }
        retArr = append(retArr,largest)
    }
    return retArr
}
func main() {
arrs := [][]int { []int{10,5,4,2}, []int{100,200,999,300}, []int{-1,-10,-6,-3}, []int{99,999,10000,9999} }

  fmt.Println(largestNum(arrs))
}