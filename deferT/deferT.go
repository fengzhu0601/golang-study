package main
//延迟函数defer,
//1.资源的回收
//2.panic异常的捕获
//3.修改返回值
import (
	"fmt"
	"time"
	"github.com/pkg/errors"
	"os"
	"io"
)

func main(){
	//test1()
	//fmt.Println("test2 = ", test2())
	//fmt.Println("test3 = ", test3())
	//fmt.Println("test4 = ", *test4())
	//fmt.Println("test5 = ", test5())
	test6()
}
// 错误版本
func CopyFile1(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}
// 经典应用实例
func CopyFile2(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close() //每次申请资源时，请习惯立即申请一个 defer 关闭资源，这样就不会忘记释放资源了

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

// defer顺序
func test1(){
	for i:=0;i < 4; i++{
		defer fmt.Println(i)
	}
}

// 匿名返回值
func test2() int{
    i := 0
    defer func (){
        i++
        fmt.Println("defer2:", i) // 打印结果为 defer2: 2
    }()
    defer func (){
        i++
        fmt.Println("defer1:", i) // 打印结果为 defer1: 1
    }()
    return i //假如返回值是a，此时a=i，defer中修改i的值不会影响返回值a，defer也根本访问不到a
}

// 有名返回值
func test3() (i int) {
    i= 5
    defer func() {
        i++
        fmt.Println("defer2:", i) // 打印结果为 b defer2: 7
    }()
    defer func() {
        i++
        fmt.Println("defer1:", i) // 打印结果为 b defer1: 6
    }()
    return i // 返回值就是i，defer修改就是i的值
}

// 匿名返回值是指针的情况
func test4() *int {
    i := 9
    defer func() {
        i++
        fmt.Println("defer2:", i, &i) // 打印结果为 defer2: 11 0xc042052088
    }()
    defer func() {
        i++
        fmt.Println("defer1:", i, &i) // 打印结果为 defer1: 10 0xc042052088
    }()
    return &i
}

func test5() int{
	i := 21

	defer func(i *int){
		*i++
		fmt.Println("i2 = ", *i)//24
	}(&i)

	defer func(i int) {
		i++
		fmt.Println("i1 = ", i)//22
	}(i)

	i++   //22

	defer func(i *int){
		*i++
		fmt.Println("i3 = ", *i ) //23
	}(&i)

	defer func(i int) {
		i++
		fmt.Println("i4 = ", i  )//23)
	}(i)

	return i  //22
}

// defer作用域
func test6(){
	defer func() {
		if r:=recover(); r != nil{
			fmt.Println("recover ===",r)
		}
	}()
	e := errors.New("error")
	fmt.Println(e)
	//panic(e)
	//os.Exit(1)

	defer fmt.Println("defer")
	//go func() { panic(e) }()
	//panic(e)
	time.Sleep(1e9)
	fmt.Println("over.")
	os.Exit(1)
}
