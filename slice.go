package main

import "fmt"

func main() {
//	p := []int{2, 3, 5, 7, 11, 13}
//	fmt.Println("p====", p)
//	fmt.Println("p[1:4] ==", p[1:4])
//
//	fmt.Println("p[:3] ==", p[:3])
//	fmt.Println("p[4:] ==", p[4:])
//	for i := 0; i < len(p); i++ {
//		fmt.Printf("p[%d] == %d\n", i, p[i])
//	}
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
	s, len(x), cap(x), x)
}
