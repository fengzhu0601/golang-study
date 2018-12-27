package main

import (
	"sync"
	"fmt"
	"time"
)

const TIMES  = 100000000
const NTimes = 10000000

func main() {
	//testSlice1()
	//testSlice2()
	//testMap1()
	testMap2()
}

func testSlice1(){
	start := time.Now()
	var wg sync.WaitGroup
    s := make([]int, 0, TIMES)

    for i := 0; i < TIMES; i++ {
        v := i
        wg.Add(1)
        go func() {
            s = append(s, v)
            wg.Done()
        }()
    }

    wg.Wait()
	end := time.Now()
	fmt.Printf("time = %f,unsafe %v\n", end.Sub(start).Seconds(), len(s))
}

func testSlice2(){
	start := time.Now()
	var (
		wg sync.WaitGroup
		mutex sync.Mutex
	)
	s := make([]int, 0, TIMES)
	for i:=0; i< TIMES; i++{
		v := i
		wg.Add(1)
		go func() {
			mutex.Lock()
			s = append(s, v)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("time = %f,safe %v\n", end.Sub(start).Seconds(), len(s))
}

func testMap1(){
	start := time.Now()
	//var wg sync.WaitGroup
    s := make(map[int]int)

    for i := 0; i < TIMES; i++ {
        //wg.Add(1)
        s[i] = i
        //go func() {
        //	s[i] = i
        //    wg.Done()
        //}()
    }

    //wg.Wait()
	end := time.Now()
	fmt.Printf("time = %f,unsafe %v\n", end.Sub(start).Seconds(), len(s))
}

func testMap2(){
	start := time.Now()
	var (
		wg sync.WaitGroup
		mutex sync.Mutex
	)
    s := make(map[int]int)
    n := TIMES / NTimes

    for i := 0; i< n ; i++ {
		wg.Add(1)
		go func(m int) {
			for j := 0; j < NTimes; j++{
				mutex.Lock()
				s[m*NTimes+j] = m*NTimes+j
				mutex.Unlock()
			}
			wg.Done()
		}(i)
	}

    wg.Wait()
	end := time.Now()
	fmt.Printf("time = %f,safe %v\n", end.Sub(start).Seconds(), len(s))
}

