package main

import (
	"fmt"
	"time"
)

func main() {
	panicExample5()
}

func panicExample1() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}

func panicExample2() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			fmt.Println("111111")
		}()
		fmt.Println("22222222")
	}()
	//go func() {
	//	panic("hello")
	//}()
	//time.Sleep(1 * time.Second)
	fmt.Println("33333333333")
}

func panicExample3() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("aaa = ", err)
			}
		}()
		var a int
		fmt.Println(3 / a)
	}()
	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		//defer fmt.Println(i)
		defer func() {
			fmt.Println(i)
		}()
	}
}

// recover一定要在func里
func panicExample4() {
	//defer recover()
	defer recover1()
	panic("ooo")
}
func panicExample5() {
	defer func() {
		recover()
	}()
	defer func() {
		panic("ooo")
	}()

}

func panicExample6() {
	defer func() {
		recover()
	}()
	panic("ooo")
}

func recover1() {
	recover()
}

func panicExample7() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("aaa = ", err)
		}
	}()
	//go func() {
	//	go func() {
	//		panic("go err")
	//	}()
	//}()
	//
	go func() {
		panic("go err")
	}()

	//panic("go err")
	time.Sleep(2 * time.Second)
}