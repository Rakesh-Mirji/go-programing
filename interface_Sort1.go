package main

import ("fmt"
		// "sort"
		"time"
		// "io"
		"os"
		// "tabwriter"	
	)

type Interface interface{
	Len()int
	Less(i,j int)bool
	Swap(i,j int)
}

type StringSlice []string

func(p StringSlice) Len()int{
	return len(p)
}

func(p StringSlice) Less(i,j int)bool{
	return p[i] < p[j]
}

func(p StringSlice) Swap(i,j int){
	p[i], p[j] = p[j], p[i]
}

type Track struct{
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

func length(s string) time.Duration{
	d,err := time.ParseDuration(s)
	if err != nil{
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track){
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(*tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _,t := range tracks{
		fmt.Fprintf(tw,format,t.Title,t.Artist,t.Album,t.Year,t.Length)
	}
	tw.flush()
}

func main(){
	var name StringSlice = []string{"10000","12","1111"}
	
	
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	fmt.Println(name,tracks)
	printTracks(tracks)
}