// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func shallowCopy(obj map[string]int) map[string]int{
    maps := make(map[string]int)
    for name, value := range obj{
        maps[name]=value
    }
    return maps
}

func toString(obj []map[string]int) string{
    var s string
    for _,c := range obj{
        for key,value := range c{
            s += key+fmt.Sprintf("%d",value)
        }
    }
    return s
}


func main() {
    var ip string = "AAABBCCGABBB"
    
    arr := make([]map[string]int,0)
    var temp rune 
    map1 := make(map[string]int)

    for _,c := range ip{
        if temp == c{
            map1[string(c)]++
        } else {
            if temp != 0{
                arr = append(arr,shallowCopy(map1))
                delete(map1,string(temp))
            }
        temp=c
        map1[string(c)]++
        }
    }
    arr = append(arr,shallowCopy(map1))


    fmt.Println(toString(arr))
}