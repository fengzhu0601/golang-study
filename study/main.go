package main

import (
	"unsafe"
	"fmt"
	"math/rand"
)

func foo() *int{
	var x int
	return &x
}

func bar() int{
	x := new(int)
	*x = 1
	return *x
}

type small struct {
	a int32
	b string
	c float32
	d []int32
}
func t() *small {
	x := new(small)
	x.a = 1
	x.b = "nihao"
	x.c = 3.1
	x.d = make([]int32, 0, 10)
	println(unsafe.Sizeof(x))
	return x
}

type student struct {
	a int32
	//b [100]int
	b [100]*int
}

func printStrudent() student {
	x := new(student)
	x.a = 1
	for i := 0; i< len(x.b); i++ {
		x.b[i] = &i
	}

	println(unsafe.Sizeof(*x))
	println(&x.b)
	return *x
}

func ps() string {
	s := new(string)
	*s = "aaaa"
	return *s
}

func keyi() {
	var t int
	t = 1
	fmt.Println(t)
}

func cross( a int) {
	a = a+1
}

func arr() int{
	b := [100]int{}
	b[0] = 1
	return b[0]
}

func main () {
	//foo()
	bar()
	//s := "abc"
	//fmt.Println(s)
	ps()
	printStrudent()
	keyi()
	var a int = 5
	cross(a)

	arr()

	//s := make([]byte, 1, 63 * 1024)
	//_ = s
	s := make([]byte, 1, rand.Int31n(10))
	_ = s

	var m []int32
	_ = m
	m = append(m , 1)
	for {}
}
