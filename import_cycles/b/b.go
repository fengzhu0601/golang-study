package b

import (
	"fmt"
	"fengzhu/golang-study/import_cycles/a"
)

type B struct {

}

func (b B) PrintB() {
	fmt.Println(b)
}

func NewB() *B{
	b := new(B)
	return b
}

func RequireA() {
	o := a.NewA()
	o.PrintA()
}
