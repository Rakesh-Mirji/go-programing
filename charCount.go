package main
import ("fmt"
"strconv"
)

func main() {
    var ip string = "AAABBCCGABBB"
    var out string
    var temp string
    for i, c := range ip{
        if len(temp)>0 && temp[:1] != string(c){
            out = out + temp[:1] + strconv.Itoa(len(temp))
            temp = string(c)
        } else {
            temp = temp + string(c)
        }
        if i == len(ip)-1{
            out = out + temp[:1] + strconv.Itoa(len(temp))
        }
    }
    fmt.Printf("input  %s\noutput %s",ip,out)
}
