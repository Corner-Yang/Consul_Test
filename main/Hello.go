package main

import "fmt"

var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	g, h := 123, "hello"
	fmt.Printf("hi Doct_Y\n")
	println(a, b, c, d, e, f, g, h)
	println("i=", i)
	println("j=", j)
	println("k=", k)
	println("l=", l)

}
