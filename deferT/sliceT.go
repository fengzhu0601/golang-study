package main

import (
	"math/rand"
	"fmt"
)

func main()  {
	//testSlice2()
	tt1()
}
func testSlice() {
    vect := make([]int, 10)
    //1. 赋初值
    for i, _ := range vect {
        vect[i] = i
    }
    for i := 0; i < 3; i++ {
        k := rand.Intn(10)
        vect[k] = -1 //将其中三个值设为-1
    }
    vect[9] = -1 //***注：将最后一个元素置为－1
    fmt.Println("before:", vect)

    //2.遍历切片，并删除值为-1的元素
    for i, v := range vect {
    	fmt.Print(i," ")
        if v == -1 {
            vect = append(vect[:i], vect[i+1:]...)
        }
		fmt.Println("del:", vect)
    }
    fmt.Println("after:", vect)
}
func testSlice2() {
	vect := make([]int, 1)
	vect[0] = 1
	fmt.Println("before:", vect)
	for i, v := range vect{
		if v == 1{
			vect = append(vect[:i], vect[i+1:]...)
		}
	}
	fmt.Println("after:", vect)
}

func tt1(){
	a := 1
	fmt.Println("111111111111:",a)
	func(t *int){
		b := 2
		t = &b
	}(&a)
	fmt.Println("222222222222:",a)
}