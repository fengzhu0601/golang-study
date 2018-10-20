package main

import "fengzhu/golang-study/import_cycles/a"

func main(){
	o := a.NewA()
	o.PrintA()
}
