package a

import (
	"fmt"
	"fengzhu/golang-study/import_cycles/b"
)

type A struct {

}

func (a A) PrintA(){
	fmt.Println(a)
}

func NewA() *A{
	a := new(A)
	return a
}

func RequireB(){
	o := b.NewB()
	o.PrintB()
}
