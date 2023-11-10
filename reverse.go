// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func rev(str string) string{
    count := len(str)-1
    if count > 0 {
        return str[count : ] + rev( str[ 0 : count ])
    }
    return str
}

func main() {
    var word string = "Pavan"
    fmt.Println(rev(word))
}