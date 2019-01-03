package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"io"
	"os"
)

func main() {
	// 计算字符串的md5
	h := md5.New()
	h.Write([]byte("123456"))
	fmt.Printf("====%s\n", hex.EncodeToString(h.Sum(nil)))

	// 计算文件的md5
	fileName := "./md5Test/md5Test1.txt"
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("os Open error", file)
	}

	defer func(){
		file.Close()
	}()

	md5 := md5.New()
	io.Copy(md5, file)
	fmt.Printf("%s\n", hex.EncodeToString(md5.Sum(nil)))

}


