package main

import(
    "fmt"
	"os"
	"golang.org/x/net/html"
)

func main(){
	doc,err := html.Parse(os.Stdin)
	fmt.Println(doc.Type,err)
}