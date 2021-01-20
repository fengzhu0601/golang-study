package main

import (
	"net"
	"fmt"
	"bufio"
	"io"
)

func main(){
	listen , err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("网络监听失败:", err)
	}
	defer listen.Close()

	for{
		c, err := listen.Accept()
		if err != nil {
			fmt.Println("错误的链接")
		}
		go Handle(c)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && err != io.EOF{
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}
