package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("连接服务器失败", err.Error())
		os.Exit(0)
	}
	defer conn.Close()

	num, err := conn.Write([]byte("hello"))
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(num)
	for {

	}
}
