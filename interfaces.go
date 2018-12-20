package main

import (
	"fmt"
	//"math"
	"os"
)

//type Abser interface {
//	Abs() float64 
//}
//
//func main() {
//	var a Abser 
//	f := MyFloat(-math.Sqrt2)
//	v := Vertex{3, 4}
//	a = f
//	a = &v
//	fmt.Println(a.Abs())
//}
//
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//type Vertex struct {
//	X, Y float64 
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y) 
//}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWrite interface {
	Reader
	Writer 
}

func main() {
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello writer\n")
}
