package main

import (
	"sync"
	"time"
	"fmt"
)

var synWait sync.WaitGroup
var normalMap map[int]int
var synMutex sync.Mutex

func testNum(num int) {
    for i := 1; i <= 1000000000; i++{
        synMutex.Lock()
        //normalMap[(num-1) * 10000000 + i] = (num-1) * 10000000 + i
		//normalMap[i] = (num-1) * 10000000 + i
        normalMap[i] = num
        synMutex.Unlock()
    }
    synWait.Done()  // 相当于 synWait.Add(-1)
}

func main() {
    normalMap = make(map[int]int)
    start := time.Now()
    for i := 1; i <= 10; i++ {
        synWait.Add(1)
        //testNum(i)      // 单协程操作
        go testNum(i)     // 多协程并发操作
    }
    synWait.Wait()
    end :=  time.Now()
    fmt.Println(end.Sub(start).Seconds(), len(normalMap))
}
