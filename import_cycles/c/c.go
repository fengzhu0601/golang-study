package c

import (
	"fengzhu/golang-study/import_cycles/a"
	"fengzhu/golang-study/import_cycles/b"
)

func PrintC()  {
	o := a.NewA()
	b.RequireA(o)
}
