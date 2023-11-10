// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import ("fmt"
       "strconv")

func findMax(arr []int, n int) int {
   if n == 1 {
      return arr[0]
   }

   max := findMax(arr, n-1)

   if arr[n-1] > max {
      return arr[n-1]
   }
   return max
}

func seriesSum(num string, value int)int{
    count := len(num)-1
    if count>0{
        i,_ := strconv.Atoi(string(num[count]))
        return i*value + seriesSum(num[0:count],value)
    }
    i,_ := strconv.Atoi(string(num[count]))
    return i*value
}

func series(num string,value int)[]int{
    var arr []int
    for i := range num {
        if i+3 <= len(num) {
            arr = append(arr,seriesSum(num[i:i+3],value))
            fmt.Println(num[i:i+3])
        }
    }
    return arr
}

func main() {
	// change a series in terms of 3 nums and multiply & add the each number with a given value
  var num = "0123456789"
  var value int = 2
  arr := series(num,value)
  fmt.Println( "product*value->",arr,"\tmax->",findMax( arr,len(arr) ) )
}