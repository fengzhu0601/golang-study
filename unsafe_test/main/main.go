package main

import (
	"fmt"
	"fengzhu/golang-study/unsafe_test/b"
)

func main() {
	s := b.Greet("world")
	fmt.Println(s)
	s = b.Hi("world")
	fmt.Println(s)
}
