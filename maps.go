package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
	  40.68433, -74.39967,
	},
	"Google": Vertex{
		37.4333, -122.09080,
	},
}

func main() {
	fmt.Println(m)
}
